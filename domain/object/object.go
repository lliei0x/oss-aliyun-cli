package object

import (
	"oss-aliyun-cli/infra/utils"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Object struct {
	Client *oss.Client
}

// PutObject 上传文件
func (object *Object) PutObject(bucketName string, objectName string, path string) error {
	bucket, err := object.Client.Bucket(bucketName)
	if err != nil {
		return err
	}
	fileName := utils.PathSplitToFileName(path)
	err = bucket.PutObjectFromFile(objectName+fileName, path)
	if err != nil {
		return err
	}
	return nil
}

// GetObject 下载文件
func (object *Object) GetObject(bucketName string, objectKey string, loaclPath string) error {
	bucket, err := object.Client.Bucket(bucketName)
	if err != nil {
		return err
	}
	err = bucket.GetObjectToFile(objectKey, loaclPath)
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
func (object *Object) DeleteObject(bucketName string, objectKey string) error {
	bucket, err := object.Client.Bucket(bucketName)
	if err != nil {
		return err
	}
	err = bucket.DeleteObject(objectKey)
	if err != nil {
		return err
	}
	return nil
}
