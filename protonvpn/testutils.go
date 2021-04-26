package protonvpn

func getLogicalServer() LogicalServer {
	return LogicalServer{
		Name:         "CA#1",
		EntryCountry: "CA",
		ExitCountry:  "CA",
		Domain:       "ca-01.protonvpn.com",
		Tier:         1,
		Features:     0,
		Region:       "",
		City:         "Toronto",
		Score:        2.6194926299999999,
		HostCountry:  "",
		ID:           "BzHqSTaqcpjIY9SncE5s7FpjBrPjiGOucCyJmwA6x4nTNqlElfKvCQFr9xUa2KgQxAiHv4oQQmAkcA56s3ZiGQ==",
		Location: LogicalServerLocation{
			Lat:  43.632899999999999,
			Long: -79.361099999999993,
		},
		Status: 1,
		Servers: []ServerInfo{
			{
				EntryIP:           "104.254.92.59",
				ExitIP:            "104.254.92.59",
				Domain:            "ca-01.protonvpn.com",
				ID:                "ceemvAaEwUf5EDfEYoBV8S4h29Cn5vxhZ5yDLR0h_ErRdnyFdt0lhKhpYbtkNNVHWKu9jbDTTf4sKP-sH_NDzA==",
				Generation:        0,
				Status:            1,
				ServiceDownReason: "",
				Label:             "",
			},
		},
		Load: 58,
	}
}
