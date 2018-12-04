package object

import (
	"fmt"
	"strings"
	// "oss-aliyun-cli/infra/utils"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Object struct {
	Client *oss.Client
}

// PutObject 上传文件
func (object *Object) PutObject(bucketName string, path string) error {
	bucket, err := object.Client.Bucket(bucketName)
	if err != nil {
		return err
	}
	// fileName := utils.PathSplitToFileName(path)
	split := strings.Split(path, `\`)
	splitLength := len(split)
	fileName := split[splitLength-1 : splitLength]
	err = bucket.PutObjectFromFile(fileName[0], path)
	if err != nil {
		return err
	}
	return nil
}

// GetObject 下载文件
func (object *Object) GetObject(bucketName string) error {
	bucket, err := object.Client.Bucket(bucketName)
	if err != nil {
		return err
	}
	err = bucket.GetObjectToFile("blog/", "blog/")
	if err != nil {
		return err
	}
	return nil
}

// ListObject 获取文件列表
func (object *Object) ListObject(bucketName string) ([]string, error) {
	bucket, err := object.Client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	lor, err := bucket.ListObjects()
	if err != nil {
		return nil, err
	}
	var listObject []string
	for _, object := range lor.Objects {
		listObject = append(listObject, object.Key)
	}
	return listObject, nil
}

// DeleteObject 删除文件
func (object *Object) DeleteObject(bucketName string, objectKeys []string) error {
	bucket, err := object.Client.Bucket(bucketName)
	if err != nil {
		return err
	}
	out, err := bucket.DeleteObjects(objectKeys)
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}
