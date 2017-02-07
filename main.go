package main

import (
	"os"

	"github.com/diyan/assimilator/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
