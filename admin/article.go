package admin

import "gitee.com/johng/gf/g/net/ghttp"

func init(){
	ghttp.GetServer().BindHandler("get:/admin/article",IndexArticle)
	ghttp.GetServer().BindHandler("get:/admin/article/publish",NewArticle)
	ghttp.GetServer().BindHandler("get:/admin/article/:cid",EditArticle)
	ghttp.GetServer().BindHandler("post:/admin/article/publish",PublishArticle)
	ghttp.GetServer().BindHandler("post:/admin/article/modify",ModifyArticle)
	ghttp.GetServer().BindHandler("/admin/article/delete",DeleteArticle)
}


func IndexArticle(r *ghttp.Request){

}

func NewArticle(r *ghttp.Request){

}

func EditArticle(r *ghttp.Request){

}

func PublishArticle(r *ghttp.Request){

}


func ModifyArticle(r *ghttp.Request){

}

func DeleteArticle(r *ghttp.Request){

}