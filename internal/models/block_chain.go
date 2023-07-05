package models

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain(data string) *BlockChain {
	genesis := &Block{
		Data: []byte(data),
	}

	return &BlockChain{
		Blocks: []*Block{genesis},
	}
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1] // Get the last block in chain
	newBlock := NewBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}
