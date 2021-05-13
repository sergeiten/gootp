package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gootp [COMMAND]",
	Short: "GoOTP - One Time Password manager",
	Long:  "GoOTP is dead simple One Time Password manager that lets you save your passwords locally",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
