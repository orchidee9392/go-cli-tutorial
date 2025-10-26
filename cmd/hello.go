package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// 挨拶メソッド
func runHello(cmd *cobra.Command, args []string) error {
	name := "textkit"
	if len(args) == 1 {
		name = args[0]
	}
	fmt.Fprintf(cmd.OutOrStdout(), "Hello, %s!\n", name)
	return nil
}

// 挨拶サブコマンド
var HelloCmd = &cobra.Command{
	Use:   "hello [name]",
	Short: "挨拶します",
	Args:  cobra.MaximumNArgs(1),
	RunE:  runHello,
}
