package bucket

import (
	"fmt"
	"testing"

	"oss-aliyun-cli/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

/*
	- 获取存储空间列表（List Bucket）
	- 创建存储空间（Create Bucket）
	- 删除存储空间（Delete Bucket）
*/
func getBucketClient() *Bucket {
	client, err := oss.New(config.Endpoint, config.AccessKeyID, config.AccessKeySecret)
	if err != nil {
		fmt.Println(err)
	}

	bucket := &Bucket{
		Client: client,
	}
	return bucket
}

func TestListBucket(t *testing.T) {
	bucket := getBucketClient()
	listBucket, err := bucket.ListBucket()
	if err != nil {
		t.Error(err)
	}
	t.Logf("My BUcket: %v", listBucket)
}

func TestCreateBucket(t *testing.T) {
	bucket := getBucketClient()
	tests := []struct {
		BucketName string
	}{
		{
			BucketName: "leeebucket",
		},
	}
	for _, test := range tests {
		err := bucket.CreateBucket(test.BucketName)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Create Bucket: %v success", test.BucketName)
		}

	}
}

func TestDeleteBucket(t *testing.T) {
	bucket := getBucketClient()
	tests := []struct {
		BucketName string
	}{
		{
			BucketName: "fafafadd",
		},
	}
	for _, test := range tests {
		err := bucket.DeleteBucket(test.BucketName)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Delete Bucket: %v success", test.BucketName)
		}
	}
}
