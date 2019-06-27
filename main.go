package main

import (
	"os"
	"time"

	"github.com/egel/build-indicator/lights"
	log "github.com/sirupsen/logrus"
	"github.com/xanzy/go-gitlab"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	trafficLight := lights.NewTrafficLight()

	git := gitlab.NewClient(nil, os.Getenv("GITLAB_ACCESS_TOKEN"))
	if err := git.SetBaseURL(os.Getenv("GITLAB_BASEURL")); err != nil {
		log.Fatalf("Can not set %s BaseURL. Error: %v", "Gitlab", err)
	}

	for {
		if err := trafficLight.GetLastPipeline(git); err != nil {
			log.Error("can not get latest pipeline status", err)
			trafficLight.SetNewLightStatus(lights.DeviceError)
		} else {
			switch trafficLight.LastPipeline.Status {
			case string(gitlab.Pending), string(gitlab.Running):
				trafficLight.SetNewLightStatus(lights.PipelinePending)
			case string(gitlab.Success):
				trafficLight.SetNewLightStatus(lights.PipelineSucceed)
			case string(gitlab.Canceled), string(gitlab.Failed), string(gitlab.Skipped):
				trafficLight.SetNewLightStatus(lights.PipelineFailed)
			default:
				trafficLight.SetNewLightStatus(lights.DeviceError)
			}
		}

		time.Sleep(time.Second * 10)
	}
}
