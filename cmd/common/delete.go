package common

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdDelete = &cobra.Command{
	Use:   "delete",
	Short: "删除指定存储空间(bucket) 或者 文件对象(object)",
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if bucket != "" && object == "" {
			err := bClient.DeleteBucket(bucket)
			if err != nil {
				fmt.Printf("删除指定存储空间(bucket) %v 出错，具体错误如下：", bucket)
				fmt.Println("")
				fmt.Println(err)
			} else {
				fmt.Printf("已删除指定存储空间(bucket)：%v", bucket)
			}
		} else if bucket != "" && object != "" {
			err := oClient.DeleteObject(bucket, object)
			if err != nil {
				fmt.Printf("删除指定文件对象(object) %v 出错，具体错误如下：", object)
				fmt.Println("")
				fmt.Println(err)
			} else {
				fmt.Printf("已删除指定文件对象(object)：%v", object)
			}
		} else {
			fmt.Println("指定命令参数不完整，需指定存储空间，文件对象，以及本地路径,具体可使用帮助命令 -h or --help")
		}
	},
}
