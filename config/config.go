package config

const (
	Endpoint        string = "XXXX"   // oss endpoint
	AccessKeyID     string = "XXXXX"  // oss accessKeyId
	AccessKeySecret string = "XXXXXX" // oss secret
	BucketName      string = "xxx"    // oss bucket名称
	ObjectName      string = ""       // oss远程目标地址
	WorkerCount     int    = 100      // 设置最大并发数
	Suffix          string = ""       // 筛选目录下需要上传的格式
)
