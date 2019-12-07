package routers

import (
	"sshDemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    log := beego.NewNamespace("/ssh" ,
    	beego.NSNamespace("log",beego.NSInclude(&controllers.LogController{})))
	beego.AddNamespace(log)
}
