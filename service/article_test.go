package service

import (
	"testing"
	"github.com/magicgravity/myblog/model"
	"strconv"
)

func TestPublishArticle(t *testing.T) {
	testCnt := 100
	tagArrs := []string{"spring","c","c++","javascript","python","go","php","sql"}
	for testCnt>0 {

		c := model.Contents{}
		c.Tags = tagArrs[testCnt%8]+strconv.Itoa(testCnt)
		c.Categories = "开发技术-"+strconv.Itoa(testCnt%8)+"-"+strconv.Itoa(testCnt)
		c.Type = "post"
		c.Slug = "about"+strconv.Itoa(testCnt)
		c.Status = "publish"
		c.Title = "测试文章标题---"+strconv.Itoa(testCnt)
		c.Content = []byte("测试内容>>"+tagArrs[testCnt%8])
		testCnt--

		ok,cid := PublishArticle(c)
		if ok{
			t.Logf("publish article successful ! cid == > %v \r\n",cid)
		}else{
			t.Error("publish article fail !")
		}
	}
}


func TestGetArticleContents(t *testing.T) {
	var totalPages uint64= 21
	var curPage uint64 = 1
	for curPage<=totalPages {
		page := GetArticleContents(curPage, 5)
		if len(page.GetList()) > 0 {
			for idx, val := range page.GetList() {
				t.Logf("[%d]>>>>[%d]  -----value===> %v", curPage,idx, val["cid"])
			}
			t.Logf(page.ToString())
		}

		t.Log("~~~~~~~~~~~~~~~~~~~~~~~")
		curPage++
	}

	//page2 := GetArticleContents(2,5)
	//if len(page2.GetList())>0 {
	//	for idx,val := range page2.GetList(){
	//		t.Logf("[%d]  -----value===> %v",idx,val["cid"])
	//	}
	//	t.Logf(page2.ToString())
	//}
}


func TestGetArticlesByKeyword(t *testing.T){
	keywordArr := []string{"java","开发","技术难度","技术","测试","试标题"}
	for idx,key := range keywordArr {
		list := GetArticlesByKeyword(key)
		if list!= nil && len(list)>0 {
			t.Logf("[%d]  find articles ok! got %d records ,keyword => %s",idx,len(list),key)
		}else{
			t.Logf("[%d]  find articles fail! got 0 records ,keyword => %s",idx,key)
		}
	}
}


