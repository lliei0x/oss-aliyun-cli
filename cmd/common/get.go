package common

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdGet = &cobra.Command{
	Use:   "get",
	Short: "获取存储空间(bucket) 或者 文件对象(object) 列表",
	// Args:  cobra.MinimumNArgs(1),
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 	},
}

var cmdGetBucket = &cobra.Command{
	Use:   "bucket",
	Short: "获取存储空间(bucket)列表",
	Run: func(cmd *cobra.Command, args []string) {
		list, err := bClient.ListBucket()
		if err != nil {
			fmt.Println("获取(bucket)列表出错，具体错误如下：")
			fmt.Println(err)
		} else {
			fmt.Printf("获取存储空间(bucket)列表：%v", list)
		}
	},
}

var cmdGetObject = &cobra.Command{
	Use:   "object",
	Short: "获取存储空间(bucket)列表",
	Run: func(cmd *cobra.Command, args []string) {
		if bucket != "" {
			list, err := oClient.ListObject(bucket)
			if err != nil {
				fmt.Println("获取文件对象(object)列表出错，具体错误如下：")
				fmt.Println(err)
			} else {
				fmt.Printf("获取文件对象(object)列表：%v", list)
			}
		}
	},
}

func init() {
	CmdGet.AddCommand(cmdGetBucket, cmdGetObject)
	cmdGetObject.Flags().StringVarP(&bucket, "bucket", "b", "", "aliyun-oss上的存储空间命名")
}
