package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/hypersleep/easyssh"
	"strings"
	"time"
)

func InstallCluster() {
	ticker := time.NewTicker(10*time.Second)
	for {
		select {
		case <- ticker.C :
			clusters , err := GetInstallingClusters()
			if err != nil {
				fmt.Println(err)
				return
			}
			for _ , cluster := range clusters {
				go ReadLog(cluster)
			}
		}
	}
}

func GetInstallingClusters() ([]Cluster , error) {
	var clusters []Cluster
	o := orm.NewOrm()
	o.Using("default")
	_ , err := o.QueryTable("Cluster").Filter("status" , "installing").All(&clusters)
	if err != nil {
		fmt.Println("get time out cluster failed: " , err)
		return nil , err
	}
	return clusters , nil
}

/*func GetTimeoutCluster(clusters []Cluster) ([]Cluster , error) {
	var timeoutCluster []Cluster
	for _ , cluster := range clusters {
		startTime , err :=
		if err != nil {
			fmt.Println("get time out cluster failed: " , err)
			return nil , err
		}
	}
}*/

func ReadLog(cluster Cluster){
	ssh := &easyssh.MakeConfig{
		User : "root" ,
		Password : "root" ,
		Server : "10.72.104.188" ,
		Port : "22" ,
	}
	command := "tail -f /var/log/" + cluster.Id + ".log"
	ticker := time.NewTicker(3*time.Second)
	ch , done , err := ssh.Stream(command)
	if err != nil {
		fmt.Println(err)
		return
	}
FOR1:
	for {
		select {
		case <- ticker.C :
			select {
			case <-done:
				break FOR1
			case line := <-ch:
				fmt.Println(cluster.Id + " log is: " + line)
				if strings.Contains(line , "failed=1") {
					fmt.Println("read log failed=1")
					UpdateClusterStatus(cluster.Id)
					break FOR1
				}
			}
		}
	}
}

func UpdateClusterStatus(clusterId string) error {
	o := orm.NewOrm()
	sql := `update cluster set status="failed" where id=`+ clusterId + ";"
	fmt.Println(sql)
	_ , err := o.Raw(sql).Exec()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
