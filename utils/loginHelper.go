package utils

import (
	"github.com/astaxie/beego"
)

//be careful, this method need not expose to user
//here just a test for sample.go  .
//actually, wo will expose a few useful interface.
//such as username, nickname and ...
func GetSessionInfo(c beego.Controller) SessionInfo {
	if !IsLogin(c) {
		panic("error: user have not login!")
	}
	return c.GetSession(c.Ctx.GetCookie(SESSION_ID)).(SessionInfo)
}

func IsLogin(c beego.Controller) bool {
	return c.GetSession(c.Ctx.GetCookie(SESSION_ID)) != nil
}
