package lights

import (
	"reflect"
	"testing"
)

func TestNewTrafficLight(t *testing.T) {
	got := NewTrafficLight()

	if reflect.ValueOf(got.LastPipeline).Kind() != reflect.Ptr {
		t.Error("should LastPipeline be a pointer")
	}

	if reflect.ValueOf(got.CurrentPipeline).Kind() != reflect.Ptr {
		t.Error("should CurrentPipeline be a pointer")
	}

	if got.CurrentLightState.Green != "off" || got.CurrentLightState.Yellow != "off" || got.CurrentLightState.Red != "off" {
		t.Error("should all lights for CurrentLightState to be off")
	}

	if got.LastLightState.Green != "off" || got.LastLightState.Yellow != "off" || got.LastLightState.Red != "off" {
		t.Error("should all lights for LastLightState to be off")
	}
}

func TestTrafficLight_SetNewLightStatus_Success(t *testing.T) {
	tl := NewTrafficLight()
	tl.SetNewLightStatus(PipelineSucceed)
	if tl.CurrentLightState.Red != LightOff || tl.CurrentLightState.Yellow != LightOff || tl.CurrentLightState.Green != LightOn {
		t.Error("should lights be set as follow: red=off yellow=off green=on")
	}
}

func TestTrafficLight_SetNewLightStatus_Failure(t *testing.T) {
	tl := NewTrafficLight()
	tl.SetNewLightStatus(PipelineFailed)
	if tl.CurrentLightState.Red != LightOn || tl.CurrentLightState.Yellow != LightOff || tl.CurrentLightState.Green != LightOff {
		t.Error("should lights be set as follow: red=on yellow=off green=off")
	}
}

func TestTrafficLight_SetNewLightStatus_Pending(t *testing.T) {
	tl := NewTrafficLight()
	tl.SetNewLightStatus(PipelinePending)
	if tl.CurrentLightState.Red != LightOff || tl.CurrentLightState.Yellow != LightOn || tl.CurrentLightState.Green != LightOff {
		t.Error("should lights be set as follow: red=off yellow=on green=off")
	}
}

func TestTrafficLight_SetNewLightStatus_AllOn(t *testing.T) {
	tl := NewTrafficLight()
	tl.SetNewLightStatus(DeviceError)
	if tl.CurrentLightState.Red != LightOn || tl.CurrentLightState.Yellow != LightOn || tl.CurrentLightState.Green != LightOn {
		t.Error("should lights be set as follow: red=on yellow=on green=on")
	}
}

func TestTrafficLight_SetNewLightStatus_AllOff(t *testing.T) {
	tl := NewTrafficLight()
	tl.SetNewLightStatus(DeviceOff)
	if tl.CurrentLightState.Red != LightOff || tl.CurrentLightState.Yellow != LightOff || tl.CurrentLightState.Green != LightOff {
		t.Error("should lights be set as follow: red=off yellow=off green=off")
	}
}
