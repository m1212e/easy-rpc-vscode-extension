package main

import "erpcLanguageServer/server"

func main() {
	server := server.NewServer()

	server.Run()
}
