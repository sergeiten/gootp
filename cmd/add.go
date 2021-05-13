package cmd

import (
	"fmt"
	"os"

	db "github.com/sergeiten/gootp/database"
	"github.com/spf13/cobra"
)

const (
	flagName      = "name"
	flagNameShort = "n"
	flagKey       = "key"
	flagKeyShort  = "k"
)

var name string
var key string

var addCmd = &cobra.Command{
	Use:   "add [NAME] [KEY]",
	Short: "Adds new OTP key",
	Run: func(cmd *cobra.Command, args []string) {
		err := db.Insert(name, key)
		if err != nil {
			fmt.Printf("failed to insert %s name for %s key: %v\n", name, key, err)
			os.Exit(1)
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		_ = db.Close()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&name, flagName, flagNameShort, "", "New OTP record name")
	addCmd.Flags().StringVarP(&key, flagKey, flagKeyShort, "", "New OTP record key")
	_ = addCmd.MarkFlagRequired(flagName)
	_ = addCmd.MarkFlagRequired(flagKey)
}
