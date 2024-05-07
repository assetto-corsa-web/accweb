package instance

import "time"

const (
	accwebConfigJsonName  = "accwebConfig.json"
	configurationJsonName = "configuration.json"
	settingsJsonName      = "settings.json"
	eventJsonName         = "event.json"
	eventRulesJsonName    = "eventRules.json"
	entrylistJsonName     = "entrylist.json"
	bopJsonName           = "bop.json"
	assistRulesJsonName   = "assistRules.json"
	configVersion         = 1

	WinCpuPriorityRealtime    = 256
	WinCpuPriorityHigh        = 128
	WinCpuPriorityAboveNormal = 32768
	WinCpuPriorityNormal      = 32
	WinCpuPriorityBelowNormal = 16384
	WinCpuPriorityLow         = 64
)

var (
	CpuPriorities = map[int]bool{
		WinCpuPriorityRealtime:    true,
		WinCpuPriorityHigh:        true,
		WinCpuPriorityAboveNormal: true,
		WinCpuPriorityNormal:      true,
		WinCpuPriorityBelowNormal: true,
		WinCpuPriorityLow:         true,
	}
)

type AccConfigFiles struct {
	Configuration ConfigurationJson `json:"configuration"`
	Settings      SettingsJson      `json:"settings"`
	Event         EventJson         `json:"event"`
	EventRules    EventRulesJson    `json:"eventRules"`
	Entrylist     EntrylistJson     `json:"entrylist"`
	Bop           BopJson           `json:"bop"`
	AssistRules   AssistRulesJson   `json:"assistRules"`
}

