package event

type EventInstanceLiveDriverBase struct {
	Name     string `json:"playerName"`
	PlayerID string `json:"playerID"`
}

func NewEventInstanceLiveDriver(name, id string) EventInstanceLiveDriverBase {
	return EventInstanceLiveDriverBase{
		Name:     name,
		PlayerID: id,
	}
}

type EventInstanceLiveCarBase struct {
	CarID      int `json:"carID"`
	RaceNumber int `json:"raceNumber"`
	CarModel   int `json:"carModel"`
}

func NewEventInstanceLiveCarBase(cId, rn, cm int) EventInstanceLiveCarBase {
	return EventInstanceLiveCarBase{
		CarID:      cId,
		RaceNumber: rn,
		CarModel:   cm,
	}
}

type EventInstanceLiveNewDriver struct {
	EventBase
	EventInstanceBase
	EventInstanceLiveDriverBase
	EventInstanceLiveCarBase
}

func EmmitEventInstanceLiveNewDriver(eib EventInstanceBase, name, pId string, cId, rn, cm int) {
	Emmit(EventInstanceLiveNewDriver{
		EventBase:                   eventBase("instance_live_new_driver"),
		EventInstanceBase:           eib,
		EventInstanceLiveDriverBase: NewEventInstanceLiveDriver(name, pId),
		EventInstanceLiveCarBase:    NewEventInstanceLiveCarBase(cId, rn, cm),
	})
}

type EventInstanceLiveRemoveConnection struct {
	EventBase
	EventInstanceBase
	EventInstanceLiveDriverBase
}

func EmmitEventInstanceLiveRemoveConnection(eib EventInstanceBase, name, pId string) {
	Emmit(EventInstanceLiveRemoveConnection{
		EventBase:                   eventBase("instance_live_remove_connection"),
		EventInstanceBase:           eib,
		EventInstanceLiveDriverBase: NewEventInstanceLiveDriver(name, pId),
	})
}

type EventInstanceLiveNewLap struct {
	EventBase
	EventInstanceBase
	EventInstanceLiveDriverBase
	EventInstanceLiveCarBase

	LapTimeMS   int    `json:"lapTimeMS"`
	TimestampMS int    `json:"timestampMS"`
	Flags       int    `json:"flags"`
	Fuel        int    `json:"fuel"`
	S1          string `json:"s1"`
	S2          string `json:"s2"`
	S3          string `json:"s3"`
	HasCut      bool   `json:"hasCut"`
	InLap       bool   `json:"inLap"`
	OutLap      bool   `json:"outLap"`
	SessionOver bool   `json:"sessionOver"`
}

func EmmitEventInstanceLiveNewLap(
	eib EventInstanceBase,
	eildb EventInstanceLiveDriverBase,
	eilcb EventInstanceLiveCarBase,
	ltms, tms, flags, fuel int,
	s1, s2, s3 string,
	fhc, fil, fol, fso bool,
) {
	Emmit(EventInstanceLiveNewLap{
		EventBase:                   eventBase("instance_live_new_lap"),
		EventInstanceBase:           eib,
		EventInstanceLiveDriverBase: eildb,
		EventInstanceLiveCarBase:    eilcb,
		LapTimeMS:                   ltms,
		TimestampMS:                 tms,
		Flags:                       flags,
		Fuel:                        fuel,
		S1:                          s1,
		S2:                          s2,
		S3:                          s3,
		HasCut:                      fhc,
		InLap:                       fil,
		OutLap:                      fol,
		SessionOver:                 fso,
	})
}

type EventInstanceLiveSessionPhaseChanged struct {
	EventBase
	EventInstanceBase

	Session   string `json:"session"`
	Phase     string `json:"phase"`
	Remaining int    `json:"remaining"`
}

func EmmitEventInstanceLiveSessionPhaseChanged(eib EventInstanceBase, s, p string, r int) {
	Emmit(EventInstanceLiveSessionPhaseChanged{
		EventBase:         eventBase("instance_live_session_phase_changed"),
		EventInstanceBase: eib,
		Session:           s,
		Phase:             p,
		Remaining:         r,
	})
}

type EventInstanceLiveNewDamageZone struct {
	EventBase
	EventInstanceBase
	EventInstanceLiveDriverBase
	EventInstanceLiveCarBase
}

func EmmitEventInstanceLiveNewDamageZone(
	eib EventInstanceBase,
	eildb EventInstanceLiveDriverBase,
	eilcb EventInstanceLiveCarBase,
) {
	Emmit(EventInstanceLiveNewDamageZone{
		EventBase:                   eventBase("instance_live_new_damage_zone"),
		EventInstanceBase:           eib,
		EventInstanceLiveDriverBase: eildb,
		EventInstanceLiveCarBase:    eilcb,
	})
}
