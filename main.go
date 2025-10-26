package main

import (
	"fmt"
	"os"

	"github.com/orchidee9392/go-cli-tutorial/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
