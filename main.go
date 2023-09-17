package main

import (
	"assignment01bca/assignment01bca"
	"fmt"
)

func main() {

	block1 := assignment01bca.NewBlock("Alice to Bob", 123, "")
	assignment01bca.AddBlock(block1)

	block2 := assignment01bca.NewBlock("Bob to Charlie", 456, block1.CurrentHash)
	assignment01bca.AddBlock(block2)

	block3 := assignment01bca.NewBlock("Charlie to Alice", 789, block2.CurrentHash)
	assignment01bca.AddBlock(block3)

	assignment01bca.ListBlocks()

	assignment01bca.ChangeBlock(block2, "New transaction")

	isValid := assignment01bca.VerifyChain()
	if isValid {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is not valid.")
	}
}
