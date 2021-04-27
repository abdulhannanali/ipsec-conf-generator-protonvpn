package protonvpn

type LogicalServers []LogicalServer

func (logicals LogicalServers) GetLogicalServerByName(name string) *LogicalServer {
	for _, logicalServer := range logicals {
		if logicalServer.Name == name {
			return &logicalServer
		}
	}

	return nil
}

func (logicals LogicalServers) GetLogicalServerByDomain(domain string) *LogicalServer {
	for _, logicalServer := range logicals {
		if logicalServer.Domain == domain {
			return &logicalServer
		}
	}

	return nil
}

func (logicals LogicalServers) FilterServers(filterOptions FilterOptions) LogicalServers {
	return FilterServers(logicals, &filterOptions)
}
