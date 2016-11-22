package duptools

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var fileCollection = NewFileCollection()

func FindDups(path string) {
	// exposed function
	inventoryFilesByName(path)
}

func inventoryFilesByName(path string) {
	// walk recursively from the given path,
	// executing fileVisited for each file found
	fmt.Println("Inventorying files in " + path)
	err := filepath.Walk(path, fileVisited)
	if err != nil {
		log.Fatal(err)
	}
}

func fileVisited(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}
	if info.IsDir() {
		return nil
	}

	if AddName(fileCollection, path, info.Name()) {
		fmt.Println(path, " is a possible duplicate.")
	}

	if AddHash(fileCollection, path) {
		fmt.Println(path, " is a duplicate file.")
	}

	return nil
}

func hashPossibleDups() {

}

func getHash(filePath string) ([]byte, error) {
	var result []byte
	file, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return result, err
	}
	return hash.Sum(result), nil
}
