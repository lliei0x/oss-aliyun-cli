package cmd

import (
	"fmt"
	"os"
	"strings"

	"oss-aliyun-cli/config"
	bu "oss-aliyun-cli/domain/bucket"
	ob "oss-aliyun-cli/domain/object"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/cobra"
)

var (
	bucket   string
	object   string
	filepath string
)

var rootCmd = &cobra.Command{
	Use:   "oss [flags]",
	Short: "针对 aliyun-oss 常规操作的命令行工具",

	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

var cmdUp = &cobra.Command{
	Use:   "up",
	Short: "上传指定文件 或者 并发上传指定目录以及子目录下的文件n",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, " "))
	},
}

var cmdDownload = &cobra.Command{
	Use:   "download",
	Short: "下载指定的文件对象(object)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, " "))
	},
}

var cmdGet = &cobra.Command{
	Use:   "get",
	Short: "获取存储空间(bucket) 或者 文件对象(object) 列表",
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if bucket != "" && object == "" {
			bclient := getBucketClient()
			list, _ := bclient.ListBucket()
			fmt.Println(list)
		} else if bucket == "" && object != "" {
			oclient := getObjectClient()
			list, _ := oclient.ListObject(object)
			fmt.Println(list)
		}
	},
}

var cmdCreat = &cobra.Command{
	Use:   "creat",
	Short: "新建存储空间(bucket) 或者 文件对象(object)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, " "))
	},
}

var cmdDelete = &cobra.Command{
	Use:   "delete",
	Short: "删除指定存储空间(bucket) 或者 文件对象(object)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(cmdUp, cmdDownload, cmdGet, cmdCreat, cmdDelete)

	// Up
	cmdUp.Flags().StringVarP(&bucket, "bucket", "b", "", "aliyun-oss上的存储空间命名")
	cmdUp.Flags().StringVarP(&object, "object", "o", "", "aliyun-oss上的文件对象命名")
	cmdUp.Flags().StringVarP(&filepath, "filepath", "f", "", "指定需要上传文件或者目录的绝对路径")

	// Download
	cmdDownload.Flags().StringVarP(&bucket, "bucket", "b", "", "aliyun-oss上的存储空间命名")
	cmdDownload.Flags().StringVarP(&object, "object", "o", "", "aliyun-oss上的文件对象命名")
	cmdDownload.Flags().StringVarP(&filepath, "filepath", "f", "", "指定需要上传文件或者目录的绝对路径")

	// Get
	cmdGet.Flags().StringVarP(&bucket, "bucket", "b", "", "aliyun-oss上的存储空间命名")
	cmdGet.Flags().StringVarP(&object, "object", "o", "", "aliyun-oss上的文件对象命名")
	cmdGet.Flags().StringVarP(&filepath, "filepath", "f", "", "指定需要上传文件或者目录的绝对路径")

	// Creat
	cmdCreat.Flags().StringVarP(&bucket, "bucket", "b", "", "aliyun-oss上的存储空间命名")
	cmdCreat.Flags().StringVarP(&object, "object", "o", "", "aliyun-oss上的文件对象命名")
	cmdCreat.Flags().StringVarP(&filepath, "filepath", "f", "", "指定需要上传文件或者目录的绝对路径")

	// Delete
	cmdDelete.Flags().StringVarP(&bucket, "bucket", "b", "", "aliyun-oss上的存储空间命名")
	cmdDelete.Flags().StringVarP(&object, "object", "o", "", "aliyun-oss上的文件对象命名")
	cmdDelete.Flags().StringVarP(&filepath, "filepath", "f", "", "指定需要上传文件或者目录的绝对路径")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func getBucketClient() *bu.Bucket {
	client, err := oss.New(config.Endpoint, config.AccessKeyID, config.AccessKeySecret)
	if err != nil {
		fmt.Println(err)
	}

	bucket := &bu.Bucket{
		Client: client,
	}
	return bucket
}

func getObjectClient() *ob.Object {
	client, err := oss.New(config.Endpoint, config.AccessKeyID, config.AccessKeySecret)
	if err != nil {
		fmt.Println(err)
	}

	object := &ob.Object{
		Client: client,
	}
	return object
}
