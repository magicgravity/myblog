package admin

import "gitee.com/johng/gf/g/net/ghttp"

func init(){
	ghttp.GetServer().BindHandler("get:/admin/category",IndexCategory)
	ghttp.GetServer().BindHandler("post:/admin/category/save",SaveCategory)
	ghttp.GetServer().BindHandler("/admin/category/delete",DeleteCategory)
}


func IndexCategory(r *ghttp.Request){

}

func SaveCategory(r *ghttp.Request){

}

func DeleteCategory(r *ghttp.Request){

}