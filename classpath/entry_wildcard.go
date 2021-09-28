package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWillcardEntry(path string) *CompositeEntry {
	baseDir := path[:len(path)-1] // remove *
	var compositeEntry CompositeEntry = []Entry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
			strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
			entry := newZipEntry(path)
			compositeEntry = append(compositeEntry, entry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn)

	return &compositeEntry
}
