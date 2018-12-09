package upload

import (
	"testing"
)

func TestUpload(t *testing.T) {
	tests := []struct {
		path string
	}{
		// {
		// 	path: `D:\Server\Go\src\oss-aliyun-cli\main.go`,
		// },
		{
			path: `D:\Server\Go\src\oss-aliyun-cli\domain`,
		},
		// {
		// 	path: `D:\Server\Go\src\oss-aliyun-cli\domain\object`,
		// },
	}

	for _, test := range tests {
		Upload(test.path)
	}
}
