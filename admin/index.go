package admin

import "gitee.com/johng/gf/g/net/ghttp"

func init(){
	ghttp.GetServer().BindHandler("get:/admin",Index)
	ghttp.GetServer().BindHandler("get:/admin/index",Index)
	ghttp.GetServer().BindHandler("get:/admin/profile",Profile)
	ghttp.GetServer().BindHandler("post:/admin/profile",SaveProfile)
	ghttp.GetServer().BindHandler("post:/admin/password",UpPwd)
}


func Index(r *ghttp.Request){

}

func Profile(r *ghttp.Request){

}

func SaveProfile(r *ghttp.Request){

}

func UpPwd(r *ghttp.Request){

}

