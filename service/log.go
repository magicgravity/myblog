package service

import (
	"github.com/magicgravity/myblog/dao"
	"github.com/magicgravity/myblog/util"
	"github.com/golang/glog"
	"gitee.com/johng/gf/g/util/gvalid"
	"strconv"
	"time"
	"gitee.com/johng/gf/g/database/gdb"
)

func AddLog(tx *gdb.Tx,action ,data ,ip string,authId uint32)bool{
	params := map[string]interface{} {
		"action"  : action,
		"data"  : data,
		"ip" : ip,
		"authId":authId,
	}
	rules := map[string]string {
		"action"  :  "required",
		"data"  : "required",
		"ip" :  "required",
		"authId" :  "required",
	}
	msgs  := map[string]interface{} {
		"action" : "动作不能为空",
		"data" : "内容不能为空",
		"ip" : "ip不能为空",
		"authId" : "id不能为空",
	}
	verifyRes :=gvalid.CheckMap(params,rules,msgs)
	glog.Info("-------------->",verifyRes)
	if len(verifyRes)>0 {
		return false
	}
	cdate := uint32(0)
	if curtime,err := strconv.Atoi(util.FormatDate(time.Now(),"yyyyMMddhh"));err== nil{
		cdate = uint32(curtime)
	}
	return dao.InsertLog(tx,action,data,ip,authId,cdate)
}