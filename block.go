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

func main() {
	block := NewBlock("123abc", []byte{12})
	fmt.Println(block.Hash)
}