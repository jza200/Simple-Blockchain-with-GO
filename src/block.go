package blockchain

import(
	"fmt"
	"encoding/hex"
	"crypto/sha256"
	"math"
)

type Block struct {
	PrevHash   []byte
	Generation uint64
	Difficulty uint8
	Data       string
	Proof      uint64
	Hash       []byte
}

// Create new initial (generation 0) block, not setting the .Hash value.
func Initial(difficulty uint8) Block {
	var initialBlock Block
	var temp [32]byte
	temp = [32]byte {}
	x := temp[:]
	//slicing method adopted from stack overflow
	initialBlock.Difficulty = difficulty
	initialBlock.Generation = 0
	initialBlock.Data = ""
	initialBlock.PrevHash = x
	return initialBlock
}


// Create new block to follow this block, with provided data, not setting the .Hash value.
func (prev_block Block) Next(data string) Block {
	var newBlock Block
	newBlock.Difficulty = prev_block.Difficulty
	newBlock.Generation = prev_block.Generation + 1
	newBlock.Data = data
	newBlock.PrevHash = prev_block.Hash
	return newBlock
}

// String that we hash for this block.
func (blk Block) hashString() string {
	return blk.hashStringProof(blk.Proof)
}

// String that we hash for this block, if we had blk.Proof == proof.
func (blk Block) hashStringProof(proof uint64) string {
	var previousHash string
	var theGneration uint64
	var theDifficulty uint8
	var theDataString string
	var theProofOfWork uint64

	previousHash = hex.EncodeToString(blk.PrevHash)
	theGneration = blk.Generation 
	theDifficulty = blk.Difficulty
	theDataString = blk.Data
	theProofOfWork = blk.Proof

	var myString string
	myString = fmt.Sprintf("%v:%v:%v:%v:%v",previousHash, theGneration, theDifficulty, theDataString, theProofOfWork)
	//fmt.Println(myString)
	return myString
}

// Calculate hash as if blk.Proof == proof.
// Separated from .CalcHash so we can test many proof values without
// modifying the Block.
func (blk Block) calcHashProof(proof uint64) []byte {
	var myString string
	myString = blk.hashStringProof(proof)
	//fmt.Println(myString)

	myByte := []byte(myString)
	slice := myByte[:]
	//slicing method adopted from stack overflow

	myHash := sha256.Sum256(slice)
	return myHash[:]
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	return blk.calcHashProof(blk.Proof)
}

// Would this hash end in enough null bits, if blk.Proof == proof?
func (blk Block) validHashProof(proof uint64) bool {
	var indice int
	theDifficulty := blk.Difficulty
	nBytes := theDifficulty / 8
	nBits := theDifficulty % 8
	indice = len(blk.Hash) - int(nBytes) - 1

	for i := 0; i < int(nBytes); i++ {
		if blk.Hash[len(blk.Hash) - i - 1] != '\x00' {
			return false
		}
	}

	temp := 1<<nBits
	bitDiv := float64(temp)
	temp1 := blk.Hash[indice]
	hashdiv := float64(temp1)
	//fmt.Println("%T", hashdiv)
	//changed all value from uint to float and magically worked.....
	if math.Mod(hashdiv, bitDiv) != 0 {
		return false
	}
	// divide by 2 ^ n

	return true
}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	return blk.validHashProof(blk.Proof)
}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
