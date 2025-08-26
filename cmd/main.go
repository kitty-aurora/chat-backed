package main

import (
	"chat-backed/internal/server"
)

func main() {
	r := server.NewHTTPServer()
	r.Run(":8081")
}
