package main

import (
	"flag"
	"log"
	"strings"

	psh "github.com/platformsh/config-reader-go/v2"

	"git.sr.ht/~shulhan/ciigo"
	"github.com/shuLhan/share/lib/memfs"
)

const (
	envNameDev        = "dev"
	envNamePlatformsh = "platform.sh"
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

		pshConfig *psh.RuntimeConfig
		flagEnv   string
		err       error
	)

	log.SetFlags(0)

	flag.StringVar(&flagEnv, "env", "", "set the environment to run")
	flag.Parse()

	if len(flagEnv) > 0 {
		flagEnv = strings.ToLower(flagEnv)
	}

	if flagEnv == envNamePlatformsh {
		pshConfig, err = psh.NewRuntimeConfig()
		if err != nil {
			log.Fatal("Not in a Platform.sh environment.")
		}
		port = pshConfig.Port()
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
