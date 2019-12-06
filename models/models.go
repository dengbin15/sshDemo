package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
	"time"
)

type Cluster struct {
	id      string		`orm:"pk;size(36);column(id)" json:"id"`
	name    string 		`orm:"size(36);column(name)" json:"name"`
	status  string		`orm:"size(36);column(status)" json:"status"`
}

type Object struct {
	Id		string		`orm:"pk;size(36);column(id)" json:"id"`
}

type CsfObject struct {
	Id            string     `orm:"pk;size(36);column(id)" json:"id"`
	Operation     string     `orm:"size(36);column(operation)"`
	CreateTime    time.Time	 `orm"type(datetime)"`
}

func OrmInit() {
	orm.RegisterModel("",new(Cluster) , new(Object) , new(CsfObject))
	orm.RunSyncdb("default", false , true)
}
