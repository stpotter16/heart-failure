package heartfailure

type HeartFailureZone int

const (
	Green HeartFailureZone = iota
	Yellow
	Red
)

var heartFailureZoneNameMap = map[HeartFailureZone]string{
	Green:  "Green",
	Yellow: "Yellow",
	Red:    "Red",
}

func (hfz HeartFailureZone) String() string {
	return heartFailureZoneNameMap[hfz]
}
