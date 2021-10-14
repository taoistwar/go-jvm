package classpath

import (
	"io/ioutil"
	"log"
	"os"
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
	_, err := os.Stat(dstFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, directoryEntry, nil
		}
		return nil, directoryEntry, err
	}

	data, err := ioutil.ReadFile(dstFile)
	log.Printf("load class:%s from file:%s", className, dstFile)
	return data, directoryEntry, err
}
