package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net/http"
	"sshDemo/models"
	"time"
)

type LogController struct {
	beego.Controller
}

// @router /:id      [get]
func (c *LogController)WriteLog() {
	id := c.GetString(":id")
	go models.WriteLog(id)
	timeNow := time.Now()
	//instanceId := fmt.Sprintf("%s",uuid.NewV4())
	cluster := models.Cluster{
		Id : id ,
		Status : "installing",
		CreateTime : timeNow ,
	}
	err := InsertCluster(cluster)
	if err != nil {
		return
	}
	c.Data["json"] = cluster
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.ServeJSON()
}

func InsertCluster(cluster models.Cluster) error {
	o := orm.NewOrm()
	_ , err := o.Insert(&cluster)
	if err != nil {
		fmt.Println("insert cluster failed: ", err)
		return err
	}
	return nil
}