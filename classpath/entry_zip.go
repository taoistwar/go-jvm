package classpath

import (
	"archive/zip"
	"io/ioutil"
	"log"
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
	data := ZipEntry{
		absPath: absPath,
	}
	return &data
}
func (zipEntry *ZipEntry) String() string {
	return zipEntry.absPath
}

func (zipEntry *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(zipEntry.absPath)
	if err != nil {
		return nil, nil, err
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
			log.Printf("load class:(%s) from zip:(%s)", className, filepath.Join(zipEntry.absPath, f.Name))
			return data, zipEntry, nil
		}
	}
	return nil, zipEntry, nil
}
