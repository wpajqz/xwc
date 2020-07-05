package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wpajqz/xwc/cmd"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "xwc",
		Short: "wrap some command with enviroment",
	}

	rootCmd.DisableFlagParsing = true
	rootCmd.AddCommand(cmd.RunInitCommand())
	rootCmd.AddCommand(cmd.RunCallCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}
