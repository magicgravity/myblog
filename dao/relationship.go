package dao

import (
	"gitee.com/johng/gf/g/database/gdb"
	"github.com/magicgravity/myblog/db"
	"github.com/golang/glog"
	"github.com/magicgravity/myblog/model"
	"gitee.com/johng/gf/g"
)

func DeleteRelationShipById(tx *gdb.Tx,cid,mid uint32)bool{
	if cid <=0 && mid <=0 {
		return false
	}
	var model *gdb.Model
	if tx==nil {
		model = db.GetMySqlDbFromCfg().Table("t_relationships")
	}else{
		model = tx.Table("t_relationships")
	}
	model = model.Where("1=1")
	if cid >0 {
		model = model.And("cid=?",cid)
	}
	if mid>0 {
		model = model.And("mid=?",mid)
	}
	r,err := model.Delete()
	if err!= nil {
		glog.Error("delete relationship fail ,reason ==> ",err)
		return false
	}else{
		if lid,err := r.RowsAffected();err==nil && lid>0 {
			return true
		}else{
			glog.Error("delete relationship fail ,reason ==> ",err)
			return false
		}
	}
}



func SelectRelationshipsById(cid,mid uint32)[]map[string]interface{}{
	if cid<=0 && mid <=0 {
		return nil
	}
	model := db.GetMySqlDbFromCfg().Table("t_relationships").Where("1=1")
	if cid >0 {
		model = model.And("cid=?",cid)
	}
	if mid >0 {
		model = model.And("mid=?",mid)
	}
	r,err := model.All()
	if err!= nil {
		glog.Error("select relationships by id fail ,reason == >",err)
		return nil
	}else{
		return r.ToList()
	}
}


func InsertRelationship(tx *gdb.Tx,r model.RelationShip)bool{
	if r.Cid>0 && r.Mid>0 {
		var model *gdb.Model
		if tx == nil {
			model = db.GetMySqlDbFromCfg().Table("t_relationships")
		} else {
			model = tx.Table("t_relationships")
		}
		_,err :=model.Data(g.Map{
			"cid":r.Cid,
			"mid":r.Mid,
		}).Insert()
		if err!=nil {
			return false
		}else{
			return true
		}
	}else{
		return false
	}
}


func SelectRelationshipsCountById(cid,mid uint32)int{
	if cid<=0 && mid <=0 {
		return 0
	}
	model := db.GetMySqlDbFromCfg().Table("t_relationships").Where("1=1")
	if cid >0 {
		model = model.And("cid=?",cid)
	}
	if mid >0 {
		model = model.And("mid=?",mid)
	}
	r,err := model.Count()
	if err!= nil {
		glog.Error("select relationships count by id fail ,reason == >",err)
		return 0
	}else{
		return r
	}
}