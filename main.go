package main

import (
	"fmt"

	"github.com/diazharizky/go-simple-blockchain/internal/models"
)

var bc *models.BlockChain

func init() {
	bc = models.NewBlockChain("String Blockchain!")
}

func main() {
	bc.AddBlock("First block")
	bc.AddBlock("Second block")
	bc.AddBlock("Third block")

	for _, b := range bc.Blocks {
		fmt.Printf("Prev hash: %v\n", string(b.PrevHash))
		fmt.Printf("Data: %v\n", string(b.Data))
		fmt.Printf("Hash: %v\n", string(b.Hash))
		fmt.Println()
	}
}
