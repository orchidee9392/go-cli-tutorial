package cmd

import "github.com/spf13/cobra"

//ルートコマンド
var RootCmd = &cobra.Command{
	Use:   "textkit",
	Short: "Minimal hello CLI",
}

//サブコマンドの登録
func init() {
	RootCmd.AddCommand(HelloCmd)
}
