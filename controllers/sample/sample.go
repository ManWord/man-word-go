package sample

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	orm2 "github.com/astaxie/beego/orm"
	"manword/models/sample"
	"strconv"
	"manword/utils"
)

/**
   这是个样例，非业务接口
 */
type SampleController struct {
	beego.Controller
}

//beego.Controller 提供了符合restful规范的一系列方法，会自动根据http中的method进行匹配
//这里复写Get方法,直接调用post进行处理
func (this *SampleController) Get() {
	//Info日志输出
	logs.Info("get method invoke...")

	username := this.Ctx.Input.Param(":username")

	test := this.GetSession(username)

	//
	logs.Info(username, test)

	this.CruSession.Set(username, username)

	if !utils.IsLogin(this.CruSession, username) {

		//登录不应太在这里调用，这个只是测试，说明登录功能可用
		//此时如果没有login，返回给前端数据即可，有前端来redirect
		utils.Login(this.CruSession, username)
	}

	this.Post();
}

//复写Post方法
func (this *SampleController) Post() {
	//Debug日志输出
	logs.Debug("post method invoke...")
	this.Ctx.WriteString(beego.AppConfig.String("jdbc.username"))

}

//自定义业务处理方法
func (this *SampleController) SayHi() {
	//错误日志输出，这里仅仅是测试
	logs.Error("It's a test error message, you can delete it ...")
	this.Ctx.WriteString("I am SayHi Method")
}

func (this *SampleController) InsertData() {
	orm := orm2.NewOrm()
	orm.Using("default")

	sampleModel := new(sample.SampleModel)
	sampleModel.Content = "Hi I am Sample Test";
	if _, err := orm.Insert(sampleModel); err != nil {
		this.Ctx.WriteString("insert failed...")
	}

	this.Ctx.WriteString("insert Success...")
}

func (this *SampleController) QueryData() {
	orm := orm2.NewOrm()

	var queryId int
	queryId, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil {
		this.Ctx.WriteString("invalid id...")
	}
	sampleModel := sample.SampleModel{Id: queryId}

	err2 := orm.Read(&sampleModel)

	if err2 == orm2.ErrNoRows {
		logs.Error("查询不到")
		this.Ctx.WriteString("can not query anything...")
	} else if err2 == orm2.ErrMissPK {
		logs.Error("找不到主键")
		this.Ctx.WriteString("no primary key...")
	} else {
		logs.Info("--------------------")
		logs.Info(sampleModel.Id, sampleModel.Content)
		logs.Info("--------------------")
	}

	this.Ctx.WriteString("query Success...")

}
