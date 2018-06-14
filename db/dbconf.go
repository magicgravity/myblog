package db

import (
	//"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"sync"
	"fmt"
	"gitee.com/johng/gf/g/os/gcfg"
	"gf/g"
)

var mysqldb *gdb.Db

var mysqlOnce = sync.Once{}

func GetMySqlDbFromCfg() *gdb.Db{
	mysqlOnce.Do(func() {
		g.Config().SetPath("D:\\DevTmp\\")
		g.Config().SetFileName("config.yaml")
		mysqldb = g.Database("default")
	})
	return mysqldb
}

func GetMySqlDb() *gdb.Db{
	mysqlOnce.Do(func(){
		cfg :=gcfg.New("D:\\DevTmp\\","config.yaml")
		cfg.SetFileName("config.yaml")
		cfg.SetPath("D:\\DevTmp\\")
		cfg.Reload()
		cfg.GetString("pass")

		//if err :=g.Config().SetPath("D:\\DevTmp");err!= nil {
		//	fmt.Errorf("set path fail >>> %v",err)
		//}
		//g.Config().SetFileName("config.yaml")

		//gdb.SetDefaultGroup("default")
		gdb.SetConfig(gdb.Config {
			"default" : gdb.ConfigGroup {
				gdb.ConfigNode {
					Host     : "127.0.0.1",
					Port     : "3306",
					User     : "tale",
					Pass     : "123456",
					Name     : "tale",
					Type     : "mysql",
					Role     : "master",
					Priority : 100,
				},
			},
		})
		if db,err := gdb.New("default");err== nil {
			mysqldb = db
		}else{
			fmt.Errorf("create db connection fail ")
		}
		//d :=g.Database("default")
		//fmt.Printf("================> %v ",d)
	})
	return mysqldb
}
