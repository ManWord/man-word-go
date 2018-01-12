package utils

import "github.com/astaxie/beego"

//对于必须要下载的静态文件，可以在此注册下载路径
//不建议在应用程序中加入静态文件，除非是js、css之类的
//绝大多数静态文件直接放到服务器上就行，不必放到工程文件里面
func init() {
	beego.SetStaticPath("/download", "static/download");
}
