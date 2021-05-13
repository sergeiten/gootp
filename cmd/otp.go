package cmd

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"time"

	db "github.com/sergeiten/gootp/database"
	"github.com/spf13/cobra"
)

var OTPCommand = &cobra.Command{
	Use: "otp [NAME]",
	Run: func(cmd *cobra.Command, args []string) {
		val, err := db.Get(name)
		if err != nil {
			fmt.Printf("failed to get value for %s: %v\n", key, err)
			os.Exit(1)
		}

		nowInSeconds := time.Now().Unix()
		stepInSeconds := int64(30)

		counter := uint64(nowInSeconds / stepInSeconds)

		val = bytes.ToUpper(val)
		secretBytes := make([]byte, len(val))
		_, err = base32.StdEncoding.Decode(secretBytes, val)
		if err != nil {
			fmt.Printf("failed to decode secret string: %v", err)
			os.Exit(1)
		}

		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, counter)

		h := hmac.New(sha1.New, secretBytes)
		h.Write(buf)

		sum := h.Sum(nil)
		offset := sum[len(sum)-1] & 0xf

		value := int64(((int(sum[offset]) & 0x7f) << 24) |
			((int(sum[offset+1] & 0xff)) << 16) |
			((int(sum[offset+2] & 0xff)) << 8) |
			(int(sum[offset+3]) & 0xff))

		length := 6
		token := int32(value % int64(math.Pow10(length)))

		fmt.Printf("%s token is: %d", name, token)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		_ = db.Close()
	},
}

func init() {
	rootCmd.AddCommand(OTPCommand)
	OTPCommand.Flags().StringVarP(&name, flagName, flagNameShort, "", "Name of item")
	OTPCommand.MarkFlagRequired(flagName)
}
