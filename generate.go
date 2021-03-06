//go:generate go run generate.go

package main

import (
	"log"

	"git.sr.ht/~shulhan/ciigo"
)

func main() {
	opts := ciigo.GenerateOptions{
		Root:           "_content",
		HTMLTemplate:   "_content/template.gohtml",
		GenPackageName: "main",
		GenVarName:     "memfsContent",
		GenGoFileName:  "cmd/www-kilabit/memfs_content.go",
	}
	err := ciigo.Generate(&opts)
	if err != nil {
		log.Fatal(err)
	}
}
