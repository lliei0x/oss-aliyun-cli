package common

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdDownload = &cobra.Command{
	Use:   "download",
	Short: "下载指定的文件对象(object)",
	// Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		if bucket != "" && object != "" && filepath != "" {
			err := oClient.GetObject(bucket, object, filepath)
			if err != nil {
				fmt.Printf("下载指定的文件对象(object) %v"+" 出错，具体错误如下：", object)
				fmt.Println("")
				fmt.Println(err)
			} else {
				fmt.Printf("已下载指定的文件对象(object)：%v 至 本地路径：%v", object, filepath)
			}
		} else {
			fmt.Println("指定命令参数不完整，需指定存储空间，文件对象，以及本地路径,具体可使用帮助命令 -h or --help")
		}
	},
}
