//go:generate go run generate.go

package main

import (
	"log"

	"git.sr.ht/~shulhan/ciigo"
	"github.com/shuLhan/share/lib/memfs"
)

func main() {
	opts := ciigo.EmbedOptions{
		ConvertOptions: ciigo.ConvertOptions{
			Root:         "_content",
			HtmlTemplate: "_content/template.gohtml",
		},
		EmbedOptions: memfs.EmbedOptions{
			PackageName: "main",
			VarName:     "memfsContent",
			GoFileName:  "cmd/www-kilabit/memfs_content.go",
		},
	}
	err := ciigo.GoEmbed(&opts)
	if err != nil {
		log.Fatal(err)
	}
}
