package main

import (
	"log"

	"git.sr.ht/~shulhan/ciigo"
)

func main() {
	log.SetFlags(0)
	ciigo.Serve("_content", ":7000", "_content/template.gohtml")
}
