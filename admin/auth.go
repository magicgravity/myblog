package admin

import "gitee.com/johng/gf/g/net/ghttp"

func init(){
	ghttp.GetServer().BindHandler("post:/admin/login",DoLogin)
	ghttp.GetServer().BindHandler("/admin/logout",Logout)
}


func DoLogin(r *ghttp.Request){

}

func Logout(r *ghttp.Request){

}
