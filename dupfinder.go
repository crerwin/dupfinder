package dupfinder

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var files = make(map[[md5.Size]byte]string)

func getHash(filePath string) ([]byte, error) {
	var result []byte
	file, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return result, err
	}
	return hash.Sum(result), nil
}

func fileVisited(path string, f os.FileInfo, err error) error {
	fmt.Printf("%s calculating hash: ", path)
	hash, _ := getHash(path)
	fmt.Printf("%x", hash)
	return nil
}

func inventoryFilesByName(path string) {
	fmt.Println("Inventorying files in " + path)
	err := filepath.Walk(path, fileVisited)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	path := flag.Arg(0)
	inventoryFilesByName(path)
}
