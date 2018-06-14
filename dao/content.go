package dao

import (
	"github.com/magicgravity/myblog/model"
	"github.com/magicgravity/myblog/db"
	"github.com/golang/glog"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func InsertContent(tx *gdb.Tx,c model.Contents)(bool,uint32){
	var model *gdb.Model
	if tx==nil {
		model = db.GetMySqlDbFromCfg().Table("t_contents")
	}else{
		model = tx.Table("t_contents")
	}
	r,err := model.Data(g.Map{
		"title":c.Title,
		"slug":c.Slug,
		"content":c.Content,
		"type":c.Type,
		"status":c.Status,
		"tags":c.Tags,
		"categories":c.Categories,
		"hits":c.Hits,
		"comments_num":c.CommentsNum,
		"author_id":c.AuthorId,
		"created":c.Created,
		"modified":c.Modified,
	}).Insert()

	if err!= nil{
		glog.Error("insert into content fail ,reason ==> ",err)
		return false,0
	}else{
		if lid,err := r.LastInsertId();err== nil&&lid>0 {
			return true,uint32(lid)
		}else{
			return false,0
		}
	}
}


func SelectContentListByTypeStatus(mtype,status string)[]map[string]interface{}{
	r,err := db.GetMySqlDbFromCfg().Table("t_contents").Where("type=?",mtype).And("status=?",status).OrderBy("created desc").All()
	if err!= nil {
		glog.Error("select content by type status fail ,reason ==> ",err)
		return nil
	}else{
		return r.ToList()
	}
}

func CountContentListByTypeStatus(mtype,status string)int {
	r,err := db.GetMySqlDbFromCfg().Table("t_contents").Where("type=?",mtype).And("status=?",status).Count()
	if err!= nil {
		glog.Error("count contents by type status fail ,reason ==> ",err)
		return 0
	}else{
		return r
	}
}

func SelectContentByCid(cid uint32)map[string]interface{}{
	r,err := db.GetMySqlDbFromCfg().Table("t_contents").Where("cid=?",cid).One()
	if err!= nil {
		glog.Error("select content by cid fail ,reason ==> ",err)
		return nil
	}else{
		return r.ToMap()
	}
}

func UpdateContent(tx *gdb.Tx,c model.Contents)bool{
	var model *gdb.Model
	if tx==nil {
		model = db.GetMySqlDbFromCfg().Table("t_contents")
	}else{
		model = tx.Table("t_contents")
	}
	r,err := model.Data(g.Map{
		"cid":c.Cid,
		"title":c.Title,
		"slug":c.Slug,
		"modified":c.Modified,
		"content":c.Content,
		"author_id":c.AuthorId,
		"type":c.Type,
		"status":c.Status,
		"tags":c.Tags,
		"categories":c.Categories,
		"hits":c.Hits,
		"comments_num":c.CommentsNum,
	}).Save()
	if err!= nil{
		glog.Error("update content fail ,reason ==> ",err)
		return false
	}else{
		if lid,err := r.RowsAffected();err== nil&&lid>0 {
			return true
		}else{
			return false
		}
	}
}


func SelectContentByKeywordType(keyword,ctype ,status string)[]map[string]interface{}{
	r,err :=db.GetMySqlDbFromCfg().Table("t_contents").Where("type=?",ctype).And("status=?",status).And("title like ?","%"+keyword+"%").OrderBy("created desc").All()
	if err!= nil {
		glog.Error("select  content by keyword type fail ,reason ==> ",err)
		return nil
	}else{
		return r.ToList()
	}
}


func DeleteContentByCid(tx *gdb.Tx,cid uint32)bool{
	var model *gdb.Model
	if tx==nil {
		model = db.GetMySqlDbFromCfg().Table("t_contents")
	}else{
		model = tx.Table("t_contents")
	}
	r,err := model.Where("cid=?",cid).Delete()
	//r,err := db.GetMySqlDbFromCfg().Delete("t_contents","cid=?",cid)
	if err!= nil {
		glog.Error("delete  content by cid fail ,reason ==> ",err)
		return false
	}else{
		if lid,err:= r.RowsAffected();lid>0&&err==nil {
			return true
		}else{
			glog.Error("delete  content by cid fail ,reason ==> ",err)
			return false
		}

	}
}


func UpdateContentCategory(tx *gdb.Tx,oldC,newC string)(bool,int64) {
	var model *gdb.Model
	if tx==nil {
		model = db.GetMySqlDbFromCfg().Table("t_contents")
	}else{
		model = tx.Table("t_contents")
	}
	r,err :=model.Data(g.Map{
		"categories":newC,
	}).Where("categories=?",oldC).Update()
	if err!= nil {
		glog.Error("update content category fail ,reason ==> ",err)
		return false,0
	}else{
		if lc,err := r.RowsAffected();err==nil {
			return true,lc
		}else{
			glog.Error("update content category fail ,reason ==> ",err)
			return false,0
		}
	}
}