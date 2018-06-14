package dao


import (
	"github.com/magicgravity/myblog/db"
	"github.com/golang/glog"
	"gitee.com/johng/gf/g"
	"github.com/magicgravity/myblog/model"
	"gitee.com/johng/gf/g/database/gdb"
)

func InsertMetaStruct(tx *gdb.Tx,meta model.Meta)(bool,uint32){
	return InsertMeta(tx,meta.Name,meta.Slug,meta.Description,meta.Type,meta.Sort)
}

func InsertMeta(tx *gdb.Tx,name,slug,desc,mtype string,sort uint32)(bool,uint32){
	var model *gdb.Model
	if tx==nil {
		model = db.GetMySqlDbFromCfg().Table("t_metas")
	}else{
		model = tx.Table("t_metas")
	}
	r,err := model.Data(g.Map{	"name":name,
																		"slug":slug,
																		"type":mtype,
																		"description":desc,
	}).Insert()

	if err!= nil{
		glog.Error("insert into meta fail ,reason ==> ",err)
		return false,0
	}else{
		if lid,err := r.LastInsertId();err== nil&&lid>0 {
			return true,uint32(lid)
		}else{
			return false,0
		}
	}
}



func UpdateMeta(tx *gdb.Tx,meta model.Meta)bool {
	var model *gdb.Model
	if tx==nil {
		model = db.GetMySqlDbFromCfg().Table("t_metas")
	}else{
		model = tx.Table("t_metas")
	}
	r,err := model.Data(g.Map{
		"mid":meta.Mid,
		"name":meta.Name,
		"slug":meta.Slug,
		"type":meta.Type,
		"description":meta.Description,
	}).Save()

	if err!= nil{
		glog.Error("update meta fail ,reason ==> ",err)
		return false
	}else{
		if lid,err := r.RowsAffected();err== nil&&lid>0 {
			return true
		}else{
			glog.Error("update meta fail ",err)
			return false
		}
	}
}


func DeleteMetaById(tx *gdb.Tx,mid string)bool {
	var model *gdb.Model
	if tx==nil {
		model = db.GetMySqlDbFromCfg().Table("t_metas")
	}else{
		model = tx.Table("t_metas")
	}
	r,err := model.Where("mid=?",mid).Delete()
	if err!= nil {
		glog.Error("delete meta fail ,reason == > ",err)
		return false
	}else{
		if lid,err := r.RowsAffected();lid>0&& err==nil {
			return true
		}else{
			glog.Error("delete meta fail ",err)
			return false
		}
	}
}


func SelectMetaByNameType(name,mtype string)map[string]interface{}{
	r,err :=db.GetMySqlDbFromCfg().Table("t_metas").Where("name=?",name).And("type=?",mtype).One()
	if err!= nil {
		glog.Error("find meta by name type fail ,reason == >",err)
		return nil
	}else{
		return r.ToMap()
	}
}

func SelectMetasByType(mtype string)[]map[string]interface{}{
	r,err :=db.GetMySqlDbFromCfg().Table("t_metas").Where("type=?",mtype).OrderBy("sort desc, mid desc").All()
	if err!= nil {
		glog.Error("find meta by  type fail ,reason == >",err)
		return nil
	}else{
		return r.ToList()
	}
}


func SelectMetaByMid(mid uint32)map[string]interface{}{
	r,err := db.GetMySqlDbFromCfg().Table("t_metas").Where("mid=?",mid).One()
	if err!= nil {
		glog.Error("find meta by  mid fail ,reason == >",err)
		return nil
	}else{
		return r.ToMap()
	}
}


func UpdateMetaNameByMid(tx *gdb.Tx,mid uint32,name string)bool {
	var model *gdb.Model
	if tx==nil {
		model = db.GetMySqlDbFromCfg().Table("t_metas")
	}else{
		model = tx.Table("t_metas")
	}
	r,err := model.Data(g.Map{
		"name":name,
	}).Where("mid=?",mid).Update()
	if err!= nil {
		glog.Error("update meta name by mid fail ,reason ==> ",err)
		return false
	}else{
		if lid,err := r.RowsAffected();err==nil && lid>0{
			return true
		}else{
			glog.Error("update meta name by mid fail ,reason ==> ",err)
			return false
		}
	}
}