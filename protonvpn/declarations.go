package protonvpn

type LogicalsResponse struct {
	Code           int
	LogicalServers LogicalServers
	FetchedOffline bool `json:"-"`
}
