package main

import (
	"os"

	"github.com/abdulhannanali/ipsec-conf-generator-protonvpn/cli"
)

func main() {
	err := cli.InitCmd(os.Args)

	if err != nil {
		panic(err)
	}
}
