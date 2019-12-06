package common

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"sshDemo/models"
)

var CommonSshPath string

func init() {
	// 注册数据库
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbUser := "root"
	dbPassword := "root"
	database := "ssh_demo"
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8" , dbUser , dbPassword , dbHost , dbPort , database )
	err := orm.RegisterDataBase("default" , "mysql" , dbLink , 30)
	if err != nil {
		fmt.Println("register database failed: ",err)
		return
	}
	CommonSshPath = beego.AppConfig.DefaultString("common_Ssh_log_path","/var/log/result.log")
	models.OrmInit()
}
