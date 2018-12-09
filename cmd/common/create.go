package common

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CmdCreate 创建存储空间
var CmdCreate = &cobra.Command{
	Use:   "create",
	Short: "新建存储空间(bucket) 或者 文件对象(object)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if bucket != "" && object == "" {
			err := bClient.CreateBucket(bucket)
			if err != nil {
				fmt.Printf("新建存储空间(bucket) %v 出错，具体错误如下：", bucket)
				fmt.Println("")
				fmt.Println(err)
			} else {
				fmt.Printf("新建存储空间(bucket)：%v ", bucket)
			}
		} else {
			fmt.Println("指定命令参数不完整，需指定存储空间，文件对象，以及本地路径,具体可使用帮助命令 -h or --help")
		}
	},
}
