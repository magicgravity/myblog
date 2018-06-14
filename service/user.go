package service

import (
	"github.com/magicgravity/myblog/dao"
	"github.com/magicgravity/myblog/model"
	"github.com/magicgravity/myblog/util"
	"time"
	"strconv"
	"github.com/golang/glog"
	"gitee.com/johng/gf/g/util/gvalid"
)
/*
插入user对象 返回成功与否和 user的uid
 */
func AddUser(user model.User)(bool,uint32){
	if len(user.UserName)<=0 || len(user.Email)<=0 {
		return false,0
	}else{
		user.Logged = 0
		user.Activated = 0
		user.Password = util.MD5encode(user.UserName+user.Password)
		if curtime,err := strconv.Atoi(util.FormatDate(time.Now(),"yyyyMMddhhmiss"));err== nil{
			user.Created = uint32(curtime)
		}else{
			user.Created = 0
		}
		return dao.InsertUser(nil,user)
	}
}


func Login(userName,pwd string)(bool,*model.User){
	user := dao.QueryUserByUserNamePwd(userName,util.MD5encode(userName+pwd))
	if user!=nil {
		return true,user
	}else{
		return false,nil
	}
}

func ModifyUserPwd(uid ,oldPwd,newPwd string)bool{
	params := map[string]interface{} {
		"passport"  : uid,
		"password"  : oldPwd,
		"password2" : newPwd,
	}
	rules := map[string]string {
		"passport"  :  "required|length:1,32",
		"password"  : "required|length:6,16|different:password2",
		"password2" :  "required|length:6,16",
	}
	msgs  := map[string]interface{} {
		"passport" : "账号不能为空|账号长度应当在:min到:max之间",
		"password" : map[string]string {
			"required" : "密码不能为空",
			"different"     : "两次密码输入相等",
		},
	}
	verifyRes :=gvalid.CheckMap(params,rules,msgs)
	glog.Info("-------------->",verifyRes)
	if len(verifyRes)>0 {
		return false
	}

	user := dao.FindUserByUid(uid)
	glog.Info("user info ===>>",user)
	username := user["username"].(string)

	//glog.Info("--------11111---",util.MD5encode(username+oldPwd))
	//glog.Info("--------11111---",user["password"])
	//glog.Info("--------11111---",username)
	if len(user)>0 && user["password"] == util.MD5encode(username+oldPwd) {
		//按id 查到结果 且 与旧密码相符
		return dao.UpdateUserPwdById(nil,uid,util.MD5encode(username+newPwd))
	}else{
		//id 或者密码不对
		glog.Error("modify user password fail,uid or pwd is not right")
		return false
	}
}