type AccWebConfigJson struct {
	ID        string             `json:"id"`
	Md5Sum    string             `json:"md5Sum"`
	AutoStart bool               `json:"autoStart"` // backward compatibility
	Settings  AccWebSettingsJson `json:"settings"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

type AccWebSettingsJson struct {
	AutoStart       bool                          `json:"autoStart"`
	EnableAdvWinCfg bool                          `json:"enableAdvWindowsCfg"`
	AdvWindowsCfg   *AccWebAdvWindowsSettingsJson `json:"advWindowsCfg"`
}

type AccWebAdvWindowsSettingsJson struct {
	CpuPriority  uint `json:"cpuPriority"`
	CoreAffinity uint `json:"coreAffinity"`
	EnableWinFW  bool `json:"enableWindowsFirewall"`
}

func (a *AccWebConfigJson) SetUpdateAt() {
	a.UpdatedAt = time.Now().UTC()
}

type ConfigurationJson struct {
	ConfigVersion   int    `json:"configVersion"`
	UdpPort         int    `json:"udpPort"`
	TcpPort         int    `json:"tcpPort"`
	MaxConnections  int    `json:"maxConnections"`
	RegisterToLobby int    `json:"registerToLobby"`
	LanDiscovery    int    `json:"lanDiscovery"`
	PublicIP        string `json:"publicIP"`
}

type SettingsJson struct {
	ConfigVersion              int    `json:"configVersion"`
	ServerName                 string `json:"serverName"`
	Password                   string `json:"password"`
	AdminPassword              string `json:"adminPassword"`
	SpectatorPassword          string `json:"spectatorPassword"`
	TrackMedalsRequirement     int    `json:"trackMedalsRequirement"`
	SafetyRatingRequirement    int    `json:"safetyRatingRequirement"`
	RacecraftRatingRequirement int    `json:"racecraftRatingRequirement"`
	IgnorePrematureDisconnects int    `json:"ignorePrematureDisconnects"`
	DumpLeaderboards           int    `json:"dumpLeaderboards"`
	IsRaceLocked               int    `json:"isRaceLocked"`
	RandomizeTrackWhenEmpty    int    `json:"randomizeTrackWhenEmpty"`
	MaxCarSlots                int    `json:"maxCarSlots"`
	CentralEntryListPath       string `json:"centralEntryListPath"`
	ShortFormationLap          int    `json:"shortFormationLap"`
	AllowAutoDQ                int    `json:"allowAutoDQ"`
	DumpEntryList              int    `json:"dumpEntryList"`
	FormationLapType           int    `json:"formationLapType"`
	CarGroup                   string `json:"carGroup"`
}

type EventJson struct {
	ConfigVersion                 int               `json:"configVersion"`
	Track                         string            `json:"track"`
	PreRaceWaitingTimeSeconds     int               `json:"preRaceWaitingTimeSeconds"`
	SessionOverTimeSeconds        int               `json:"sessionOverTimeSeconds"`
	AmbientTemp                   int               `json:"ambientTemp"`
	TrackTemp                     int               `json:"trackTemp,omitempty"`
	CloudLevel                    float64           `json:"cloudLevel"`
	Rain                          float64           `json:"rain"`
	WeatherRandomness             int               `json:"weatherRandomness"`
	Sessions                      []SessionSettings `json:"sessions"`
	MetaData                      string            `json:"metaData"`
	PostQualySeconds              int               `json:"postQualySeconds"`
	PostRaceSeconds               int               `json:"postRaceSeconds"`
	SimracerWeatherConditions     int               `json:"simracerWeatherConditions"`
	IsFixedConditionQualification int               `json:"isFixedConditionQualification"`
}

type EventRulesJson struct {
	ConfigVersion                        int  `json:"configVersion"`
	QualifyStandingType                  int  `json:"qualifyStandingType"`
	PitWindowLengthSec                   int  `json:"pitWindowLengthSec"`
	DriverStintTimeSec                   int  `json:"driverStintTimeSec"`
	MandatoryPitstopCount                int  `json:"mandatoryPitstopCount"`
	MaxTotalDrivingTime                  int  `json:"maxTotalDrivingTime"`
	MaxDriversCount                      int  `json:"maxDriversCount"`
	IsRefuellingAllowedInRace            bool `json:"isRefuellingAllowedInRace"`
	IsRefuellingTimeFixed                bool `json:"isRefuellingTimeFixed"`
	IsMandatoryPitstopRefuellingRequired bool `json:"isMandatoryPitstopRefuellingRequired"`
	IsMandatoryPitstopTyreChangeRequired bool `json:"isMandatoryPitstopTyreChangeRequired"`
	IsMandatoryPitstopSwapDriverRequired bool `json:"isMandatoryPitstopSwapDriverRequired"`
	TyreSetCount                         int  `json:"tyreSetCount"`
}

type SessionSettings struct {
	HourOfDay              int    `json:"hourOfDay"`
	DayOfWeekend           int    `json:"dayOfWeekend"`
	TimeMultiplier         int    `json:"timeMultiplier"`
	SessionType            string `json:"sessionType"`
	SessionDurationMinutes int    `json:"sessionDurationMinutes"`
}

type EntrylistJson struct {
	ConfigVersion  int             `json:"configVersion"`
	Entries        []EntrySettings `json:"entries"`
	ForceEntryList int             `json:"forceEntryList"`
}

type EntrySettings struct {
	Drivers                      []DriverSettings `json:"drivers"`
	RaceNumber                   *int             `json:"raceNumber,omitempty"`
	ForcedCarModel               *int             `json:"forcedCarModel,omitempty"`
	OverrideDriverInfo           int              `json:"overrideDriverInfo"`
	IsServerAdmin                *int             `json:"isServerAdmin,omitempty"`
	CustomCar                    *string          `json:"customCar,omitempty"`
	OverrideCarModelForCustomCar int              `json:"overrideCarModelForCustomCar"`
	BallastKg                    *int             `json:"ballastKg,omitempty"`
	Restrictor                   *int             `json:"restrictor,omitempty"`
	DefaultGridPosition          *int             `json:"defaultGridPosition,omitempty"`
}

type DriverSettings struct {
	FirstName      *string `json:"firstName,omitempty"`
	LastName       *string `json:"lastName,omitempty"`
	ShortName      *string `json:"shortName,omitempty"`
	DriverCategory *int    `json:"driverCategory,omitempty"`
	PlayerID       string  `json:"playerID"`
	Nationality    *int    `json:"nationality,omitempty"`
}

type BopJson struct {
	ConfigVersion int           `json:"configVersion"`
	Entries       []BopSettings `json:"entries"`
}

type BopSettings struct {
	Track      string `json:"track"`
	CarModel   int    `json:"carModel"`
	BallastKg  int    `json:"ballastKg"`
	Restrictor int    `json:"restrictor"`
}

type AssistRulesJson struct {
	ConfigVersion            int `json:"configVersion"`
	StabilityControlLevelMax int `json:"stabilityControlLevelMax"`
	DisableAutosteer         int `json:"disableAutosteer"`
	DisableAutoLights        int `json:"disableAutoLights"`
	DisableAutoWiper         int `json:"disableAutoWiper"`
	DisableAutoEngineStart   int `json:"disableAutoEngineStart"`
	DisableAutoPitLimiter    int `json:"disableAutoPitLimiter"`
	DisableAutoGear          int `json:"disableAutoGear"`
	DisableAutoClutch        int `json:"disableAutoClutch"`
	DisableIdealLine         int `json:"disableIdealLine"`
}
