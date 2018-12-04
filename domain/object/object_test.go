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
		Bucket string
		Path   string
	}{
		{
			Bucket: "",
			Path:   "",
		},
	}
	for _, test := range tests {
		object.PutObject(test.Bucket, test.Path)
	}

}
