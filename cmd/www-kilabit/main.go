package main

import (
	"log"

	"github.com/shuLhan/ciigo"
)

func main() {
	log.SetFlags(0)
	srv := ciigo.NewServer("", ":7000", "templates/html.tmpl")
	srv.Start()
}
