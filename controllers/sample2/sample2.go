package sample

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

/**
   这个样例存在的目的是呈现session共享，即一个controller设置了session信息后，
所有controller都能得到该信息
 */
type Sample2Controller struct {
	beego.Controller
}

//beego.Controller 提供了符合restful规范的一系列方法，会自动根据http中的method进行匹配
//这里复写Get方法,直接调用post进行处理
func (this *Sample2Controller) Get() {
	//Info日志输出
	logs.Info("sample2 get method invoke...")

	username := this.Ctx.Input.Param(":username")

	status := this.GetSession(username)

	if status == nil {
		this.Ctx.WriteString("have not login!")
		return
	}

	this.Ctx.WriteString("have login!")
}
