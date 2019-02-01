## OSS-Aliyun-CLI

![](https://travis-ci.org/wuxiaoxiaoshen/2019-daily.svg?branch=master)
![leeifme on github](https://img.shields.io/badge/github-@leeifme-red.svg)](https://github.com/leeifme)
![package version](https://img.shields.io/badge/package-v0.1.0-blue.svg)](https://github.com/leeifme/oss-aliyun-cli)
![language by golang](https://img.shields.io/badge/language-@golang-green.svg)](https://github.com/leeifme/oss-aliyun-cli)

> aliyun-oss 命令行工具 :man_technologist:

### 支持

- [x] 获取存储空间列表（List Bucket）
- [x] 创建存储空间（Create Bucket）
- [x] 删除存储空间（Delete Bucket）

- [x] 上传文件（Put Object）
- [x] 下载文件 (Get Object)
- [x] 获取文件列表（List Objects）
- [x] 删除文件 (Delete Object)

### 主要功能
- 上传文件（Put Object）
    - 单个文件 && 读取目录及子目录下文件
    - 并发上传多个文件
    - 支持过滤筛选文件
  
### 使用
```sh
针对 aliyun-oss 常规操作的命令行工具

Usage:
  oss [command]

Available Commands:
  create      新建存储空间(bucket) 或者 文件对象(object)
  delete      删除指定存储空间(bucket) 或者 文件对象(object)
  download    下载指定的文件对象(object)
  get         获取存储空间(bucket) 或者 文件对象(object) 列表
  help        Help about any command
  up          上传指定文件 或者 并发上传指定目录以及子目录下符合过滤条件的文件

Flags:
  -h, --help   help for oss

Use "oss [command] --help" for more information about a command.
```
**注：** 可 `clone` 源码后自行编译后体验

### 最后
> 感谢
- [aliyun-oss-go-sdk](https://github.com/aliyun/aliyun-oss-go-sdk)
- [cobra](https://github.com/spf13/cobra)
- [go-oss-uploads](https://github.com/363182860/go-oss-uploads)