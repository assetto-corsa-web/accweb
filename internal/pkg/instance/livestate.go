package instance

type ServerState string

const (
	ServerStateOffline       ServerState = "offline"
	ServerStateStarting      ServerState = "starting"
	ServerStateNotRegistered ServerState = "not_registered"
	ServerStateOnline        ServerState = "online"
)

type LiveState struct {
	ServerState ServerState `json:"server_state"`
	NrClients   int         `json:"nr_clients"`
	Track       string      `json:"track"`
}

func newLiveState() *LiveState {
	return &LiveState{
		ServerState: ServerStateOffline,
	}
}

func (l *LiveState) setServerState(s ServerState) {
	l.ServerState = s
}

func (l *LiveState) setNrClients(nr int) {
	l.NrClients = nr
}

func (l *LiveState) setTrack(track string) {
	l.Track = track
}
