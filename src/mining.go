package blockchain

import (
	"work_queue"
	"fmt"
	"math"
)

// Mine in a very simple way: check sequentically until a valid hash is found.
// This doesn't *need* to be used in any way, but could be used to do some mining
// before your .Mine is complete. Results should be the same as .Mine (but slower).
func (blk *Block) mineSequential() {
	proof := uint64(0)
	for !blk.validHashProof(proof) {
		proof += 1
	}
	blk.SetProof(proof)
}

// carries range to mine and basic info needed for mining
type miningWorker struct {
	start uint64
	end   uint64
	blk   Block
}

//check proof values in the range and return a MiningResult with possible valid proof
func (task miningWorker) Run() interface{} {
	var result MiningResult
	result.Found = false
	for i := task.start; i < task.end; i++ {
		task.blk.SetProof(i)
		//fmt.Println(i)
		//fmt.Println(task.blk.Proof)
		if (task.blk.ValidHash()) {
			//fmt.Println("validHash passes")
			//fmt.Println(task.blk)
			result.Found = true
			result.Proof = task.blk.Proof
			return result
		}
	}	
	return result
}

type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // was a valid proof-of-work found?
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.

// Call MineRange with start and end to split it into length/chunks and .Run() them concurrently, send them to work_queue
// MineRange and Mine and mine Sequencial does the same thing, excpet for concurrency.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	var noResult, myResult MiningResult
	queue := work_queue.Create(uint(workers), uint(chunks))
	length := uint64(math.Ceil(float64((end - start) / chunks) + 0.0001))

	for i := start; i < end; i+= length {
		mineWorker := miningWorker{i, i + length, blk}
		queue.Enqueue(mineWorker)
		//fmt.Println(i, i+length)
	}

	//fmt.Printf("type of queue %T, value of queue %v", queue, queue)
	//fmt.Println(len(queue.Jobs))
	//fmt.Println(len(queue.Results))


	for i := range queue.Results {
		temp := i // channel of interface{} type
		//MiningResult(temp) would not work, but type assert would
		myResult = temp.(MiningResult)
		if (myResult.Found == true) {
			queue.Shutdown()
			//fmt.Println("got something")
			return myResult
		}
	}
	queue.Shutdown()
	fmt.Println("freed")
	return noResult
}

// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << blk.Difficulty) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4567)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}
