package imgconv

import "os"

func GetOutputPath(path, dst string) string {
	return getOutputPath(path, dst)
}

func IsTargetFile(info os.FileInfo, path, format string) bool {
	return isTargetFile(info, path, format)
}
