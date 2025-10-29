package b64

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

// 文字列をbase64にエンコードするメソッド
func runEncode(cmd *cobra.Command, args []string) error {
	encoded := base64.StdEncoding.EncodeToString([]byte(args[0]))
	_, err := fmt.Fprintln(cmd.OutOrStdout(), encoded)
	return err
}

// encode サブコマンド
var EncodeCmd = &cobra.Command{
	Use:   "encode <text>",
	Short: "文字列を Base64 にエンコード",
	Args:  cobra.ExactArgs(1),
	RunE:  runEncode,
}
