package cmd

import (
	"fmt"
	"os"

	"oss-aliyun-cli/cmd/common"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "oss [flags]",
	Short: "针对 aliyun-oss 常规操作的命令行工具",

	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

func init() {
	rootCmd.AddCommand(common.CmdUp, common.CmdDownload, common.CmdGet, common.CmdCreate, common.CmdDelete)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
