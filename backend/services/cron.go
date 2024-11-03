package services

import (
	"fmt"
	"time"
)

func RunCron() {
	for {
		fmt.Println("running cron")
		go RunApplets()
		time.Sleep(60 * time.Second)
	}
}
