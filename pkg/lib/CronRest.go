package lib

import (
	"github.com/RogerDurdn/MonitoringApp/pkg/http"
	"github.com/robfig/cron/v3"
	"log"
)

var c *cron.Cron

func InitCron()  {
	log.Println("init cron")
	c = cron.New()
	_, err := c.AddFunc("@every 2s", func() {
		obtainData(&http.MockSource{})
	})
	if err != nil {
		log.Panic("Failed to initialize cron: ", err)
	}
	c.Start()
}
func StopCron() {
	c.Stop()
}

func obtainData(dr http.DataSource){
	log.Println("executing change data")
	data := dr.FetchData()
	ChangeData(data)
}
