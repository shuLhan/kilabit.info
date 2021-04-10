//go:generate go run generate.go

package main

import (
	"log"

	"git.sr.ht/~shulhan/ciigo"
)

func main() {
	opts := ciigo.GenerateOptions{
		ConvertOptions: ciigo.ConvertOptions{
			Root:         "_content",
			HtmlTemplate: "_content/template.gohtml",
		},
		GenPackageName: "main",
		GenVarName:     "memfsContent",
		GenGoFileName:  "cmd/www-kilabit/memfs_content.go",
	}
	err := ciigo.Generate(&opts)
	if err != nil {
		log.Fatal(err)
	}
}
