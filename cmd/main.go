package main

import (
	"fmt"
	"youtube-downloader/internal/server"
)

func main() {
	fmt.Println("Server started at http://localhost:8080/")
	server.Start()
}
