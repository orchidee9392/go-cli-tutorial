package b64

import (
	"encoding/base64"
	"strings"

	"github.com/orchidee9392/go-cli-tutorial/internal/iox"
	"github.com/spf13/cobra"
)

// base64 形式のテキストをデコードするメソッド
func runDecode(cmd *cobra.Command, args []string) error {
	inPath, _ := cmd.Flags().GetString("in")
	outPath, _ := cmd.Flags().GetString("out")

	arg := ""
	if len(args) == 1 {
		arg = args[0]
	}

	in, err := iox.ReadInput(arg, inPath, cmd.InOrStdin())
	if err != nil {
		return err
	}

	dec, err := base64.StdEncoding.DecodeString(strings.TrimSpace(string(in)))
	if err != nil {
		return err
	}

	return iox.WriteOutput(outPath, cmd.OutOrStdout(), dec)
}

// decode サブコマンド
var DecodeCmd = &cobra.Command{
	Use:   "decode <b64text>",
	Short: "Base64 テキストをデコード",
	Args:  cobra.MaximumNArgs(1),
	RunE:  runDecode,
}
