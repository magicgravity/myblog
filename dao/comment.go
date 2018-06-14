package dao

import (
	"github.com/magicgravity/myblog/db"
	"github.com/golang/glog"
	"gitee.com/johng/gf/g"
	"github.com/magicgravity/myblog/model"
	"gitee.com/johng/gf/g/database/gdb"
)


func InsertComment(tx *gdb.Tx,comment model.Comments)(bool,uint32){
	var model *gdb.Model
	if tx==nil {
		model = db.GetMySqlDbFromCfg().Table("t_comments")
	}else{
		model = tx.Table("t_comments")
	}
	r,err := model.Data(g.Map{
		"author":comment.Author,
		"cid":comment.Cid,
		"author_id":comment.AuthorId,
		"owner_id":comment.OwnerId,
		"mail":comment.Mail,
		"url":comment.Url,
		"ip":comment.Ip,
		"agent":comment.Agent,
		"content":comment.Content,
		"type":comment.Type,
		"status":comment.Status,
		"parent":comment.Parent,
	}).Insert()

	if err!= nil{
		glog.Error("insert into comment fail ,reason ==> ",err)
		return false,0
	}else{
		if lid,err := r.LastInsertId();err== nil&&lid>0 {
			return true,uint32(lid)
		}else{
			return false,0
		}
	}
}

func SelectCommentListByCid(cid uint32)[]map[string]interface{}{
	if cid>0 {
		r,err :=db.GetMySqlDbFromCfg().Table("t_comments").Where("cid=?",cid).And("parent=?",0).And("status is not null").And("status = ?","approved").OrderBy("coid desc").All()
		if err!= nil {
			glog.Error("select comment list by cid fail ,reason ==> ",err)
			return nil
		}else{
			return r.ToList()
		}
	}else{
		return nil
	}
}


func DeleteCommentByCoId(tx *gdb.Tx,coid uint32)bool{
	if coid >0 {
		var model *gdb.Model
		if tx==nil {
			model = db.GetMySqlDbFromCfg().Table("t_comments")
		}else{
			model = tx.Table("t_comments")
		}
		r,err := model.Where("coid=?",coid).Delete()
		//r,err := db.GetMySqlDbFromCfg().Delete("t_comments","coid=?",coid)
		if err!= nil {
			glog.Error("delete comment  by coid fail ,reason ==> ",err)
			return false
		}else{
			if lid,err := r.RowsAffected();lid>0&&err==nil {
				return true
			}else{
				glog.Error("delete comment  by coid fail ,reason ==> ",err)
				return false
			}
		}
	}else{
		return false
	}
}

func SelectCommentByCoId(coid uint32)map[string]interface{}{
	if coid >0 {
		r,err := db.GetMySqlDbFromCfg().Table("t_comments").Where("coid=?",coid).One()
		if err!= nil {
			glog.Error("select comment  by coid fail ,reason ==> ",err)
			return nil
		}else{
			return r.ToMap()
		}
	}else{
		return nil
	}
}


