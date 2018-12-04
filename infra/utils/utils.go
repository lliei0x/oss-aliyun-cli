package utils

import (
	"fmt"
	"os"
	"strings"
)

func PathSplitToFileName(path string) string {
	split := strings.Split(path, `\`)
	splitLength := len(split)
	fileName := split[splitLength-1 : splitLength]
	return fileName[0]
}

func CheckFileOrDir(path string) (bool bool, err error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

func Upload(path string) {
	//分析输入是目录还是文件,以及文件的合法性
	isDir, err := CheckFileOrDir(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("目录或文件不存在，请核对后再试！")
		} else {
			fmt.Println(err)
		}
	} else {
		if !isDir {
			fmt.Printf("准备上传单个文件！%s", path)
			fmt.Println("")
			//上传单个
			UploadOne(path)
		} else {
			fmt.Printf("准备批量上传文件！%s", path)
			fmt.Println("")
			//并发上传
			UploadMany(path)
		}
	}
}
