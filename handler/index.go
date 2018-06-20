package handler

import (
	"gitee.com/johng/gf/g/net/ghttp"
	"strconv"
	"gitee.com/johng/gf/g/frame/gins"
)

func init(){
	ghttp.GetServer().BindHandler("get:/",     Index)
	ghttp.GetServer().BindHandler("get:/page/:pageNo",IndexPage)
	ghttp.GetServer().BindHandler("get:/article/:cid",GetArticle)
	ghttp.GetServer().BindHandler("get:/article/:cid.html",GetArticle)
	ghttp.GetServer().BindHandler("get:/article/:cid/preview/",ArticlePreview)
	ghttp.GetServer().BindHandler("/logout",Logout)
	ghttp.GetServer().BindHandler("post:/comment",Comment)
	ghttp.GetServer().BindHandler("get:/category/:keyword",Categories)
	ghttp.GetServer().BindHandler("get:/category/:keyword/:page",CategoriesByKeywordPage)
	ghttp.GetServer().BindHandler("get:/archives",Archives)
	ghttp.GetServer().BindHandler("get:/links",Links)
	ghttp.GetServer().BindHandler("get:/:pagename",Page)
	ghttp.GetServer().BindHandler("get:/search/:keyword",Search)
	ghttp.GetServer().BindHandler("get:/search/:keyword/:page",SearchPage)
	ghttp.GetServer().BindHandler("get:/tag/:name",Tags)
	ghttp.GetServer().BindHandler("get:/tag/:name/:page",TagsPage)
}

/**
* 首页
* @return
*/
func Index(r *ghttp.Request){
	limitParam := r.GetRequestString("limit")
	var limit = 12
	if limitParam!="" {
		limit,_ = strconv.Atoi(limitParam)
	}
	getContentPages(0,limit,r)
}


/*
*首页分页
* @param request request
* @param pageNo       第几页
* @param limit   每页大小
* @return 主页
 */
func IndexPage(r *ghttp.Request){
	pageNo := r.Get("pageNo")
	limitParam := r.GetRequestString("limit")
	var limit = 12
	if limitParam!="" {
		limit,_ = strconv.Atoi(limitParam)
	}
	p,_ := strconv.Atoi(pageNo)
	getContentPages(p,limit,r)
}

/*
首页内容的分页处理方法
 */
func getContentPages(p,limit int,r *ghttp.Request){
	content, _ := gins.View().Parse("index.tpl", map[string]interface{}{
		"id"   : 123,
		"name" : "john",
	})
	r.Response.Write(content)
}


/*
* 文章页
*
* @param request 请求
* @param cid     文章主键
* @return
 */
func GetArticle(r *ghttp.Request){

}

/*
文章页 预览
* @param request 请求
* @param cid     文章主键
* @return
 */
func ArticlePreview(r *ghttp.Request){

}


/*
登出注销
 */
func Logout(r *ghttp.Request){

}

/*
评论操作
 */
func Comment(r *ghttp.Request){

}

/*
分类页
 */
func Categories(r *ghttp.Request){

}

/*
分类页
 */
func CategoriesByKeywordPage(r *ghttp.Request){

}

/*
归档页
 */
func Archives(r *ghttp.Request){

}

/*
友链页
 */
func Links(r *ghttp.Request){

}

/*
自定义页面,如关于的页面
 */
func Page(r *ghttp.Request){

}

/*
搜索页
 */
func Search(r *ghttp.Request){

}


/*
搜索 分页
 */
func SearchPage(r *ghttp.Request){

}

/*
标签页
 */
func Tags(r *ghttp.Request){

}

/*
标签 分页
 */
func TagsPage(r *ghttp.Request){

}

