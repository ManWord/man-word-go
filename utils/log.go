package utils

/**
全局日志工具
 */
import (
	"github.com/astaxie/beego/logs"
)

func init() {
	//日志设置，日志便于调试，鉴于服务器本身容量问题，本项目不建议保存到文件中，要不然会很快
	//挤爆服务器~，故这设置控制台输出，调试信息可以在console中看到
	logs.SetLogger("console")
}
