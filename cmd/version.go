package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print GoOTP version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gootp version %s\n", "0.0.1")
	},
}
