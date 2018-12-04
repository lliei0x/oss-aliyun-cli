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
		t.Error(err)
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
		bucketName string
	}{
		{
			bucketName: "leeebucket",
		},
	}
	for _, test := range tests {
		err := bucket.CreateBucket(test.bucketName)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Create Bucket: %v success", test.bucketName)
		}

	}
}

func TestDeleteBucket(t *testing.T) {
	bucket := getBucketClient()
	tests := []struct {
		bucketName string
	}{
		{
			bucketName: "leeebucket",
		},
	}
	for _, test := range tests {
		err := bucket.DeleteBucket(test.bucketName)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Delete Bucket: %v success", test.bucketName)
		}
	}
}
