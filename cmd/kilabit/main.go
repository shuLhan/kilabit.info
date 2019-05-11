package main

import (
	"github.com/shuLhan/ciigo"
)

func main() {
	srv := ciigo.NewServer(":7000")
	srv.Start()
}
