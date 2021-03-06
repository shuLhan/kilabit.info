package main

import (
	"log"

	"git.sr.ht/~shulhan/ciigo"
	"github.com/shuLhan/share/lib/memfs"
)

var memfsContent *memfs.MemFS

func main() {
	log.SetFlags(0)
	err := ciigo.Serve(memfsContent, "_content", ":7000", "_content/template.gohtml")
	if err != nil {
		log.Fatal(err)
	}
}
