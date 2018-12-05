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
		fmt.Println("")
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
	var temps []string
	files, err := getAllFiles(path, temps)
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
			fmt.Println("")
		}
		go func() { done <- true }()
	}
}

//获取指定目录下的所有文件,包含子目录下的文件
func getAllFiles(dirPth string, temps []string) (files []string, err error) {
	// 拼接上一个目录下的文件路径
	for _, temp := range temps {
		files = append(files, temp)
	}

	fileInfos, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			// 目录, 递归遍历
			files, _ = getAllFiles(dirPth+`\`+fileInfo.Name(), files)
		} else {
			// 过滤指定格式
			if suffix != "" {
				ok := strings.HasSuffix(fileInfo.Name(), suffix)
				if ok {
					files = append(files, dirPth+`\`+fileInfo.Name())
				}
			} else {
				files = append(files, dirPth+`\`+fileInfo.Name())
			}
		}
	}
	return files, nil
}
