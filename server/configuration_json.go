package server

// ConfigurationJSON represents the configuration.json.
type ConfigurationJSON struct {
	UDPPort         int  `json:"udpPort" label:"UDP Port"`
	TCPPort         int  `json:"tcpPort" label:"TCP Port"`
	MaxConnections  int  `json:"maxConnections" label:"Max Connections"`
	LANDiscovery    bool `json:"lanDiscovery" label:"LAN Discovery"`
	RegisterToLobby bool `json:"registerToLobby" label:"Register to Lobby"`
	ConfigVersion   int  `json:"configVersion"`

	// TEST
	SubTest    SubTest `json:"sub_test" label:"Struct Test"`
	DoesntWork []int   `json:"doesnt_work" label:"This shouldn't be rendered'"`
}

// TEST
type SubTest struct {
	Str          string    `json:"str" label:"String Test"`
	Select       string    `json:"select" label:"Select Test" options:"opt1:Option 1,opt2:Option 2,opt3:Option 3"`
	Select2      int       `json:"select2" label:"Select Test 2" options:"42:Option 1,43:Option 2,44:Option 3"`
	SubTestSlice []SubTest `json:"sub_test_slice" label:"Struct Test Array"`
}
