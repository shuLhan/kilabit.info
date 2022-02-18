package main

import (
	"flag"
	"log"
	"strings"

	"git.sr.ht/~shulhan/ciigo"
	"github.com/shuLhan/share/lib/memfs"
)

var memfsContent *memfs.MemFS

func main() {
	var (
		flagEnv string
	)

	log.SetFlags(0)

	flag.StringVar(&flagEnv, "env", "", "set the environment to run")
	flag.Parse()

	if len(flagEnv) > 0 {
		flagEnv = strings.ToLower(flagEnv)
	}

	serveOpts := &ciigo.ServeOptions{
		ConvertOptions: ciigo.ConvertOptions{
			Root:         "_content",
			HtmlTemplate: "_content/template.gohtml",
		},
		Mfs:     memfsContent,
		Address: "127.0.0.1:7000",
	}

	if flagEnv == "dev" {
		serveOpts.IsDevelopment = true
	}

	err := ciigo.Serve(serveOpts)
	if err != nil {
		log.Fatal(err)
	}
}
