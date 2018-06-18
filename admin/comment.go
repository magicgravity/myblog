package admin

import "gitee.com/johng/gf/g/net/ghttp"

func init(){
	ghttp.GetServer().BindHandler("get:/admin/comments",IndexComment)
	ghttp.GetServer().BindHandler("post:/admin/comments/delete",DeleteComment)
	ghttp.GetServer().BindHandler("post:/admin/comments/status",DeleteCommentStatus)
}


func IndexComment(r *ghttp.Request){

}


func DeleteComment(r *ghttp.Request){

}

func DeleteCommentStatus(r *ghttp.Request){

}

