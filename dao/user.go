package dao

import (
	"github.com/magicgravity/myblog/db"
	"github.com/magicgravity/myblog/model"
	"github.com/golang/glog"
	"gitee.com/johng/gf/g"
	//"gitee.com/johng/gf/g/database/gdb"

	"gitee.com/johng/gf/g/database/gdb"
)

func QueryUserByUserNamePwd(id ,pwd string)*model.User{
	glog.Infof("QueryUserByUserNamePwd ==> username=> %s ; pwd=> %s ",id,pwd)
	r,err := db.GetMySqlDbFromCfg().Table("t_users").Where("username = ?",id).And("password = ?",pwd).One()
	//如果没查到 len(r) 是0
	glog.Info("----------------ret len->",len(r))
	if err!= nil || len(r)==0  {
		if err!= nil {
			glog.Warning("query user by username and pwd fail, reason == >> ", err)
		}else{
			glog.Warning("query user by username and pwd fail,no data found")
		}
		return nil
	}else{
		user := model.User{}
		if err:= r.ToStruct(&user);err== nil{
			return &user
		}else{
			glog.Warning("query user by username and pwd ok,but toStuct fail, reason == >> ",err)
			return nil
		}
	}
}

func InsertUser(tx *gdb.Tx,u model.User)(bool,uint32){
	var model *gdb.Model
	if tx==nil {
		model = db.GetMySqlDbFromCfg().Table("t_users")
	}else{
		model = tx.Table("t_users")
	}
	r,err := model.Data(g.Map{	"username":u.UserName,
														"password":u.Password,
														"email":u.Email,
														"home_url":u.HomeUrl,
														"screen_name":u.ScreenName,
														"created":u.Created,
														"activated":u.Activated,
														"logged":u.Logged,
														"group_name":u.GroupName}).Insert()

	//r,err :=db.GetMySqlDbFromCfg().Insert("t_users",gdb.Map(g.Map{	"username":u.UserName,
	//												"password":u.Password,
	//												"email":u.Email,
	//												"home_url":u.HomeUrl,
	//												"screen_name":u.ScreenName,
	//												"created":u.Created,
	//												"activated":u.Activated,
	//												"logged":u.Logged,
	//												"group_name":u.GroupName}))
	if err!= nil{
		glog.Error("insert into user fail ,reason ==> ",err)
		return false,0
	}else{
		if lid,err := r.LastInsertId();err== nil {
			return true,uint32(lid)
		}else{
			return false,0
		}
	}
}


func ModifyUser(tx *gdb.Tx,u model.User)bool{
	var model *gdb.Model
	if tx==nil {
		model = db.GetMySqlDbFromCfg().Table("t_users")
	}else{
		model = tx.Table("t_users")
	}
	glog.Infof("ModifyUser ==> user => %v  ",u)
	r,err := model.Data(g.Map{	"username":u.UserName,
																"password":u.Password,
																"email":u.Email,
																"home_url":u.HomeUrl,
																"screen_name":u.ScreenName,
																"created":u.Created,
																"activated":u.Activated,
																"logged":u.Logged,
																"group_name":u.GroupName,
																"uid":u.Uid}).Save()

	if err!= nil{
		glog.Error("modify  user fail ,reason ==> ",err)
		return false
	}else{
		if lid,err := r.RowsAffected();err== nil && lid >0 {
			return true
		}else{
			return false
		}
	}
}

func UpdateUserPwdById(tx *gdb.Tx,id , pwd string) bool{
	var model *gdb.Model
	if tx==nil {
		model = db.GetMySqlDbFromCfg().Table("t_users")
	}else{
		model = tx.Table("t_users")
	}
	r,err := model.Data(g.Map{"password":pwd}).Where("uid=?",id).Update()
	//r,err := db.GetMySqlDbFromCfg().Update("t_users","password=?","uid=?",pwd,id)
	if err!= nil  {
		glog.Error("update user pwd fail ,reason == > ",err)
		return false
	}else{
		if lid,err := r.RowsAffected();err==nil && lid>0  {
			return true
		}else{
			return false
		}
	}
}

func FindUserByUid(id string) map[string]interface{}{
	glog.Infof("FindUserByUid ==> uid=> %s ",id)
	r,err := db.GetMySqlDbFromCfg().Table("t_users").Where("uid = ?",id).One()
	//如果没查到 len(r) 是0
	glog.Info("----------------ret len->",len(r))
	if err!= nil || len(r)==0  {
		if err!= nil {
			glog.Warning("query user by uid  fail, reason == >> ", err)
		}else{
			glog.Warning("query user by uid  fail,no data found")
		}
		return nil
	}else{
		glog.Info("----------> ",r)
		return r.ToMap()
	}
}