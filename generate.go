//go:generate go run generate.go

package main

import (
	"git.sr.ht/~shulhan/ciigo"
)

func main() {
	ciigo.Generate("_content", "cmd/www-kilabit/static.go", "_content/template.gohtml")
}
