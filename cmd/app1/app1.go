package main

import (
	"app1/pkg/server"
)

func main() {

	server := server.New()
	server.Start("Server starting ...")

}
