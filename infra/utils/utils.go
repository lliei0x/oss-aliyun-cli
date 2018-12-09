package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

// PathSplitToFileName 分割路径
func PathSplitToFileName(path string) string {
	split := strings.Split(path, `\`)
	splitLength := len(split)
	fileName := split[splitLength-1 : splitLength]
	return fileName[0]
}

// CheckFileOrDir 判断路径是文件还是文件夹
func CheckFileOrDir(path string) (bool bool, err error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

//GetAllFiles 获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(dirPth string, temps []string, suffix string) (files []string, err error) {
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
			files, _ = GetAllFiles(dirPth+`\`+fileInfo.Name(), files, suffix)
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
