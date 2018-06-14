package service

import (
	"github.com/magicgravity/myblog/dao"
	"github.com/golang/glog"
	"gitee.com/johng/gf/g/database/gdb"

	"github.com/magicgravity/myblog/model"
)

func SaveMeta(tx *gdb.Tx,mtype,name string ,mid uint32)bool{
	if len(mtype)>0 && len(name)>0 {
		r := dao.SelectMetaByNameType(name,mtype)
		if r!=nil {
			//已经存在该项
			glog.Error("save meta fail ,reason ===>already exist this meta ")
			return false
		}else{
			if mid >0 {
				//查找以前的 更新
				existOld := dao.SelectMetaByMid(mid)
				if existOld!= nil {
					//确实存在
					ok := dao.UpdateMetaNameByMid(tx,mid,name)
					if ok {
						// 更新原有文章的categories
						UpdateArticleCategory(tx,existOld["name"].(string),name)
						return true
					}else{
						return false
					}
				}
			}

			//插入新的meta
			newMeta := model.Meta{}
			newMeta.Type = mtype
			newMeta.Name = name
			ok,_ :=dao.InsertMetaStruct(tx,newMeta)
			if ok {
				return true
			}else{
				return false
			}
		}
	}
}
