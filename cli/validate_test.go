package cli

import (
	"strings"
	"testing"
)

func Test_validateStringFlag(t *testing.T) {
	tests := []struct {
		input         []string
		errorExpected bool
	}{
		{
			[]string{"one", ""},
			true,
		},
		{
			[]string{"two", "expected"},
			false,
		},
		{
			[]string{"three", "exprected_another"},
			false,
		},
	}

	for i, tt := range tests {
		errorExpected := tt.errorExpected
		err := validateStringFlag(tt.input[0], tt.input[1])

		if errorExpected && err == nil {
			t.Errorf("Error was expected but no error for test case %d", i)
		} else if !errorExpected && err != nil {
			t.Errorf("Error not expected but there's error for test case %d\n%v", i, err)
		}
	}
}

func Test_validatePathFlag(t *testing.T) {
	tests := []struct {
		input   string
		errPart string
	}{
		{
			"./fixtures/validatePathFlag/one",
			"already exists",
		},

		{
			"./fixtures/validatePathFlag/three",
			"",
		},

		{
			"./fixtures/validatePathFlag/sampleDirectory",
			"is a directory",
		},

		{
			"./fixtures/validatePathFlag/doesNotExist",
			"",
		},
	}

	for i, tt := range tests {
		errPart := tt.errPart
		err := validatePathFlag("", tt.input)

		if err == nil && errPart != "" {
			t.Errorf("error expected but no error returned for test case: %d", i)
		}

		if err != nil && errPart != "" {
			if !(strings.Contains(err.Error(), errPart)) {
				t.Errorf("error does not contain \"%s\"\nfullError: %s", errPart, err.Error())
			}
		}
	}
}
