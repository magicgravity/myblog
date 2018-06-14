package dao

import (
	"github.com/magicgravity/myblog/db"
	"github.com/golang/glog"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func InsertLog(tx *gdb.Tx,action ,data ,ip string,authId,cdate uint32)bool {
	var model *gdb.Model
	if tx==nil {
		model = db.GetMySqlDbFromCfg().Table("t_logs")
	}else{
		model = tx.Table("t_logs")
	}
	r,err := model.Data(g.Map{	"action":action,
																		"data":data,
																		"ip":ip,
																		"author_id":authId,
																		"created":cdate,
																		}).Insert()
	if err!= nil{
		glog.Error("insert into log fail ,reason ==> ",err)
		return false
	}else{
		if lid,err := r.LastInsertId();err== nil&&lid>0 {
			return true
		}else{
			return false
		}
	}
}
