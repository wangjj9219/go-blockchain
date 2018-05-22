package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	PrevBlockHash []byte
	Hash []byte
	Data []byte
}


// NewBlock用于生成新块，需要参数Data和PrevBlockHash
// 当前块的Hash会根据Data和PrevBlockHash的计算
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Data:          []byte(data),
	}
	block.SetHash()

	return block
}


// SetHash 设置当前块hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}


// NewGenesisBlock 生成创世块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}