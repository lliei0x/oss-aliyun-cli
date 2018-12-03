package domain

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

/*
	- 获取存储空间列表（List Bucket）
	- 创建存储空间（Create Bucket）
	- 删除存储空间（Delete Bucket）
*/
type Bucket struct {
	Client *oss.Client
}

// ListBucket 获取存储空间列表
func (bucket *Bucket) ListBucket() {
	lbr, err := bucket.Client.ListBuckets()
	if err != nil {
		return
	}
	for _, bucket := range lbr.Buckets {
		fmt.Println("my buckets:", bucket.Name)
	}
}

// CreateBucket 创建存储空间
func (bucket *Bucket) CreateBucket(bucketName string) error {
	err := bucket.Client.CreateBucket(bucketName)
	return err
}

// DeleteBucket 删除存储空间
func (bucket *Bucket) DeleteBucket(bucketName string) error {
	err := bucket.Client.DeleteBucket(bucketName)
	return err
}
