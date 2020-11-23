package measure

import (
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

type SSAData struct {
	FuncMap map[string]*ssa.Function
}

type blockQueue struct {
	q []*ssa.BasicBlock
}

var indexExist = struct{}{}

// CyclomaticComplexity 対象の関数のCYCLOを計算する
func CyclomaticComplexity(funcName string, ssaData *SSAData) int {
	ssaFunc := ssaData.getSSAFunc(funcName)
	if ssaFunc == nil {
		return -1
	}
	// v(G) = e-n+p
	n, e := ssaData.countFlowGraphValues(ssaFunc)
	p := 2
	return e - n + p
}

func GetSSAData(ssaInfo *buildssa.SSA) *SSAData {
	ssaFuncMap := map[string]*ssa.Function{}

	for _, ssaFunc := range ssaInfo.SrcFuncs {
		ssaFuncMap[ssaFunc.Name()] = ssaFunc
	}

	return &SSAData{FuncMap: ssaFuncMap}
}

func (sd *SSAData) countFlowGraphValues(ssaFunc *ssa.Function) (nodeCount int, edgeCount int) {
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
				if cSSAFunc := sd.getSSAFunc(mClosure.Fn.Name()); cSSAFunc != nil {
					cNodeCount, cEdgeCount := sd.countFlowGraphValues(cSSAFunc)
					nodeCount += cNodeCount + 1
					edgeCount += cEdgeCount + 2
				}
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

func (sd *SSAData) getSSAFunc(funcName string) *ssa.Function {
	if funcName == "init" {
		funcName = "init#1"
	}

	return sd.FuncMap[funcName]
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
