package admin

import "gitee.com/johng/gf/g/net/ghttp"

func init(){
	ghttp.GetServer().BindHandler("get:/admin/attach",IndexAttach)
	ghttp.GetServer().BindHandler("post:/admin/attach/upload",UploadAttach)
	ghttp.GetServer().BindHandler("/admin/attach/delete",DeleteAttach)

}

func IndexAttach(r *ghttp.Request){

}

func UploadAttach(r *ghttp.Request){

}

func DeleteAttach(r *ghttp.Request){

}