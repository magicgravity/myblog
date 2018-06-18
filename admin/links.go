package admin

import "gitee.com/johng/gf/g/net/ghttp"

func init(){
	ghttp.GetServer().BindHandler("get:/admin/links",IndexLink)
	ghttp.GetServer().BindHandler("post:/admin/links/save",SaveLink)
	ghttp.GetServer().BindHandler("/admin/links/delete",DeleteLink)
}

func IndexLink(r *ghttp.Request){

}

func SaveLink(r *ghttp.Request){

}


func DeleteLink(r *ghttp.Request){

}