package duptools

import "crypto/sha256"

type FileCollection struct {
	filesByName map[string][]string            // key: filename value: full path
	filesByHash map[[sha256.Size]byte][]string // key: hash value: full path
}

func NewFileCollection() *FileCollection {
	// use to create a new FileCollection and build the maps
	var fc FileCollection
	fc.filesByName = make(map[string][]string)
	fc.filesByHash = make(map[[sha256.Size]byte][]string)
	return &fc
}

func AddName(fc *FileCollection, path string, filename string) bool {
	// add file to filesByName and return True if there's already a
	// file (or many files) with that name, otherwise return false
	fc.filesByName[filename] = append(fc.filesByName[filename], path)
	if len(fc.filesByName[filename]) > 1 {
		return true
	} else {
		return false
	}
}
