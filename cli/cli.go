package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/abdulhannanali/ipsec-conf-generator-protonvpn/generator"
	"github.com/abdulhannanali/ipsec-conf-generator-protonvpn/protonvpn"
)

const defaultIntValue int = -1
const defaultStringValue string = ""

const githubRepoLink string = "https://github.com/abdulhannanali/ikev2-protonvpn-generator"
const flagSetName string = "protonvpn-profiles-generator"

type FlagValues struct {
	certPath                  *string
	eapIdentity               *string
	fetchOffline              *bool
	filterByStatus            *int
	filterByTier              *int
	filterByExcludedCountries *string
	filterByEntryCountry      *string
	filterByExitCountry       *string
	writeToFile               *string
}

var flagValues FlagValues

func InitCmd(args []string) error {
	cmd, flagValues := createCmdFlagSet()

	if len(os.Args) == 1 {
		cmd.Usage()
		return nil
	}

	err := cmd.Parse(args[1:])

	if err != nil {
		return err
	}

	if err := validateFlagValues(flagValues); err != nil {
		return err
	}

	err = execute(flagValues)

	if err != nil {
		return err
	}

	return nil
}

func execute(flagValues *FlagValues) error {
	filePath := *flagValues.writeToFile
	configurationFile, err := GenerateteConfigurationFile(flagValues)

	if err != nil {
		return err
	}

	if filePath != "" {
		err = os.WriteFile(filePath, []byte(configurationFile), 0666)

		if err != nil {
			return err
		}
	} else {
		os.Stdout.WriteString(configurationFile)
	}

	return nil
}

func createCmdFlagSet() (*flag.FlagSet, *FlagValues) {
	var Cmd *flag.FlagSet = flag.NewFlagSet(flagSetName, flag.ExitOnError)
	flagValues = FlagValues{}

	Cmd.Usage = func() {
		err := helpTemplate.Execute(Cmd.Output(), helpTemplateParams{githubRepoLink})

		if err != nil {
			panic(err)
		}

		fmt.Fprintf(Cmd.Output(), "\n\nUsage of %s:\n\n", flagSetName)
		Cmd.PrintDefaults()
		fmt.Println("")
	}

	flagValues.certPath = Cmd.String("certPath", defaultStringValue, "path of the protonvpn's certificate")
	flagValues.eapIdentity = Cmd.String("eapIdentity", defaultStringValue, "eap_identity going to be used")
	flagValues.fetchOffline = Cmd.Bool("fetchOffline", false, "uses the cached the protonvpn server list, it may be stale or compromised, do it at your own risk")
	flagValues.filterByStatus = Cmd.Int("filterByStatus", defaultIntValue, "Filters the list of protonvpn servers by status")
	flagValues.filterByTier = Cmd.Int("filterByTier", defaultIntValue, "Default Tier to be used, accepted values are ")
	flagValues.filterByExcludedCountries = Cmd.String("filterByExcludedCountries", defaultStringValue, "Filters by a comma separated list of excluded countries")
	flagValues.filterByEntryCountry = Cmd.String("filterByEntryCountry", defaultStringValue, "Filters by Entry Country")
	flagValues.filterByExitCountry = Cmd.String("filterByExitCountry", defaultStringValue, "Filters by Exit Country")
	flagValues.writeToFile = Cmd.String("writeToFile", "", "Writes to a file at given location if specified, otherwise prints to stdout")

	return Cmd, &flagValues
}

func GenerateteConfigurationFile(flagValues *FlagValues) (string, error) {
	filterOptions := constructFilterOptions(flagValues)

	inputParameters := generator.GenerationInputParameters{
		CertPath:      *flagValues.certPath,
		EapIdentity:   *flagValues.eapIdentity,
		FetchOffline:  *flagValues.fetchOffline,
		FilterOptions: &filterOptions,
	}

	str, err := generator.CreateFile(inputParameters)

	if err != nil {
		return "", err
	}

	return str, nil
}

func constructFilterOptions(flagValues *FlagValues) generator.FilterOptions {
	fo := generator.FilterOptions{}

	fo.ExitCountry = *flagValues.filterByExitCountry
	fo.EntryCountry = *flagValues.filterByEntryCountry
	fo.ExcludedCountries = strings.Split(*flagValues.filterByExcludedCountries, ",")

	for i, country := range fo.ExcludedCountries {
		fo.ExcludedCountries[i] = strings.Trim(country, " ")
	}

	isDefaultInt := func(i int) bool {
		return i == defaultIntValue
	}

	if !isDefaultInt(*flagValues.filterByStatus) {
		fo.Status = protonvpn.OptionalIntField{Exists: true, Value: *flagValues.filterByStatus}
	}

	if !isDefaultInt(*flagValues.filterByTier) {
		fo.Tier = protonvpn.CreateOptionalIntField(*flagValues.filterByTier)
	}

	return fo
}
