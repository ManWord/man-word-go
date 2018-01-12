package main

import (
	_ "manword/routers"
	"github.com/astaxie/beego"
	_ "manword/utils"
)

func main() {
	beego.Run()
}

