package main

import (
	"github.com/magicgravity/myblog/service"

	"github.com/golang/glog"
)

func main(){


	//ok,u :=service.Login("admin","123456")
	//if ok{
	//	glog.Info("get user =========> ",u)
	//}

	//ok := service.ModifyUserPwd("4","3333","111111")
	//if ok{
	//	glog.Info("upedate user pwd ok ")
	//}

	list := service.GetArticlesByKeyword("标题")
	if list!= nil {
		glog.Info("query result ==> ",list)
	}
}
