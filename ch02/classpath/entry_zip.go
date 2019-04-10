package classpath

import (
	"archive/zip"
	"github.com/pkg/errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath: absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// zip 每次打开、关闭 性能较差
	// TODO 查看程序的优化
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, err
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

// zipEntry String
func (self *ZipEntry) String() string {
	return self.absPath
}
