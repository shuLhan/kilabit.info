package main

import (
	"log"

	"git.sr.ht/~shulhan/ciigo"
	"github.com/shuLhan/share/lib/memfs"
)

var memfsContent *memfs.MemFS

func main() {
	log.SetFlags(0)
	serveOpts := &ciigo.ServeOptions{
		ConvertOptions: ciigo.ConvertOptions{
			Root:         "_content",
			HtmlTemplate: "_content/template.gohtml",
		},
		Address: "127.0.0.1:7000",
		Mfs:     memfsContent,
	}

	err := ciigo.Serve(serveOpts)
	if err != nil {
		log.Fatal(err)
	}
}
