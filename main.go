package main

import (
	"fmt"
)

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Alice send 1 BTC to Bob")
	bc.AddBlock("Bob send 2 BTC to Casey")

	for _, block := range bc.blocks {
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
