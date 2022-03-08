package main

import (
	"bufio"
	// "erpcLanguageServer/jsonrpc"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting Easy-RPC language server")
	// msg, err := jsonrpc.ReadNextMessage(os.Stdin)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	r := bufio.NewReader(os.Stdin)
	b := make([]byte, 5000)
	r.Read(b)
	fmt.Printf("%q", b)

	os.Exit(0)
}
