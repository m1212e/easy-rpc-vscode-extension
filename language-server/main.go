package main

import (
	"sync"

	"fmt"
	"os"
)



func main() {
	fmt.Println("Starting Easy-RPC language server")

	wg.Add(3)
	go readerLoop()
	go writerLoop()
	go handlerLoop()

	wg.Wait()
}
