package common

import (
	"fmt"
	"oss-aliyun-cli/infra/upload"

	"github.com/spf13/cobra"
)

var CmdUp = &cobra.Command{
	Use:   "up",
	Short: "上传指定文件 或者 并发上传指定目录以及子目录下符合过滤条件的文件",
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if bucket != "" && filepath != "" {
			err := upload.Upload(filepath, suffix, bucket, object, oClient)
			if err != nil {
				fmt.Println("上传指定文件或目录出错，具体错误如下：")
				fmt.Println(err)
			} else {
				fmt.Printf("上传指定文件或目录成功")
			}
		} else {
			fmt.Println("指定命令参数不完整，需指定存储空间，文件对象，以及本地路径,具体可使用帮助命令 -h or --help")
		}
	},
}
