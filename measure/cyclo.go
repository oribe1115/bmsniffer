package measure

import (
	"fmt"

	"golang.org/x/tools/go/ssa"
)

type blockQueue struct {
	q []*ssa.BasicBlock
}

var indexExist = struct{}{}

// CyclomaticComplexity 対象の関数のCYCLOを計算する
func CyclomaticComplexity(ssaFunc *ssa.Function) int {
	if ssaFunc == nil {
		return -1
	}
	// v(G) = e-n+p
	n, e := countFlowGraphValues(ssaFunc)
	p := 2 // 連結されたコンポーネントの数 決め打ちでこれでいいのか確認
	return e - n + p
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

		for _, instr := range tmpNode.Instrs {
			if mClosure, _ := instr.(*ssa.MakeClosure); mClosure != nil {
				fmt.Println(mClosure.String(), ": ", mClosure.Fn.Name())
			}
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
