package config

const (
	Endpoint        string = "oss-cn-shanghai.aliyuncs.com"   // oss endpoint
	AccessKeyID     string = "LTAILC66xL8WBKrV"               // oss accessKeyId
	AccessKeySecret string = "O7J3LqnUXvtWZJ0XSCGFoPBHAGZYDL" // oss secret
	BucketName      string = "xxx"                            // oss bucket名称
	ObjectName      string = ""                               // oss远程目标地址
	WorkerCount     int    = 100                              // 设置最大并发数
	Suffix          string = ""                               // 筛选目录下需要上传的格式
)
