package utils

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/session"
)

func IsLogin(cruSession session.Store, username string) bool {
	//TODO 1、解析客户端数据，查询数据库是否存在
	//TODO 2、根据username保存session信息
	logs.Info(cruSession.Get(username), "=====Islogin")
	return cruSession.Get(username) != nil
}

func Login(cruSession session.Store, uniquesign string) bool {
	signval := cruSession.Get(uniquesign);
	if signval == nil {
		cruSession.Set(uniquesign, signval)
	}
	logs.Info(signval == nil)
	return true
}
