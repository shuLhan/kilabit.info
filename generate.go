//go:generate go run generate.go

package main

import (
	"github.com/shuLhan/ciigo"
)

func main() {
	ciigo.Generate("_content", "cmd/www-kilabit/static.go", "_content/template.gohtml")
}
