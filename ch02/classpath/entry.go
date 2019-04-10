package classpath

import (
	"os"
	"strings"
)

// 路径分隔符
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// 寻找和加载类文件
	readClass(className string) ([] byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, "*") {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") || strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)

	}

	return newDirEntry(path)
}