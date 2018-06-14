package util

import (
	"testing"
	"time"
)

func TestMD5encode(t *testing.T) {
	ret := MD5encode("admin123456")
	t.Logf("md5 ==> %s",ret)
}


func TestFormatDate(t *testing.T) {
	str := FormatDate(time.Now(),"yyyy-MM-dd")
	t.Logf("curDay ==> %s ",str)

	str = FormatDate(time.Now(),"yyyy-MM-dd hh:mi:ss")
	t.Logf("curDay ==> %s ",str)

	str = FormatDate(time.Now(),"yyyyMMddhhmiss")
	t.Logf("curDay ==> %s ",str)
}