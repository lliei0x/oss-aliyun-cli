package object

import (
	"fmt"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

/*
	- 上传文件（Put Object）
	- 下载文件 (Get Object)
	- 获取文件列表（List Objects）
	- 删除文件 (Delete Object)
*/

type Object struct {
	Client *oss.Client
}

func (object *Object) PutObject(bucketName string, path string) error {
	bucket, err := object.Client.Bucket(bucketName)
	if err != nil {
		return err
	}
	objectName := ""
	split := strings.Split(path, `\`)
	splitLength := len(split)
	fileName := split[splitLength-1 : splitLength]
	signedURL, err := bucket.SignURL(objectName+fileName[0], oss.HTTPPut, 60)
	if err != nil {
		return err
	}
	fmt.Println(signedURL)
	err = bucket.PutObjectFromFile(objectName+fileName[0], path)
	if err != nil {
		return err
	}
	return nil
}

func (object *Object) GetObject(bucketName string) error {
	bucket, err := object.Client.Bucket(bucketName)
	if err != nil {
		return err
	}
	err = bucket.GetObjectToFile("blog/", "blog/")
	// bucket.GetObjectWithURL(signedURL, options)
	if err != nil {
		return err
	}
	return nil
}

func (object *Object) ListObject(bucketName string) error {
	bucket, err := object.Client.Bucket(bucketName)
	if err != nil {
		return err
	}
	lor, err := bucket.ListObjects()
	if err != nil {
		return err
	}
	for _, object := range lor.Objects {
		// if strings.Split(object.Key, )
		fmt.Println(object.Key)
	}
	return nil
}
