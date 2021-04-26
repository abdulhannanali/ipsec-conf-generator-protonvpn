package protonvpn

type OptionalIntField struct {
	Value  int
	Exists bool
}

func CreateOptionalIntField(val int) OptionalIntField {
	return OptionalIntField{val, true}
}

type FilterOptions struct {
	Tier     OptionalIntField
	Status   OptionalIntField
	Features OptionalIntField

	ExcludedCountries []string
	EntryCountry      string
	ExitCountry       string
}

/**
 * Filters the servers based on various metrics
 */
func FilterServers(logicalServers LogicalServers, filterOptions *FilterOptions) LogicalServers {
	var filteredLogicalServers LogicalServers = make([]LogicalServer, 0)

	for _, logicalServer := range logicalServers {
		if filter(&logicalServer, filterOptions) {
			filteredLogicalServers = append(filteredLogicalServers, logicalServer)
		}
	}

	return filteredLogicalServers
}

/**
 * Should filter the logical server or not
 */
func filter(logicalServer *LogicalServer, filterOptions *FilterOptions) bool {
	if filterOptions.EntryCountry != "" && (filterOptions.EntryCountry != logicalServer.EntryCountry) {
		return false
	}

	if filterOptions.ExitCountry != "" && (filterOptions.ExitCountry != logicalServer.ExitCountry) {
		return false
	}

	if filterOptions.ExcludedCountries != nil {
		isExcludedEntryCountry := isCountryExcluded(filterOptions.ExcludedCountries, logicalServer.EntryCountry)
		isExcludedExitCountry := isCountryExcluded(filterOptions.ExcludedCountries, logicalServer.ExitCountry)

		if isExcludedEntryCountry || isExcludedExitCountry {
			return false
		}
	}

	if filterOptions.Status.Exists && filterOptions.Status.Value != logicalServer.Status {
		return false
	}

	if filterOptions.Tier.Exists && filterOptions.Tier.Value != logicalServer.Tier {
		return false
	}

	if filterOptions.Features.Exists && Feature(filterOptions.Features.Value) != logicalServer.Features {
		return false
	}

	return true
}

func isCountryExcluded(excludedCountries []string, country string) bool {
	for _, excludedCountry := range excludedCountries {
		if excludedCountry == country {
			return true
		}
	}

	return false
}
