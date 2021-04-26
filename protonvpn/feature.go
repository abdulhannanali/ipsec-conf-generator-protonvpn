package protonvpn

import (
	"fmt"
)

type Feature int

const (
	FeatureP2P     Feature = 1
	FeatureS2C     Feature = 2
	FeatureTOR     Feature = 4
	FeatureUNKNOWN Feature = 0
)

func ParseFeatureFromString(str string) (Feature, error) {
	switch str {
	default:
		return FeatureUNKNOWN, fmt.Errorf("unrecognized feature '%s'\n, please just use the number for this feature/file an issue in this package's repo", str)
	case "P2P":
		return FeatureP2P, nil
	case "S2C":
		return FeatureS2C, nil
	case "TOR":
		return FeatureTOR, nil
	}
}
