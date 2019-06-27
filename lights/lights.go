package lights

const (
	LightOff string = "off"
	LightOn  string = "on"
)

type Light struct {
	Red    string
	Yellow string
	Green  string
}

func GetRedOnlySolid() *Light {
	return &Light{
		Red:    LightOn,
		Yellow: LightOff,
		Green:  LightOff,
	}
}

func GetYellowOnlySolid() *Light {
	return &Light{
		Red:    LightOff,
		Yellow: LightOn,
		Green:  LightOff,
	}
}

func GetGreenOnlySolid() *Light {
	return &Light{
		Red:    LightOff,
		Yellow: LightOff,
		Green:  LightOn,
	}
}

func GetAllOn() *Light {
	return &Light{
		Red:    LightOn,
		Yellow: LightOn,
		Green:  LightOn,
	}
}

func GetAllOff() *Light {
	return &Light{
		Red:    LightOff,
		Yellow: LightOff,
		Green:  LightOff,
	}
}
