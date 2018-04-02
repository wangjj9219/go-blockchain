package main

// BlockChain是一个Block指针数组
type BlockChain struct {
	blocks []*Block
}

// NewBlockChain 传建一个有创世块的链
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

// AddBlock向链中加入一个新块
// data在实际中就是交易
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
