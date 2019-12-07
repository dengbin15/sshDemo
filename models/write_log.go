package models

import (
	"fmt"
	"github.com/hypersleep/easyssh"
	"math/rand"
	"time"
)

func WriteLog(clusterId string) error {
	ssh := &easyssh.MakeConfig {
			User : "root" ,
			Password : "root" ,
			Port : "22" ,
			Server : "10.72.104.188",
	}
	ticker := time.NewTicker(3*time.Second)
	FOR1:
	for {
		select {
		case <- ticker.C :
			var str string
			timeNow := time.Now().Format("2006-01-02 15:04:05")
			randNum := rand.Intn(100)
			if randNum < 50 {
				str = timeNow + " failed=1"
				command := "echo " + str + " >> /var/log/" + clusterId + ".log"
				fmt.Println(command)
				_ , err := ssh.Run(command)
				if err != nil {
					fmt.Println("run command failed: " , err)
					return err
				}
				break FOR1
			} else {
				str = timeNow + " failed=0"
			}
			command := "echo " + str + " >> /var/log/" + clusterId + ".log"
			fmt.Println(command)
			_ , err := ssh.Run(command)
			if err != nil {
				fmt.Println("run command failed: " , err)
				return err
			}
		}
	}
	return nil
}
