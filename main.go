package main

import (
	"github.com/magicgravity/myblog/service"

	//"github.com/golang/glog"
	"log"
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

	//list := service.GetArticlesByKeyword("标题")
	//if list!= nil {
	//	glog.Info("query result ==> ",list)
	//}

	var totalPages uint64= 21
	var curPage uint64 = 19
	for curPage<=totalPages {
		page := service.GetArticleContents(curPage, 5)
		if page!= nil &&len(page.GetList()) > 0 {
			for idx, val := range page.GetList() {
				log.Printf("[%d]>>>>[%d]  -----value===> %v", curPage,idx, val["cid"])
			}
			log.Printf(page.ToString())
		}else{
			break
		}

		log.Printf("~~~~~~~~~~~~~~~~~~~~~~~")
		curPage++
	}
}
