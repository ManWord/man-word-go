package routers

import (
	"github.com/astaxie/beego"
	"manword/controllers/sample"
	"manword/controllers/login"
)

func init() {

	//登录测试
	beego.Router("/login", &login.LoginController{}, "*:Login")

	//这种形式的路由注册会根据http method类型自动匹配
	beego.Router("/sample/?:username:string", &sample.SampleController{})
	//这种形式的路由注册，指定了请求交由哪个方法进行处理
	beego.Router("/sample/sayhi", &sample.SampleController{}, "*:SayHi")

	beego.Router("/sample/insert", &sample.SampleController{}, "*:InsertData")

	beego.Router("/sample/query/?:id:int", &sample.SampleController{}, "*:QueryData")

	///////////////////////////////////////以上为测试路由///////////////////////////////////
}
