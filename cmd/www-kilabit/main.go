package main

import (
	"flag"
	"log"
	"strings"

	"git.sr.ht/~shulhan/ciigo"
	"github.com/shuLhan/share/lib/memfs"
)

const (
	cmdNameEmbed = `embed`
)

var memfsContent *memfs.MemFS

func main() {
	var (
		convertOpts = ciigo.ConvertOptions{
			Root:         `_content`,
			HTMLTemplate: `_content/template.gohtml`,
		}
		serveOpts = ciigo.ServeOptions{
			ConvertOptions: convertOpts,
			Mfs:            memfsContent,
			Address:        `127.0.0.1:7000`,
		}

		cmd string
		err error
	)

	flag.BoolVar(&serveOpts.IsDevelopment, "dev", false, "Run in development mode")
	flag.Parse()

	cmd = strings.ToLower(flag.Arg(0))

	switch cmd {
	case cmdNameEmbed:
		var embedOpts = ciigo.EmbedOptions{
			ConvertOptions: convertOpts,
			EmbedOptions: memfs.EmbedOptions{
				PackageName: `main`,
				VarName:     `memfsContent`,
				GoFileName:  `cmd/www-kilabit/memfs_content.go`,
			},
		}
		err = ciigo.GoEmbed(&embedOpts)
		if err != nil {
			log.Fatal(err)
		}

	default:
		err = ciigo.Serve(&serveOpts)
		if err != nil {
			log.Fatal(err)
		}
	}
}
