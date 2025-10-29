package b64

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// base64 形式のテキストをデコードするメソッド
func runDecode(cmd *cobra.Command, args []string) error {
	dec, err := base64.StdEncoding.DecodeString(strings.TrimSpace(args[0]))
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(cmd.OutOrStdout(), string(dec))
	return err
}

// decode サブコマンド
var DecodeCmd = &cobra.Command{
	Use:   "decode <b64text>",
	Short: "Base64 テキストをデコード",
	Args:  cobra.ExactArgs(1),
	RunE:  runDecode,
}
