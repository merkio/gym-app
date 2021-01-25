package tasks

import (
	config "gym-app/app-config"
	"gym-app/app/program"
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
	s, err := scheduler.NewScheduler(1)

	if err != nil {
		log.Error(err) // just example
	}

	s.Every().Minute(29).Hour(19).Do(vkCollectorTask, "", 10, 0)
}

func vkCollectorTask(query string, count, offset int) {
	client := newClient()
	response := WallResponseData{}

	for i := offset; i < offset+100; i = i + count {
		time.Sleep(20 * time.Second)
		log.Infof("Request with count %v and offset %v", count, i)
		if err := client.CallMethod("wall.get", vk.RequestParams{
			//"query":   query,
			"count":    count,
			"offset":   i,
			"owner_id": config.VkConnectionConfig.GroupID,
		}, &response); err != nil {
			log.Error("Request to the vk api failed", err)
		}
		programRepo := program.PRepository{}
		if len(response.Items) <= 0 {
			return
		}
		for _, post := range response.Items {
			if programRepo.GetByText(post.Text) {
				log.Errorf("Program already exist for date %v\n%s", post.ID, time.Unix(post.Date, 0).String())
				continue
			}
			str, err := programRepo.Create(program.Program{
				Text:    post.Text,
				DateInt: post.Date,
				Tags:    "raw,post",
				Date:    time.Unix(post.Date, 0),
			})
			if err != nil {
				log.Error("Error during save the program", err)
			}
			log.Infof("Saved new program with ID %s and Date %s", str, time.Unix(post.Date, 0).String())
		}
	}
}
