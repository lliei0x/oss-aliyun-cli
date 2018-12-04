package object

import (
	"fmt"
	"testing"

	"oss-aliyun-cli/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

/*
	- 上传文件（Put Object）
	- 下载文件 (Get Object)
	- 获取文件列表（List Objects）
	- 删除文件 (Delete Object)
*/
func getObjectClient() *Object {
	client, err := oss.New(config.Endpoint, config.AccessKeyID, config.AccessKeySecret)
	if err != nil {
		fmt.Println(err)
	}

	object := &Object{
		Client: client,
	}
	return object
}

func TestPutObject(t *testing.T) {
	object := getObjectClient()
	tests := []struct {
		bucketName string
		path       string
	}{
		{
			bucketName: "leeebucket",
			path:       `D:\Server\Go\src\oss-aliyun-cli\main.go`,
		},
	}
	for _, test := range tests {
		err := object.PutObject(test.bucketName, test.path)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Put Object: %v success", test.path)
		}
	}

}

func TestGetObject(t *testing.T) {
	object := getObjectClient()
	tests := []struct {
		bucketName string
		Path       string
	}{
		{
			bucketName: "leeebucket",
			Path:       "",
		},
	}
	for _, test := range tests {
		err := object.GetObject(test.bucketName)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Put Object: %v success", test.Path)
		}
	}

}

func TestListObject(t *testing.T) {
	object := getObjectClient()
	tests := []struct {
		bucketName string
	}{
		{
			bucketName: "leeebucket",
		},
	}
	for _, test := range tests {
		listObject, err := object.ListObject(test.bucketName)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("My Object: %v", listObject)
		}
	}

}

func TestDeleteObject(t *testing.T) {
	object := getObjectClient()
	tests := []struct {
		bucketName string
		objectKeys []string
	}{
		{
			bucketName: "leeebucket",
			objectKeys: []string{"main.go"},
		},
	}
	for _, test := range tests {
		err := object.DeleteObject(test.bucketName, test.objectKeys)
		if err != nil {
			t.Error(err)
		}
		// else {
		// 	t.Logf("My Object: %v", listObject)
		// }
	}
}
