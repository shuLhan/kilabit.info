package main

import (
	"log"

	"github.com/shuLhan/ciigo"
)

func main() {
	log.SetFlags(0)
	ciigo.Serve("_content", ":7000", "_content/template.gohtml")
}
