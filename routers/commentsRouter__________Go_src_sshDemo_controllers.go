package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["sshDemo/controllers:LogController"] = append(beego.GlobalControllerRouter["sshDemo/controllers:LogController"],
		beego.ControllerComments{
			Method: "WriteLog",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
