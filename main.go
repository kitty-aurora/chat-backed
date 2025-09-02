package main

import (
	"chat-backed/router"
	"fmt"
)

func main() {
	r := router.SetupRouter()

	port := ":8080"
	fmt.Println("Server running at http://localhost" + port)
	r.Run(port)
}
