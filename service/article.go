package service

import (
	"github.com/magicgravity/myblog/dao"
	"github.com/magicgravity/myblog/model"
	"gitee.com/johng/gf/g/util/gvalid"
	"github.com/golang/glog"
	"github.com/magicgravity/myblog/util"
	"github.com/magicgravity/myblog/db"
	"gitee.com/johng/gf/g/database/gdb"
)

func GetArticlesByKeyword(keyword string)[]map[string]interface{}{
	return dao.SelectContentByKeywordType(keyword,"post","publish")
}


func PublishArticle(content model.Contents)(bool,uint32){

	if !checkContentValid(content) {
		glog.Error("publish article param is not valid")
		return false,0
	}else{
		slug := content.Slug
		if len(slug)>0 {
			//检查是否合法
			if len(slug)<5{
				glog.Error("slug length is too short")
				return false,0
			}else{
				if !util.IsPath(slug) {
					glog.Error("slug path is not valid")
					//您输入的路径不合法
					return false,0
				}

				if cc :=dao.CountContentListByTypeStatus(content.Type,slug);cc>0 {
					//该路径已经存在，请重新输入
					glog.Error("slug is already exist")
					return false,0
				}
			}
		}

		//TODO  EmojiParser 需要添加对emoji 的支持，解析content内容
		content.Hits= 0
		content.CommentsNum = 0
		content.Created = util.GetCurTimeAsInt()
		content.Modified = content.Created

		tx,err := db.GetMySqlDbFromCfg().Begin()
		if err== nil{
			ok,cid := dao.InsertContent(tx,content)
			if ok {
				//插入content 成功 还需要保存meta

				ok1 := SaveMeta(tx,"tag",content.Tags,cid)
				ok2 := SaveMeta(tx,"category",content.Categories,cid)
				if ok1 && ok2 {
					if err := tx.Commit(); err == nil {
						return true, cid
					} else {
						glog.Error("commit transaction fail ，reason ==> ", err)
						tx.Rollback()
						return false, 0
					}
				}else{
					glog.Error("save meta fail ,prepare to rollback ")
					tx.Rollback()
					return false,0
				}
			}else{
				tx.Rollback()
				return false,0
			}
		}else{
			return false,0
		}
	}
}


func UpdateArticleCategory(tx *gdb.Tx,ordinal,newCatefory string)bool {
	if ordinal==newCatefory {
		return true
	}else{
		ok,_ := dao.UpdateContentCategory(tx,ordinal,newCatefory)
		return ok
	}
}


func UpdateContentByCid(c model.Contents){
	if c.Cid>0 {
		dao.UpdateContent(nil,c)
	}else{
		glog.Error("not found content cid ,can't update ")
	}
}

func GetArticlesByCatalog(mid uint32){
	//TODO  待分析
}


func DeleteArticleByCid(cid uint32)bool{
	tx,err := db.GetMySqlDbFromCfg().Begin()
	if err== nil {
		ok := dao.DeleteContentByCid(tx, cid)
		if ok {
			if dao.DeleteRelationShipById(tx,cid,0) && tx.Commit()==nil {
				return true
			}else{
				tx.Rollback()
				return false
			}
		}else{
			glog.Error("delete content by cid fail ,prepare to rollback")
			tx.Rollback()
			return false
		}
	}else{
		return false
	}
}


func UpdateArticle(content model.Contents)bool {
	if !checkContentValid(content){
		glog.Error("update article param is not valid")
		return false
	}else{
		//TODO  EmojiParser 需要添加对emoji 的支持，解析content内容
		content.Modified = util.GetCurTimeAsInt()

		tx,err := db.GetMySqlDbFromCfg().Begin()
		if err== nil {
			if dao.UpdateContent(tx,content) &&
				dao.DeleteRelationShipById(tx,content.Cid,0) &&
					SaveMeta(tx,"tag",content.Tags,content.Cid) &&
						SaveMeta(tx,"category",content.Categories,content.Cid) {
				return true
			}else{
				glog.Error("update article fail ,prepare to rollback")
				tx.Rollback()
				return false
			}
		}else{
			return false
		}
	}
}


func checkContentValid(content model.Contents)bool{
	params := map[string]interface{} {
		"title"  : content.Title,
		"content"  : string(content.Content),
		"authorId" : content.AuthorId,
	}
	rules := map[string]string {
		"title"  :  "required|length:1,200",
		"content"  : "required|length:1,200000",
		"authorId" :  "required|min:1",
	}
	msgs  := map[string]interface{} {
		"title" : "标题不能为空|标题长度应当在:min到:max之间",
		"content" : "内容不能为空|内容长度应当在:min到:max之间",
		"authorId": "请登陆后发布文章|请登陆后发布文章",
	}
	verifyRes :=gvalid.CheckMap(params,rules,msgs)
	if len(verifyRes)>0 {
		glog.Error("publish article param is not valid, reason ==>>",verifyRes)
		return false
	}else{
		return true
	}
}