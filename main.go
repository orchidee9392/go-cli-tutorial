package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	//ルートコマンド
	rootCmd := &cobra.Command{
		Use:          "textkit",
		Short:        "Minimal hello CLI",
		SilenceUsage: true,
	}

	//挨拶サブコマンド
	helloCmd := &cobra.Command{
		Use:   "hello [name]",
		Short: "挨拶します",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := "world"
			if len(args) == 1 {
				name = args[0]
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Hello, %s\n", name)
		},
	}

	//サブコマンドの追加
	rootCmd.AddCommand(helloCmd)

	//実行
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
