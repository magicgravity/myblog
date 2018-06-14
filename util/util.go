package util

import (
	"crypto/md5"
	"encoding/hex"
	"time"
	"strings"
	"bytes"
	"github.com/magicgravity/myblog/common"
	"strconv"
)

func MD5encode(str string)string{
	if len(str)<=0{
		return ""
	}else{
		md5 := md5.New()
		_,err :=md5.Write(([]byte)(str))
		if err!=nil{
			return ""
		}else{
			cipherStr :=md5.Sum(nil)
			return hex.EncodeToString(cipherStr)
		}
	}
}


func FormatCurrentDateYYYYMMdd()string{
	now := time.Now()
	return now.Format("20060102150405")
}


/*
格式化时间
 */
func FormatDate(date time.Time,fmtype string ) string{
	uperType := strings.ToUpper(fmtype)
	fmtStr := date.Format("20060102150405")

	buf :=bytes.Buffer{}
	if strings.Contains(uperType,"YYYY") {
		buf.WriteString(strings.Replace(uperType,"YYYY",fmtStr[0:4],-1))
	}else{
		buf.WriteString(uperType)
	}
	if strings.Contains(buf.String(),"MM") {
		month :=strings.Replace(buf.String(),"MM",fmtStr[4:6],-1)
		buf.Reset()
		buf.WriteString(month)
	}

	if strings.Contains(buf.String(),"DD") {
		day := strings.Replace(buf.String(),"DD",fmtStr[6:8],-1)
		buf.Reset()
		buf.WriteString(day)
	}

	if strings.Contains(buf.String(),"HH") {
		hour := strings.Replace(buf.String(),"HH",fmtStr[8:10],-1)
		buf.Reset()
		buf.WriteString(hour)
	}

	if strings.Contains(buf.String(),"MI") {
		minute := strings.Replace(buf.String(),"MI",fmtStr[10:12],-1)
		buf.Reset()
		buf.WriteString(minute)
	}

	if strings.Contains(buf.String(),"SS") {
		second := strings.Replace(buf.String(),"SS",fmtStr[12:],-1)
		buf.Reset()
		buf.WriteString(second)
	}

	return buf.String()
}


func IsPath(p string)bool{
	if len(p)>0{
		if strings.Contains(p,"/") || strings.Contains(p," ") || strings.Contains(p,".") {
			return false
		}else{
			if ps := common.SLUG_REGEX.Find([]byte(p));ps == nil {
				return false
			}else{
				return true
			}
		}
	}else{
		return false
	}

}


func GetCurTimeAsInt()uint32{
	if curtime,err := strconv.Atoi(FormatDate(time.Now(),"yyyyMMddhh"));err== nil{
		return uint32(curtime)
	}else{
		return 0
	}
}