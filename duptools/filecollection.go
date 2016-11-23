package duptools

import (
	"crypto/sha256"
	"fmt"
)

type FileCollection struct {
	filesByName map[string][]string            // key: filename value: full path
	filesByHash map[[sha256.Size]byte][]string // key: hash value: full path
}

func newFileCollection() *FileCollection {
	// use to create a new FileCollection and build the maps
	var fc FileCollection
	fc.filesByName = make(map[string][]string)
	fc.filesByHash = make(map[[sha256.Size]byte][]string)
	return &fc
}

func addName(fc *FileCollection, path string, filename string) bool {
	// add file to filesByName and return True if there's already a
	// file (or many files) with that name, otherwise return false
	fc.filesByName[filename] = append(fc.filesByName[filename], path)
	if len(fc.filesByName[filename]) > 1 {
		return true
	} else {
		return false
	}
}

func addHash(fc *FileCollection, path string) bool {
	// add file keyed on its hash to filesByHash and return True if
	// there's already a file (or many files) with that hash, otherwise
	// return false
	fmt.Print("hashing ", path, ": ")
	hash, _ := getHash(path)
	fmt.Printf("%x\n", hash)
	var hashArray [sha256.Size]byte
	copy(hashArray[:], hash)

	fc.filesByHash[hashArray] = append(fc.filesByHash[hashArray], path)
	if len(fc.filesByHash[hashArray]) > 1 {
		return true
	} else {
		return false
	}
}

func getNameDups(fc *FileCollection) []string {
	var nameDups []string
	for _, paths := range fc.filesByName {
		if len(paths) > 1 {
			nameDups = append(nameDups, paths...)
		}
	}
	return nameDups
}

func getHashDups(fc *FileCollection) []string {
	var hashDups []string
	for _, paths := range fc.filesByHash {
		if len(paths) > 1 {
			hashDups = append(hashDups, paths...)
		}
	}
	return hashDups
}

func getFileCount(fc *FileCollection) int {
	var count int
	for _, paths := range fc.filesByName {
		count += len(paths)
	}
	return count
}
