package lights

import (
	"bytes"
	"os/exec"

	cc "github.com/egel/build-indicator/clewarecontrol"
	log "github.com/sirupsen/logrus"
	"github.com/xanzy/go-gitlab"
)

// LightType represents a Mini Traffic Light status.
type LightType string

// These constants represent all valid Mini Traffic Light status.
const (
	DeviceOff       LightType = "device-off"
	DeviceError     LightType = "device-error"
	PipelinePending LightType = "pipeline-pending"
	PipelineSucceed LightType = "pipeline-succeed"
	PipelineFailed  LightType = "pipeline-failed"
)

type TrafficLight struct {
	LastPipeline      *gitlab.PipelineInfo
	CurrentPipeline   *gitlab.PipelineInfo
	LastLightState    *Light
	CurrentLightState *Light
}

func NewTrafficLight() *TrafficLight {
	return &TrafficLight{
		LastPipeline:      &gitlab.PipelineInfo{},
		CurrentPipeline:   &gitlab.PipelineInfo{},
		LastLightState:    GetAllOff(),
		CurrentLightState: GetAllOff(),
	}
}

func (tl *TrafficLight) SetNewLightStatus(lightType LightType) {
	tl.LastLightState = tl.CurrentLightState
	switch lightType {
	case PipelineSucceed:
		tl.CurrentLightState = GetGreenOnlySolid()
	case PipelineFailed:
		tl.CurrentLightState = GetRedOnlySolid()
	case PipelinePending:
		tl.CurrentLightState = GetYellowOnlySolid()
	case DeviceError:
		tl.CurrentLightState = GetAllOn()
	case DeviceOff:
		tl.CurrentLightState = GetAllOff()
	}
	tl.applyNewSetup()
}

func (tl *TrafficLight) applyNewSetup() {
	var programArgs []string
	programArgs = append(programArgs, cc.RepeatSingle...)
	programArgs = append(programArgs, tl.convertCurrentLightStateToProgramArgs()...)

	cmd := exec.Command(cc.Program, programArgs...)
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Errorf("Error: %s, Stderr: %s", err, stderr.String())
		return
	}
	log.Infof("Result: %s", out.String())
}

func (tl *TrafficLight) convertCurrentLightStateToProgramArgs() []string {
	var args []string
	switch tl.CurrentLightState.Red {
	case LightOff:
		args = append(args, cc.RedLightOff...)
	case LightOn:
		args = append(args, cc.RedLightOn...)
	}

	switch tl.CurrentLightState.Yellow {
	case LightOff:
		args = append(args, cc.YellowLightOff...)
	case LightOn:
		args = append(args, cc.YellowLightOn...)
	}

	switch tl.CurrentLightState.Green {
	case LightOff:
		args = append(args, cc.GreenLightOff...)
	case LightOn:
		args = append(args, cc.GreenLightOn...)
	}

	return args
}
