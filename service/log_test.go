package service

import "testing"

func TestAddLog(t *testing.T) {
	ok := AddLog(nil,"添加测试","测试内容222","127.0.0.1",1)
	if ok {
		t.Log("增加日志成功")
	}else{
		t.Error("增加日志失败")
	}
}
