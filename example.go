package main

import (
	"fmt"
)

type Block struct {
	id   int
	data string
}

var BlockChain []Block

func newBlock(data string) {
	var newID int

	if len(BlockChain) == 0 {
		newID = 0
	} else {
		newID = len(BlockChain) + 1
	}

	var newBlock Block

	newBlock.id = newID
	newBlock.data = data
	BlockChain = append(BlockChain, newBlock)

}

func displayAllBlocks() {
	var i int
	fmt.Println("Printing all blocks...")
	for i = 0; i < len(BlockChain); i++ {
		fmt.Println(BlockChain[i].id, ", ", BlockChain[i].data)
	}

}

func ModifyBlock(num int, newData string) {
	if num >= 0 && num < len(BlockChain) {
		BlockChain[num].data = newData
		fmt.Printf("Block modified - Block ID: %d, New Data: %s\n\n", num, newData)
	}
}
