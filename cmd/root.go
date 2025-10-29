package cmd

import (
	"github.com/orchidee9392/go-cli-tutorial/cmd/b64"
	"github.com/spf13/cobra"
)

// ルートコマンド
var RootCmd = &cobra.Command{
	Use:   "textkit",
	Short: "Minimal hello CLI",
}

// サブコマンドの登録
func init() {
	RootCmd.AddCommand(HelloCmd)
	RootCmd.AddCommand(b64.B64Cmd)
}
