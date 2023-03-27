package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// このファイルではサブコマンドを作成している
var (
	Version   string
	ReVersion string
)

func init() {
	RootCmd.AddCommand(newVersionCmd())
}

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print Version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("hoge version: %s, reversion: %s\n", Version, ReVersion)
		},
	}

	return cmd
}
