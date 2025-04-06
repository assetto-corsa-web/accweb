package instance

import "github.com/assetto-corsa-web/accweb/internal/pkg/event"

type LapState struct {
	CarID       int          `json:"carID"`
	DriverIndex int          `json:"driverIndex"`
	Car         *CarState    `json:"-"`
	Driver      *DriverState `json:"-"`
	LapTimeMS   int          `json:"lapTimeMS"`
	TimestampMS int          `json:"timestampMS"`
	Flags       int          `json:"flags"`
	S1          string       `json:"s1"`
	S1MS        int          `json:"s1MS"`
	S2          string       `json:"s2"`
	S2MS        int          `json:"s2MS"`
	S3          string       `json:"s3"`
	S3MS        int          `json:"s3MS"`
	Fuel        int          `json:"fuel"`
	HasCut      bool         `json:"hasCut"`
	InLap       bool         `json:"inLap"`
	OutLap      bool         `json:"outLap"`
	SessionOver bool         `json:"sessionOver"`
}

func (l LapState) IsValid() bool {
	return l.Flags == 0 && !l.HasCut && !l.InLap && !l.OutLap && !l.SessionOver
}

func (l LapState) ToEILS() event.EventLapState {
	return event.EventLapState{
		DriverIndex: l.DriverIndex,
		LapTimeMS:   l.LapTimeMS,
		TimestampMS: l.TimestampMS,
		Flags:       l.Flags,
		S1:          l.S1,
		S1MS:        l.S1MS,
		S2:          l.S2,
		S2MS:        l.S2MS,
		S3:          l.S3,
		S3MS:        l.S3MS,
		Fuel:        l.Fuel,
		HasCut:      l.HasCut,
		InLap:       l.InLap,
		OutLap:      l.OutLap,
		SessionOver: l.SessionOver,
	}
}
