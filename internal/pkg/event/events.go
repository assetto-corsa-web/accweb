package event

import (
	"time"
)

type Eventer interface {
	GetInfo() EventBase
}

type EventBase struct {
	Name      string    `json:"eventName"`
	Timestamp time.Time `json:"timestamp"`
}

func (e EventBase) GetInfo() EventBase {
	return e
}

func eventBase(name string) EventBase {
	return EventBase{
		Name:      name,
		Timestamp: time.Now().UTC(),
	}
}

type EventInstanceBase struct {
	InstanceId string `json:"instanceId"`
	Name       string `json:"instanceName"`
	Track      string `json:"track"`
	UdpPort    int    `json:"udpPort"`
	TcpPort    int    `json:"tcpPort"`
	NrDrivers  int    `json:"nrDrivers"`
}

func NewEventInstanceBase(id, name, track string, tcp, udp, nrDrivers int) EventInstanceBase {
	return EventInstanceBase{
		InstanceId: id,
		Name:       name,
		Track:      track,
		UdpPort:    udp,
		TcpPort:    tcp,
		NrDrivers:  nrDrivers,
	}
}

type EventInstanceBeforeStart struct {
	EventBase
	EventInstanceBase
}

func EmmitEventInstanceBeforeStart(eib EventInstanceBase) {
	Emmit(EventInstanceBeforeStart{
		EventBase:         eventBase("instance_before_start"),
		EventInstanceBase: eib,
	})
}

type EventInstanceStarted struct {
	EventBase
	EventInstanceBase
}

func EmmitEventInstanceStarted(eib EventInstanceBase) {
	Emmit(EventInstanceStarted{
		EventBase:         eventBase("instance_started"),
		EventInstanceBase: eib,
	})
}

type EventInstanceBeforeStop struct {
	EventBase
	EventInstanceBase
}

func EmmitEventInstanceBeforeStop(eib EventInstanceBase) {
	Emmit(EventInstanceBeforeStop{
		EventBase:         eventBase("instance_before_stop"),
		EventInstanceBase: eib,
	})
}

type EventInstanceStopped struct {
	EventBase
	EventInstanceBase
}

func EmmitEventInstanceStopped(eib EventInstanceBase) {
	Emmit(EventInstanceStopped{
		EventBase:         eventBase("instance_stopped"),
		EventInstanceBase: eib,
	})
}

type EventInstanceOutput struct {
	EventBase
	EventInstanceBase

	Output []byte `json:"output"`
}

func EmmitEventInstanceOutput(eib EventInstanceBase, o []byte) {
	Emmit(EventInstanceOutput{
		EventBase:         eventBase("instance_output"),
		EventInstanceBase: eib,
		Output:            o,
	})
}
