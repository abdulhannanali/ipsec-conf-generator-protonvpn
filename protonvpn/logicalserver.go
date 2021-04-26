package protonvpn

type LogicalServer struct {
	Name         string
	EntryCountry string
	ExitCountry  string
	Domain       string
	Tier         int
	Features     Feature
	Region       string
	City         string
	Score        float64
	HostCountry  string
	ID           string
	Location     LogicalServerLocation
	Status       int
	Servers      []ServerInfo
	Load         int
}

type LogicalServerLocation struct {
	Lat  float64
	Long float64
}

func (logicalServer *LogicalServer) isSecureCore() bool {
	return logicalServer.EntryCountry == logicalServer.ExitCountry
}
