package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Cluster struct {
	Id      string		`orm:"pk;size(36);column(id)" json:"id"`
	Name    string 		`orm:"size(36);column(name)" json:"name"`
	Status  string		`orm:"size(36);column(status)" json:"status"`
	CreateTime time.Time `orm:"type(datatime)" json:"createTime"`
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
	orm.RegisterModelWithPrefix("",new(Cluster) , new(Object) , new(CsfObject),)
	orm.RunSyncdb("default", false , true)
}
