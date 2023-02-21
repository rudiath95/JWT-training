package main

import (
	"github.com/rudiath95/SWT-training/api"
)

func main() {
	server := api.NewAPIServer(":3000")
	server.Run()
}
