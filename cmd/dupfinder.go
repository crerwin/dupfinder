package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/crerwin/dupfinder/duptools"
)

var filesByHash = make(map[[sha256.Size]byte]string)

func main() {
	if len(os.Args) != 2 {
		println("invalid arguments.  Must specify one folder.")
	} else {
		flag.Parse()
		path := flag.Arg(0)
		fmt.Println(len(os.Args), os.Args)
		inventoryFilesByName(path)
	}
}

func fileVisited(path string, f os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}
	if f.IsDir() {
		return nil
	}
	fmt.Printf("%s calculating hash: ", path)
	hash, _ := duptools.GetHash(path)
	fmt.Printf("%x\n", hash)
	var hashArray [sha256.Size]byte
	copy(hashArray[:], hash)
	if p, ok := filesByHash[hashArray]; ok {
		fmt.Printf("%q is a duplicate of %q\n", path, p)
	} else {
		filesByHash[hashArray] = path
	}
	return nil
}

func inventoryFilesByName(path string) {
	fmt.Println("Inventorying files in " + path)
	err := filepath.Walk(path, fileVisited)
	if err != nil {
		log.Fatal(err)
	}
}
