package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dga",
	Short: "Docker Github Action for simplifying repetitive actions",
	Long: `Sometimes you'll want to put 4 tags on a container... and push 
them all to a registry... after logging in. This tool will aid 
in making the number of steps to do that as small as possible.`,
	Version: "v0.1",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
