package dao

import (
	"github.com/magicgravity/myblog/db"
	"github.com/golang/glog"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g"
)

func SelectAttachById(aid uint32)map[string]interface{}{
	if aid>0 {
		r,err :=db.GetMySqlDbFromCfg().Table("t_attach").Where("id=?",aid).One()
		if err!= nil {
			glog.Error("select attach by id fail ,reason == > ",err)
			return nil
		}else{
			return r.ToMap()
		}
	}else{
		return nil
	}
}


func InsertAttach(tx *gdb.Tx,fname,fkey,ftype string,authorId,created uint32)bool{
	var model *gdb.Model
	if tx==nil {
		model = db.GetMySqlDbFromCfg().Table("t_attach")
	}else{
		model = tx.Table("t_attach")
	}
	_,err := model.Data(g.Map{
		"fname":fname,
		"ftype":ftype,
		"fkey":fkey,
		"author_id":authorId,
		"created":created,
	}).Insert()
	if err!= nil {
		glog.Error("insert attach fail ,reason ==> ",err)
		return false
	}else{
		return true
	}
}


func DeleteAttachById(tx *gdb.Tx,aid uint32)bool {
	if aid>0 {
		var model *gdb.Model
		if tx==nil {
			model = db.GetMySqlDbFromCfg().Table("t_attach")
		}else{
			model = tx.Table("t_attach")
		}
		if _,err:=model.Where("id=?",aid).Delete();err==nil {
			return true
		}else{
			glog.Error("delete attach by id fail ,reason == > ",err)
			return false
		}
	}else{
		return false
	}
}


func SelectAllAttachs()[]map[string]interface{}{
	r,err := db.GetMySqlDbFromCfg().Table("t_attach").OrderBy("id desc").All()
	if err!= nil {
		glog.Error("select all attachs fail ,reason == > ",err)
		return nil
	}else{
		return r.ToList()
	}
}