package cli

import (
	"fmt"
	"os"
	"path/filepath"
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
	if len(pathValue) > 0 {
		dir := filepath.Dir(pathValue)
		dirStat, err := os.Stat(dir)

		if err != nil {
			return err
		}

		if dirStat != nil && !dirStat.IsDir() {
			return fmt.Errorf("\"%s\" is a file", pathValue)
		}

		fileStat, err := os.Stat(pathValue)

		if err != nil && strings.Contains(err.Error(), "no such file") {
			return nil
		}

		if fileStat.IsDir() {
			return fmt.Errorf("\"%s\" is a directory", pathValue)
		}

		if fileStat != nil {
			return fmt.Errorf("\"%s\" already exists", pathValue)
		}
	}

	return nil
}
