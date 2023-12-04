package config

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
)

func InitTask() {
	cronJob := cron.New()
	//分(0-59) 时(0-23) 天(1-31) 月(1-12) 星期(0-6)
	times := fmt.Sprintf("*/%v * * * *", GlobalConfig.AuthorizationFresh)
	// ----------------------------- 定时刷新本服务的accessToken ------------------------------
	if _, err := cronJob.AddFunc(times, func() {
		if err := Authorization(); err != nil {
			fmt.Errorf("定时刷新accessToken失败!", err)
		}
	}); err != nil {
		log.Fatal("定时刷新accessToken定时任务增加失败!", err)
	}
	cronJob.Start()
}
