package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir	string // 绝对路径
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir:absDir}
}

func (self *DirEntry) readClass(className string)([] byte, Entry, error)  {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (selef * DirEntry) String() string {
	return selef.absDir
}
