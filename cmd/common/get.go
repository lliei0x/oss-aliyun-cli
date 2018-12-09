package common

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdGet = &cobra.Command{
	Use:   "get",
	Short: "获取存储空间(bucket) 或者 文件对象(object) 列表",
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if bucket == "" && object == "" {
			list, err := bClient.ListBucket()
			if err != nil {
				fmt.Println("获取(bucket)列表出错，具体错误如下：")
				fmt.Println(err)
			} else {
				fmt.Printf("获取存储空间(bucket)列表：%v", list)
			}
		} else if bucket != "" {
			list, err := oClient.ListObject(bucket)
			if err != nil {
				fmt.Println("获取文件对象(object)列表出错，具体错误如下：")
				fmt.Println(err)
			} else {
				fmt.Printf("获取文件对象(object)列表：%v", list)
			}
		} else {
			fmt.Println("指定命令参数不完整，需指定存储空间，文件对象，以及本地路径,具体可使用帮助命令 -h or --help")
		}
	},
}
