package protonvpn

import (
	"testing"
)

func TestFilterServers(t *testing.T) {
	logicalServers := []LogicalServer{
		{EntryCountry: "PK"},
		{EntryCountry: "US"},
		{EntryCountry: "US"},
		{EntryCountry: "SG"},
		{EntryCountry: "CA"},
		{EntryCountry: "UK"},
		{Status: 1},
		{Status: 2},
		{Status: 3},
		{Status: 1},
		{ExitCountry: "IN"},
		{ExitCountry: "AU"},
		{ExitCountry: "UK"},
		{ExitCountry: "UK"},
		{Tier: 1},
		{Tier: 2},
		{Tier: 3},
		{Tier: 1},
	}

	if len(logicalServers) != 18 {
		t.Errorf(
			"Seems like `logicalServers` length has been updated from %d, please update the filter_test file accordingly",
			18,
		)

	}

	tests := []struct {
		input    FilterOptions
		expected int
	}{
		{
			FilterOptions{EntryCountry: "UK"},
			1,
		},

		{
			FilterOptions{ExcludedCountries: []string{"AU", "UK"}},
			14,
		},

		{
			FilterOptions{Tier: CreateOptionalIntField(3), Status: CreateOptionalIntField(3)},
			0,
		},

		{
			FilterOptions{Tier: CreateOptionalIntField(1)},
			2,
		},

		{
			FilterOptions{Status: CreateOptionalIntField(1)},
			2,
		},
	}

	for i, tt := range tests {
		got := len(FilterServers(logicalServers, &tt.input))

		if got != tt.expected {
			t.Errorf(`
				Test Case idx: %d
				Expected: %d
				Got: %d
			`, i, tt.expected, got)
		}
	}
}
