package utils

import (
	"strings"
)

func PathSplitToFileName(path string) string {
	split := strings.Split(path, `\`)
	splitLength := len(split)
	fileName := split[splitLength-1 : splitLength]
	return fileName[0]
}
