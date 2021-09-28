package classpath

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

type DirectoryEntry struct {
	absDir string
}

func newDirectoryEntry(path string) *DirectoryEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	data := DirectoryEntry{
		absDir: absDir,
	}
	return &data
}
func (directoryEntry *DirectoryEntry) String() string {
	return directoryEntry.absDir
}

func (directoryEntry *DirectoryEntry) readClass(className string) ([]byte, Entry, error) {
	dstFile := filepath.Join(directoryEntry.absDir, className)
	data, err := ioutil.ReadFile(dstFile)
	log.Printf("load class:%s from file:%s", className, dstFile)
	return data, directoryEntry, err
}
