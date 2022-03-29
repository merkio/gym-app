package tasks

import (
	config "gym-app/app-config"
	"gym-app/app/model"
	"gym-app/app/program"
	"gym-app/common/db"
	"time"

	"github.com/go-vk-api/vk"
	"github.com/prprprus/scheduler"
)

func newClient() *vk.Client {
	client, err := vk.NewClientWithOptions(
		vk.WithToken(config.VkConnectionConfig.AccessToken),
	)
	if err != nil {
		log.Error("Failed to create vk client", err)
	}
	return client
}

// CollectVkMessages run task to collect messages from vk
func CollectVkMessages() {
	log.Info("Start vk task")
	s, err := scheduler.NewScheduler(1)

	if err != nil {
		log.Error(err) // just example
	}
	s.Every().
		Hour(config.VkConnectionConfig.Hour).
		Minute(config.VkConnectionConfig.Minute).
		Do(vkCollectorTaskForGroups, 10, 0)
}

func vkCollectorTaskForGroups(count, offset int) {
	for k, v := range config.VkConnectionConfig.GetGroups() {
		vkCollectorTask(k, v, count, offset)
	}
}

func vkCollectorTask(group_name, group_id string, count, offset int) {
	log.Info("Collecting data from vk")
	client := newClient()
	response := WallResponseData{}
	countErrors := 0
	i := offset
	for {
		time.Sleep(20 * time.Second)
		log.Infof("Request with count %v and offset %v and group %v", count, i, group_name)
		if err := client.CallMethod("wall.get", vk.RequestParams{
			"count":    count,
			"offset":   i,
			"owner_id": group_id,
		}, &response); err != nil {
			log.Error("Request to the vk api failed", err)
		}
		programRepo := program.NewPRepository(db.GetDB(config.DataConnectionConfig), log)
		if len(response.Items) <= 0 {
			return
		}
		for _, post := range response.Items {
			if programRepo.GetByText(post.Text) {
				log.Errorf("Program already exist for date %v\n%s", post.ID, time.Unix(post.Date, 0).String())
				countErrors += 1
				continue
			}
			str, err := programRepo.Create(model.Program{
				Text:      post.Text,
				DateInt:   post.Date,
				Tags:      "raw,post",
				GroupName: group_name,
				GroupID:   group_id,
				Date:      time.Unix(post.Date, 0),
			})
			if err != nil {
				log.Error("Error during save the program", err)
			}
			log.Infof("Saved new program with ID %s and StartDate %s", str, time.Unix(post.Date, 0).String())
		}
		i = i + count
		if countErrors > 30 {
			return
		}
	}
}
