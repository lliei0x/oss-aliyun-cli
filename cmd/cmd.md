```sh
/*
	- 获取存储空间列表（List Bucket）
	- 创建存储空间（Create Bucket）
	- 删除存储空间（Delete Bucket）
*/

/*
	- 上传文件（Put Object）
	- 下载文件 (Get Object)
	- 获取文件列表（List Objects）
	- 删除文件 (Delete Object)
*/

oss 
[object]
- up 
	- object         // oss up -b leeifme -o leeifme -f oss/main.go
- download 
	- object         // oss download -b leeifme -o leeifme 

[object|bucket]
- get 
	- bucket         // oss get -b leeifme 
	- object         // oss get -b leeifme -o leeifme
- create
	- bucket         // oss create -b leeifme
	- object	 	 // oss create -b leeifme -o leeifme
- delete 
	- bucket         // oss delete -b leeifme
	- object         // oss delete -b leeifme -o leeifme


```