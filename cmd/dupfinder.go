package main

import (
	"flag"
	"os"

	"github.com/crerwin/dupfinder/duptools"
)

func main() {
	if len(os.Args) != 2 {
		println("invalid arguments.  Must specify one folder.")
	} else {
		flag.Parse()
		path := flag.Arg(0)
		duptools.FindDups(path)
	}
}
