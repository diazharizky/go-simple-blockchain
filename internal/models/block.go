package models

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Data     []byte
	PrevHash []byte
	Hash     []byte
}

func NewBlock(data string, prevHash []byte) (newBlock *Block) {
	newBlock = &Block{
		Data:     []byte(data),
		PrevHash: prevHash,
	}
	newBlock.resolveHash()

	return
}

func (b *Block) resolveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}
