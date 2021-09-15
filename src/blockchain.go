package blockchain

import (
	"bytes"
)

type Blockchain struct {
	Chain []Block
}

func (chain *Blockchain) Add(blk Block) {
	// You can remove the panic() here if you wish.
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}

	chain.Chain = append(chain.Chain, blk)
	// TODO
	// func append(s []T, value T) []T
}

// Check a few aspects of the validity of the Blockchain
func (chain Blockchain) IsValid() bool {
	for i, blk := range chain.Chain {
		if i != 0 {
			prevBlk := chain.Chain[i-1]
			if !bytes.Equal(prevBlk.Hash, blk.PrevHash) {
				return false // prev_hash doesn't match
			}
		}
		if !bytes.Equal(blk.Hash, blk.CalcHash()) {
			return false // hash doesn't match contents
		}
		if !blk.ValidHash() {
			return false // not enough trailing nulls
		}
	}
	return true
}
