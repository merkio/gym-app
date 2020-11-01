package tasks

import (
	_ "github.com/go-vk-api/vk"
	"github.com/prprprus/scheduler"
)

func collectNewMesssages() {
	s, err := scheduler.NewScheduler(1000)

	if err != nil {
		log.Error(err) // just example
	}

	s.Every().Second(45).Minute(20).Hour(13).Day(23).Weekday(3).Month(6).Do(vkCollectorTask, "prprprus", 23)
}

func vkCollectorTask() {

}
