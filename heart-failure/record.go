package heartfailure

import "time"

type Record struct {
	ID                   string
	Date                 time.Time
	MorningBloodPressure string
	MorningWeight        float32
	UrinationCount       int8
	Sodium               int16
	Water                int16
	Activity             string
	HeartFailureZone     HeartFailureZone
}
