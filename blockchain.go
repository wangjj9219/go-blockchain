package main

// BlockChain是一个Block指针数组
type BlockChain struct {
	blocks []*Block
}

// NewBlockChain 创建一个有创世块的链
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

// AddBlock 向链中添加入一个块
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
