//go:generate go run generate.go

package main

import (
	"github.com/shuLhan/ciigo"
)

func main() {
	ciigo.Generate("./content", "cmd/kilabit/static.go")
}
