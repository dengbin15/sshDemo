package main

import (
	"github.com/astaxie/beego"
	_ "sshDemo/common"
	"sshDemo/models"
	_ "sshDemo/routers"
)

func main() {
	go models.WriteLog()
	beego.Run()
}

