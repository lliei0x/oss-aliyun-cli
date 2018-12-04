package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"oss-aliyun-cli/config"
	"oss-aliyun-cli/domain/object"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Worker struct {
	in   chan string
	done chan bool
}

const (
	workerCount = 10 //设置最大并发数
	suffix      = ""
)

func getObjectClient() *object.Object {
	client, err := oss.New(config.Endpoint, config.AccessKeyID, config.AccessKeySecret)
	if err != nil {
		fmt.Println(err)
	}

	object := &object.Object{
		Client: client,
	}
	return object
}

//上传单个文件
func UploadOne(path string) {
	// 创建OSSClient实例
	object := getObjectClient()
	err := object.PutObject("leeebucket", path)
	if err != nil {
		handleError(err)
	} else {
		fmt.Printf("Put Object: %v success", path)
	}
}
func createWorker(id int) Worker {
	w := Worker{
		in:   make(chan string),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}

//并发上传目录下所有文件
func UploadMany(path string) {
	files, err := getAllFiles(path)
	if err != nil {
		handleError(err)
	}
	fileCount := len(files)
	if fileCount == 0 {
		handleError(errors.New("目录下没有指定文件，请重试！"))
	}
	var workers [workerCount]Worker
	for i := 0; i < fileCount; i++ {
		workers[i] = createWorker(i)
	}
	for i := 0; i < fileCount; i++ {
		workers[i].in <- files[i]
	}
	for i := 0; i < fileCount; i++ {
		<-workers[i].done
	}
}

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

func doWork(id int, c chan string, done chan bool) {
	object := getObjectClient()
	for n := range c {
		fmt.Printf("worker : %d, uploading file %v \n", id, n)
		err := object.PutObject("leeebucket", n)
		if err != nil {
			handleError(err)
		} else {
			fmt.Printf("Put Object: %v success", n)
		}
		go func() { done <- true }()
	}
}

//获取指定目录下的所有文件,包含子目录下的文件
func getAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			getAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			if suffix != "" {
				ok := strings.HasSuffix(fi.Name(), suffix)
				fmt.Println("111111111")
				if ok {
					files = append(files, dirPth+PthSep+fi.Name())
				}
			} else {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	return files, nil
}
