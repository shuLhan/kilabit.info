package main

import (
	"flag"
	"log"
	"strings"

	"git.sr.ht/~shulhan/ciigo"
	"github.com/shuLhan/share/lib/memfs"
)

const (
	envNameDev = "dev"
)

var memfsContent *memfs.MemFS

func main() {
	var (
		port      = "7000"
		serveOpts = ciigo.ServeOptions{
			ConvertOptions: ciigo.ConvertOptions{
				Root:         "_content",
				HtmlTemplate: "_content/template.gohtml",
			},
			Mfs: memfsContent,
		}

		flagEnv string
		err     error
	)

	log.SetFlags(0)

	flag.StringVar(&flagEnv, "env", "", "set the environment to run")
	flag.Parse()

	if len(flagEnv) > 0 {
		flagEnv = strings.ToLower(flagEnv)
	}

	serveOpts.Address = ":" + port

	if flagEnv == envNameDev {
		serveOpts.IsDevelopment = true
	}

	err = ciigo.Serve(&serveOpts)
	if err != nil {
		log.Fatal(err)
	}
}
