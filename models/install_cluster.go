package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/hypersleep/easyssh"
	"golang.org/x/tools/go/ssa/interp/testdata/src/strings"
	"time"
)

func InstallCluster() {
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <- ticker.C :
			clusters , err := GetInstallingClusters()
			if err != nil {
				fmt.Println(err)
				return
			}
			for _ , cluster := range clusters {
				flag , err := ReadLog(cluster)
				if err != nil {
					fmt.Println(err)
					return
				}
				if flag == false {
					UpdateClusterStatus(cluster.id)
				}
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

func ReadLog(cluster Cluster) (bool , error){
	ssh := &easyssh.MakeConfig{
		User : "root" ,
		Password : "root" ,
		Server : "10.72.104.188" ,
		Port : "22" ,
	}
	command := "tail -f /var/log/" + cluster.id + ".log"
	ticker := time.NewTicker(3*time.Second)
	var flag = true
	for {
		select {
		case <- ticker.C :
			ch , done , err := ssh.Stream(command)
			if err != nil {
				fmt.Println(err)
				return false , err
			}
			select {
			case <-done:
				break
			case line := <-ch:
				if strings.Contains(line , "failed=") {
					done <- true
					if strings.Contains(line , "failed=1") {
						flag = false
					}
				}
			}
		}
	}
	return flag , nil
}

func UpdateClusterStatus(clusterId string) error {
	o := orm.NewOrm()
	sql := "update status failed from cluster where id=" + clusterId + ";"
	_ , err := o.Raw(sql).Exec()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
