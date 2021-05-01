package generator

import (
	"testing"
)

func Test_convertDomainToName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"nl-84.protonvpn.com", "nl_84_protonvpn_com"},
		{"nl-83.protonvpn.com", "nl_83_protonvpn_com"},
		{"a.example.com", "a_example_com"},
		{"b.example.com", "b_example_com"},
		{"c.adsa.asdad.asda.example.com", "c_adsa_asdad_asda_example_com"},
	}

	for _, tt := range tests {
		got := convertProtonVpnDomainToName(tt.input)

		if got != tt.expected {
			t.Errorf("Expected: %s, got: %s", tt.expected, got)
		}
	}
}

func Test_labelConnection(t *testing.T) {
	type labelConnectionInput struct {
		tier      int
		domain    string
		serverIdx int
	}

	tests := []struct {
		input    labelConnectionInput
		expected string
	}{
		{
			labelConnectionInput{1, "ca.01.protonvpn.com", 0},
			"ca_01_protonvpn_com",
		},
	}

	for _, tt := range tests {
		input := tt.input
		result := generateConnectionLabel(input.domain, input.tier, input.serverIdx)

		if result != tt.expected {
			t.Errorf("result: %s, expected: %s", result, tt.expected)
		}
	}
}
