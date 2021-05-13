package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	db "github.com/sergeiten/gootp/database"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display list all keys name",
	Run: func(cmd *cobra.Command, args []string) {
		keys := db.All()

		for _, k := range keys {
			fmt.Printf("%s\n", k)
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		_ = db.Close()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
