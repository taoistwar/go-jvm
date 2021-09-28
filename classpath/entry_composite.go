package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) *CompositeEntry {
	var compositeEntry CompositeEntry = []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return &compositeEntry
}
func (compositeEntry *CompositeEntry) String() string {
	items := make([]string, len(*compositeEntry))
	for i, entry := range *compositeEntry {
		items[i] = entry.String()
	}
	return strings.Join(items, pathListSeparator)
}

func (compositeEntry *CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range *compositeEntry {
		data, from, err := entry.readClass(className)
		if err != nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}
