package cli

import "testing"

func TestResolveCountry(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"US", "US"},
		{"Bangladesh", "BD"},
		{"Pakistan", "PK"},
		{"India", "IN"},
		{"China", "CN"},
		{"CN", "CN"},
		{"AU", "AU"},
		{"iNdoNesia", "ID"},
		{"bRAziL", "BR"},
		{"fR", "FR"},
	}

	for _, test := range tests {
		resolvedCountryCode, err := ResolveCountry(test.input)

		if err != nil {
			t.Errorf("Didn't expect error for input:%s\n%v", test.input, err)
		}

		if resolvedCountryCode != test.expected {
			t.Errorf("Resolved Country Code (%s) != Expected (%s)", resolvedCountryCode, test.expected)
		}
	}
}
