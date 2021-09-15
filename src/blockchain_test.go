package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)



func TestInitial(t *testing.T) {
	b0 := Initial(7)
	assert.Equal(t, 7, b0.Difficulty, "correct")
	assert.Equal(t, 0, b0.Generation, "correct")
	assert.Equal(t, "", b0.Data, "correct")
}

func TestMine(t *testing.T) {
	b0 := blockchain.Initial(7)
	b0.Mine(1)
	assert.Equal(t, 385, b0.Proof, "correct")
	assert.Equal(t, "379bf2fb1a558872f09442a45e300e72f00f03f2c6f4dd29971f67ea4f3d5300", hex.EncodeToString(b0.Hash), "correct")
}

func TestMine2(t *testing.T) {
	b0 := blockchain.Initial(7)
	b0.Mine(1)
	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	assert.Equal(t, 20, b1.Proof, "correct")
	assert.Equal(t, "4a1c722d8021346fa2f440d7f0bbaa585e632f68fd20fed812fc944613b92500", hex.EncodeToString(b1.Hash), "correct")
}

func TestMine3(t *testing.T) {
	b0 := blockchain.Initial(7)
	b0.Mine(1)
	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	assert.Equal(t, 40, b2.Proof, "correct")
	assert.Equal(t, "ba2f9bf0f9ec629db726f1a5fe7312eb76270459e3f5bfdc4e213df9e47cd380", hex.EncodeToString(b2.Hash), "correct")
}
// TODO: some useful tests of Blocks
