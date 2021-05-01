package generator

import (
	"strings"
	"testing"
)

func TestConvertToProtonVpnFilterOptions(t *testing.T) {
	gFilterOptions := FilterOptions{
		ExitCountry:  "US",
		EntryCountry: "US",
	}

	pFilterOptions := convertToProtonVpnFilterOptions(gFilterOptions)

	t.Logf("pFilterOptions is %v", pFilterOptions)

	compareStringField := func(field string, gValue string, pValue string) {
		if gValue != pValue {
			t.Errorf(`
				generator.FilterOptions is not being properly converted to
				protonvpn.FilterOptions as %s is not equal for both

				(%v) != (%v)
			`, field, gValue, pValue)
		}
	}

	compareStringField("ExitCountry", gFilterOptions.ExitCountry, pFilterOptions.ExitCountry)
	compareStringField("EntryCountry", gFilterOptions.EntryCountry, pFilterOptions.EntryCountry)
}

func Test_getConnectionString(t *testing.T) {
	tests := []struct {
		input         []string
		expectedParts []string
	}{
		{
			[]string{"Label", "a.example.com", "X", "/home/random"},

			[]string{
				"right=a.example.com",
				"eap_identity=X",
				"rightid=%a.example.com",
				"conn Label",
				"rightca=/home/random",
			},
		},
		{
			[]string{"Label2", "b.example.com", "Y", "/home/random1"},
			[]string{
				"right=b.example.com",
				"rightid=%b.example.com",
				"eap_identity=Y",
				"rightca=/home/random1",
				"conn Label2",
			},
		},
		{
			[]string{"Label3", "c.example.com", "Z", "/home/random2"},
			[]string{
				"rightca=/home/random2",
				"right=c.example.com",
				"eap_identity=Z",
				"rightid=%c.example.com",
				"conn Label3",
			},
		},
	}

	for i, test := range tests {
		input := test.input
		label := input[0]
		domain := input[1]
		eapIdentity := input[2]
		certPath := input[3]

		expectedParts := test.expectedParts
		connectionString, err := getConnectionString(label, domain, eapIdentity, certPath)

		t.Logf("connectionString (%d): %s", i, connectionString)

		if err != nil {
			t.Errorf("%v", err)
		}

		for _, part := range expectedParts {
			if !strings.Contains(connectionString, part) {
				t.Errorf("connectionString of %d does not contain %s", i, part)
			}
		}

	}
}
