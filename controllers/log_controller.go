package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
	"sshDemo/models"
)

type LogController struct {
	beego.Controller
}

// @router /:id      [get]
func (lc *LogController)WriteLog() {
	id := lc.GetString(":id")
	err := models.WriteLog(id)
	if err != nil {
		fmt.Println("write log failed: ",err)
		return
	}
	lc.Data["json"] = nil
	lc.Ctx.Output.SetStatus(http.StatusOK)
	lc.ServeJSON()
}
