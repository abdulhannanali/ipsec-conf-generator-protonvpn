package generator

import (
	"strings"

	"github.com/abdulhannanali/protonvpn-ikev2-generator/protonvpn"
)

type FilterOptions protonvpn.FilterOptions

type GenerationInputParameters struct {
	CertPath      string
	EapIdentity   string
	FetchOffline  bool
	FilterOptions *FilterOptions
}

type OutputFile struct {
	FileName   string
	FileOutput string
}

func CreateFile(inputParameters GenerationInputParameters) (string, error) {
	response, err := protonvpn.GetLogicals(inputParameters.FetchOffline)

	if err != nil {
		return "", err
	}

	logicalServers := response.LogicalServers.FilterServers(protonvpn.FilterOptions(*inputParameters.FilterOptions))
	connections, err := generateConnections(inputParameters, logicalServers)

	if err != nil {
		return "", err
	}

	var connectionFile strings.Builder
	err = connectionFileTemplate.Execute(&connectionFile, strings.Join(connections, "\n\n"))

	if err != nil {
		return "", err
	}

	return connectionFile.String(), nil
}

func generateConnections(inputParameters GenerationInputParameters, logicalServers protonvpn.LogicalServers) ([]string, error) {
	connections := make([]string, len(logicalServers))

	for _, logicalServer := range logicalServers {
		for serverIdx, serverInfo := range logicalServer.Servers {
			label := generateConnectionLabel(serverInfo.Domain, logicalServer.Tier, serverIdx)

			connectionString, err := getConnectionString(label, logicalServer.Domain, inputParameters.EapIdentity, inputParameters.CertPath)

			if err != nil {
				return nil, err
			}

			connections = append(connections, connectionString)
		}
	}

	return connections, nil
}

func convertToProtonVpnFilterOptions(opts FilterOptions) *protonvpn.FilterOptions {
	return &protonvpn.FilterOptions{
		EntryCountry:      opts.EntryCountry,
		ExitCountry:       opts.ExitCountry,
		ExcludedCountries: opts.ExcludedCountries,
		Status:            opts.Status,
		Tier:              opts.Tier,
		Features:          opts.Features,
	}
}

func getConnectionString(label string, domain string, eapIdentity string, certPath string) (string, error) {
	type connectionTemplateData struct {
		ConnectionLabel string
		EapIdentity     string
		Domain          string
		CertPath        string
	}

	var b strings.Builder
	err := connectionTemplate.Execute(&b, connectionTemplateData{label, eapIdentity, domain, certPath})

	if err != nil {
		return "", err
	}

	return b.String(), nil
}
