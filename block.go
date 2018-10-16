package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp    int64
	Data         []byte
	PreBlockHash []byte
	Hash         []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	test := [][]byte{b.PreBlockHash, b.Data, timestamp}
	//fmt.Println(test)
	headers := bytes.Join(test, []byte{})
	//fmt.Println(headers)
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

type BlockChain struct {
	blocks []*Block
}

func (blockChain *BlockChain) AddBlock(data string) {
	prevHash := blockChain.blocks[len(blockChain.blocks)-1].Hash
	newBlock := NewBlock(data, prevHash)
	blockChain.blocks = append(blockChain.blocks, newBlock)
}

//创世区块
func NewGenesisBlock() *Block {
	return NewBlock("GenesisBlock", []byte{})
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PreBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
