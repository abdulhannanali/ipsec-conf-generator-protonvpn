package protonvpn

import (
	"embed"
	_ "embed"
)

//go:embed data/logical.json
var embedFS embed.FS
