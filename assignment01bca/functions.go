package assignment01bca

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	CurrentHash  string
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.CurrentHash = CalculateHash(fmt.Sprintf("%v%d%s", block.Transaction, block.Nonce, block.PreviousHash))
	return block
}

var blockchain []*Block

func ListBlocks() {
	for _, block := range blockchain {
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Current Hash: %s\n", block.CurrentHash)
		fmt.Println("----------------------------------------------------")
	}
}

func ChangeBlock(block *Block, newTransaction string) {
	block.Transaction = newTransaction
	block.CurrentHash = CalculateHash(fmt.Sprintf("%v%d%s", block.Transaction, block.Nonce, block.PreviousHash))
}

func VerifyChain() bool {
	for i := 1; i < len(blockchain); i++ {
		if blockchain[i].PreviousHash != blockchain[i-1].CurrentHash {
			return false
		}
	}
	return true
}

func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	return fmt.Sprintf("%x", hash)
}

func AddBlock(block *Block) {
	blockchain = append(blockchain, block)
}
