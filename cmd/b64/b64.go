package b64

import "github.com/spf13/cobra"

//b64コマンド
var B64Cmd = &cobra.Command{
	Use: "b64 [encode/decode] [args]",
	Short: "Base64の変換ツール",
}

//サブコマンドの登録
func init(){
	B64Cmd.PersistentFlags().StringP("in", "i", "", "入力ファイルパス")
	B64Cmd.PersistentFlags().StringP("out", "o", "", "出力ファイルパス")

	B64Cmd.AddCommand(EncodeCmd)
	B64Cmd.AddCommand(DecodeCmd)
}