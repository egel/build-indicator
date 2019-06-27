package lights

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/xanzy/go-gitlab"
)

var options = &gitlab.ListProjectPipelinesOptions{
	Scope: gitlab.String("branches"),
	Ref:   gitlab.String(os.Getenv("GITLAB_BRANCH_NAME")),
	Sort:  gitlab.String("desc"),
}

func (tl *TrafficLight) GetLastPipeline(git *gitlab.Client) error {
	pipelines, _, err := git.Pipelines.ListProjectPipelines(os.Getenv("GITLAB_PROJECT_ID"), options)
	if err != nil {
		log.Fatalf("can not get list of pipelines. error: %v", err)
		return err
	}
	tl.LastPipeline = tl.CurrentPipeline
	tl.CurrentPipeline = pipelines[0]
	log.Infof("Previous '%s' pipeline (%d) status %s", tl.LastPipeline.Ref, tl.LastPipeline.ID, tl.LastPipeline.Status)
	log.Infof("Current '%s' pipeline (%d) status %s", tl.CurrentPipeline.Ref, tl.CurrentPipeline.ID, tl.CurrentPipeline.Status)
	return nil
}
