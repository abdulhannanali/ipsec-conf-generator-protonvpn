package main

import (
	"os"

	"github.com/abdulhannanali/protonvpn-ikev2-generator/cli"
)

func main() {
	err := cli.InitCmd(os.Args)

	if err != nil {
		panic(err)
	}
}
