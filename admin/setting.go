package admin

import "gf/g/net/ghttp"

func init(){
	ghttp.GetServer().BindHandler("get:/admin/setting",Setting)
	ghttp.GetServer().BindHandler("post:/admin/setting",SaveSetting)
	ghttp.GetServer().BindHandler("post:/admin/setting/backup",BackUpSetting)

}

/*
系统设置
 */
func Setting(r *ghttp.Request){

}

/*
保存设置
 */
func SaveSetting(r *ghttp.Request){

}

/*
备份设置
 */
func BackUpSetting(r *ghttp.Request){

}