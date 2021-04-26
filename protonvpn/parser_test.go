package protonvpn

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseLogical(t *testing.T) {
	localLogicalFile, err := embedFS.ReadFile("data/logical.json")

	if err != nil {
		t.Errorf("%v", err)
	}

	logicalsResponse, err := parseLogicalResponse(localLogicalFile)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if logicalsResponse.Code != 1000 {
		t.Errorf("Code is not '1000'")
	}

	logicals := logicalsResponse.LogicalServers
	logical := *logicals.GetLogicalServerByName("CA#1")

	diff := cmp.Diff(logical, getLogicalServer())

	fmt.Println("diff", diff)

	if diff != "" {
		t.Errorf("parseLogical failed, mismatch (-want, +got):\n%s", diff)
	}
}
