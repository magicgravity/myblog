package service

import (
	"testing"
	"github.com/magicgravity/myblog/model"
)

func TestLogin(t *testing.T) {
	ok,u := Login("admin","123456")
	if !ok {
		t.Fatal("登陆失败")
	}else{
		t.Log("登陆成功,用户信息=== >>> ",u)
	}
}


func TestAddUser(t *testing.T) {
	user := model.User{}
	user.UserName = "ceshi344"
	user.GroupName = "测试组11"
	user.HomeUrl ="http://www.baidu.com"
	user.Password = "333333"
	user.ScreenName = "测试用户4444"
	user.Email = "ceshi_9153@163.com"

	ok,uid := AddUser(user)
	if ok {
		t.Logf("增加用户成功,uid ==> %d",uid)
	}else{
		t.Error("增加用户失败")
	}

}


func TestModifyUserPwd(t *testing.T) {
	//333333		c677bce3f82801e0c731404578eab800
	ok := ModifyUserPwd("4","333333","222222")
	if ok {
		t.Log("修改密码成功")
	}else{
		t.Error("修改密码失败")
	}
}