package main

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"leeif.me/Go-utils/oss-aliyun/config"
	"leeif.me/Go-utils/oss-aliyun/domain"
)

const (
	path       = `D:\Server\Go\src\leeif.me\Go-utils\oss-aliyun\main.go`
	bucketName = "leeifme"
	objectName = ""
)

func main() {
	client, err := oss.New(config.Endpoint, config.AccessKeyID, config.AccessKeySecret)
	if err != nil {
		fmt.Println(err)
	}
	bucket := &domain.Bucket{
		Client: client,
	}
	bucket.ListBucket()

	object := &domain.Object{
		Client: client,
	}
	// fileinfo, err := os.Stat(path)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(fileinfo.Name())  //获取文件名
	// fmt.Println(fileinfo.IsDir()) //判断是否是目录，返回bool类型
	err = object.PutObject(bucketName, path)
	if err != nil {
		fmt.Println(err)
	}
	// err = object.GetObject(bucketName)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
