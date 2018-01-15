package login

import (
	"github.com/astaxie/beego"
	"manword/utils"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Login() bool {

	r := this.Ctx.Request
	r.ParseForm()

	//TODO 1、query database first, if user not register , notify client
	//here， just a test case , so there is a sample solution
	username := ""
	password := ""

	if len(r.Form["username"]) != 0 {
		username = r.Form["username"][0]
	}

	if len(r.Form["password"]) != 0 {
		password = r.Form["password"][0]
	}

	if username == "" || password == "" {
		panic("error: username or password wrong")
	}

	//save user info
	userInfo := utils.UserInfo{
		Username: username,
		Password: password,
	}

	sessionInfo := utils.SessionInfo{
		userInfo,
		utils.CartInfo{},
	}

	this.SessionRegenerateID()
	this.SetSession(this.CruSession.SessionID(), sessionInfo)

	this.Ctx.WriteString("login success!")

	return true
}
