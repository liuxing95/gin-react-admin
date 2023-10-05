package initialize

import (
	"fmt"

	"github.com/liuxing95/gin-react-admin/config"
	"github.com/liuxing95/gin-react-admin/global"
	"github.com/liuxing95/gin-react-admin/utils"
	"github.com/robfig/cron/v3"
)

func Timer() {
	if global.GRA_CONFIG.Timer.Start {
		for i := range global.GRA_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				var option []cron.Option
				if global.GRA_CONFIG.Timer.WithSeconds {
					option = append(option, cron.WithSeconds())
				}
				_, err := global.GRA_Timer.AddTaskByFunc("ClearDB", global.GRA_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.GRA_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				}, option...)
				if err != nil {
					fmt.Println("add timer error:", err)
				}
			}(global.GRA_CONFIG.Timer.Detail[i])
		}
	}
}
