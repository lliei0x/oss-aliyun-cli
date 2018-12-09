package upload

import (
	"errors"
	"fmt"
	"os"

	"oss-aliyun-cli/domain/object"
	"oss-aliyun-cli/infra/utils"
)

type Worker struct {
	in   chan string
	done chan bool
}

const (
	workerCount = 10 //设置最大并发数
	suffix      = ""
)

func Upload(path string, suffix string, bucketName string, objectKey string, oClient *object.Object) error {

	var (
		err   error
		isDir bool
	)
	//分析输入是目录还是文件,以及文件的合法性
	isDir, err = utils.CheckFileOrDir(path)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("目录或文件不存在，请核对后再试！")
		}
	} else {
		if !isDir {
			fmt.Printf("准备上传单个文件！%s", path)
			fmt.Println("")
			//上传单个
			err = uploadOne(path, bucketName, objectKey, oClient)
		} else {
			fmt.Printf("准备批量上传文件！%s", path)
			fmt.Println("")
			//并发上传
			err = uploadMany(path, suffix, bucketName, objectKey, oClient)
		}
	}
	if err != nil {
		return err
	}
	return nil
}

//上传单个文件
func uploadOne(path string, bucketName string, objectKey string, object *object.Object) error {
	err := object.PutObject(bucketName, objectKey, path)
	if err != nil {
		return err
	} else {
		fmt.Printf("Put Object: %v success", path)
		fmt.Println("")
	}
	return nil
}

//并发上传目录下所有文件
func uploadMany(path string, suffix string, bucketName string, objectKey string, oClient *object.Object) error {
	var temps []string
	files, err := utils.GetAllFiles(path, temps, suffix)
	if err != nil {
		return err
	}
	fileCount := len(files)
	if fileCount == 0 {
		return errors.New("目录下没有指定文件，请重试！")
	}
	var workers [workerCount]Worker
	for i := 0; i < fileCount; i++ {
		workers[i], err = createWorker(i, bucketName, objectKey, oClient)
		if err != nil {
			return err
		}
	}
	for i := 0; i < fileCount; i++ {
		workers[i].in <- files[i]
	}
	for i := 0; i < fileCount; i++ {
		<-workers[i].done
	}
	return nil
}

func createWorker(id int, bucketName string, objectKey string, oClient *object.Object) (Worker, error) {
	var err error
	w := Worker{
		in:   make(chan string),
		done: make(chan bool),
	}
	go func(err error) error {
		err = doWork(id, w.in, w.done, bucketName, objectKey, oClient)
		if err != nil {
			return err
		}
		return nil
	}(err)
	return w, err
}

func doWork(id int, c chan string, done chan bool, bucketName string, objectKey string, oClient *object.Object) error {
	for n := range c {
		fmt.Printf("worker : %d, uploading file %v \n", id, n)
		err := oClient.PutObject(bucketName, objectKey, n)
		if err != nil {
			return err
		} else {
			fmt.Printf("Put Object: %v success", n)
			fmt.Println("")
		}
		go func() { done <- true }()
	}
	return nil
}
