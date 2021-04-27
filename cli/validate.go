package cli

import (
	"fmt"
	"os"
	"strings"
)

func validateFlagValues(fv *FlagValues) error {
	err := validateStringFlag("certPath", *fv.certPath)
	if err != nil {
		return err
	}

	err = validateStringFlag("eapIdentity", *fv.eapIdentity)
	if err != nil {
		return err
	}

	err = validatePathFlag("writeToFile", *fv.writeToFile)
	if err != nil {
		return err
	}

	return nil
}

func validateStringFlag(fieldName, str string) error {
	if len(str) < 1 {
		return fmt.Errorf("valid %s not provided", fieldName)
	}

	return nil
}

func validatePathFlag(fieldName, pathValue string) error {
	fileInfo, err := os.Stat(pathValue)

	if err != nil {
		if strings.Contains(err.Error(), "no such file") {
			return nil
		} else {
			return err
		}
	}

	if fileInfo != nil {
		return fmt.Errorf("something already exists at %s", pathValue)
	}

	return err
}
