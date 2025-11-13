package b64

import (
	"encoding/base64"

	"github.com/orchidee9392/go-cli-tutorial/internal/iox"
	"github.com/spf13/cobra"
)

// 文字列をbase64にエンコードするメソッド
func runEncode(cmd *cobra.Command, args []string) error {
	//フラグ取得
	inPath, _ := cmd.Flags().GetString("in")
	outPath, _ := cmd.Flags().GetString("out")

	//引数取得
	arg := ""
	if len(args) == 1 {
		arg = args[0]
	}

	//入力受け取り
	in, err := iox.ReadInput(arg, inPath, cmd.InOrStdin())
	if err != nil {
		return err
	}

	//エンコード
	encoded := base64.StdEncoding.EncodeToString(in) + "\n"

	//出力
	return iox.WriteOutput(outPath, cmd.OutOrStdout(), []byte(encoded))
}

// encode サブコマンド
var EncodeCmd = &cobra.Command{
	Use:   "encode [text]",
	Short: "文字列を Base64 にエンコード",
	Args:  cobra.MaximumNArgs(1),
	RunE:  runEncode,
}

func init() {
	EncodeCmd.Flags().StringP("in", "i", "", "入力ファイルパス")
	EncodeCmd.Flags().StringP("out", "o", "", "出力ファイルパス")
}
