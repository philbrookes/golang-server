// main
package main

import (
	"server"
)

func main() {
	srv := server.NewServer(8080)
	srv.Start()
}
