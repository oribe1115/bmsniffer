package measure

import (
	"golang.org/x/tools/go/ssa"
)

type blockQueue struct {
	q []*ssa.BasicBlock
}

var indexExist = struct{}{}

func CyclomaticComplexity(ssaFunc *ssa.Function) {
	// for _, ssaBlock := range ssaFunc.Blocks {
	// 	fmt.Println(ssaBlock.Index, ssaBlock.Instrs, ssaBlock.Succs)
	// }
}

func countFlowGraphValues(ssaFunc *ssa.Function) (nodeCount int, edgeCount int) {
	checkedNodeIndexMap := map[int]struct{}{}
	nodeQueue := initIntQueue()
	nodeQueue.enqueue(ssaFunc.Blocks[0])

	for nodeQueue.len() != 0 {
		tmpNode := nodeQueue.dequeue()
		if _, exist := checkedNodeIndexMap[tmpNode.Index]; exist {
			continue
		}

		nodeCount++
		checkedNodeIndexMap[tmpNode.Index] = indexExist
		for _, succ := range tmpNode.Succs {
			edgeCount++
			nodeQueue.enqueue(succ)
		}
	}

	return nodeCount, edgeCount
}

func initIntQueue() *blockQueue {
	queue := &blockQueue{}
	queue.q = make([]*ssa.BasicBlock, 0)
	return queue
}

func (queue *blockQueue) enqueue(block *ssa.BasicBlock) {
	queue.q = append(queue.q, block)
}

func (queue *blockQueue) dequeue() *ssa.BasicBlock {
	result := queue.q[0]
	queue.q = queue.q[1:]
	return result
}

func (queue *blockQueue) len() int {
	return len(queue.q)
}
