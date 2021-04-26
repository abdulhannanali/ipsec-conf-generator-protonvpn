package protonvpn

type ServerInfo struct {
	EntryIP           string
	ExitIP            string
	Domain            string
	ID                string
	Generation        int
	Status            int
	ServiceDownReason string
	Label             string
}
