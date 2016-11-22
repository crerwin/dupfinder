package duptools

import "crypto/sha256"

type FileCollection struct {
	filesByName map[string]string
	filesByHash map[[sha256.Size]byte]string
}

func NewFileCollection() *FileCollection {
	// use to create a new FileCollection and build the maps
	var fc FileCollection
	fc.filesByName = make(map[string]string)
	fc.filesByHash = make(map[[sha256.Size]byte]string)
	return &fc
}

func CheckName(fc *FileCollection, path string, filename string) bool {
	if _, ok := fc.filesByName[filename]; ok {
		return true
	} else {
		fc.filesByName[filename] = path
		return false
	}
}
