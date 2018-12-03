package config

const (
	Endpoint        string = "oss-cn-hangzhou.aliyuncs.com" // oss endpoint
	AccessKeyID     string = "<AccessKeyId>"                // oss accessKeyId
	AccessKeySecret string = "xxx"                          // oss secret
	BucketName      string = "xxx"                          // oss bucket名称
	ObjectName      string = ""                             // oss远程目标地址
	WorkerCount     int    = 100                            // 设置最大并发数
	Suffix          string = ""                             // 筛选目录下需要上传的格式
)
