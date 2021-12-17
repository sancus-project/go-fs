package main

import (
	"log"

	"github.com/spf13/cobra"
)

const (
	CmdName = "dumb"
)

var rootCmd = &cobra.Command{
	Use: CmdName,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
