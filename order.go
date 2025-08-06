package graphlib

type topologicalOrder[V comparable] struct {
	graph         *Graph[V]
	traversedNums uint       // next traversed node idx
	traversed     map[V]uint // traversed node idx
	notReadyNodes map[V]int  // node with predecessor node's nums
}

func (ts *topologicalOrder[V]) done(nodes *[]V) []V {
	readyNodes := make([]V, 0, 8)
	for _, node := range *nodes {
		ts.traversed[node] = ts.traversedNums
		ts.traversedNums++

		nodeInfo := ts.graph.node2info[node]
		for _, successor := range nodeInfo.Successors {
			successorInfo := ts.graph.node2info[successor]
			predecessorNums := len(successorInfo.Predecessors)
			if predecessorNums == 1 || ts.notReadyNodes[successor] == 1 {
				readyNodes = append(readyNodes, successor)
				delete(ts.notReadyNodes, successor)
			} else {
				ts.notReadyNodes[successor] = predecessorNums - 1
			}
		}
	}
	return readyNodes
}

func (ts *topologicalOrder[V]) traverse() {
	//find roots
	rootNodes := make([]V, 0, len(ts.graph.node2info)/2)
	for _, v := range ts.graph.NodeSortSet() {
		nodeInfo := ts.graph.node2info[v]
		if len(nodeInfo.Predecessors) == 0 {
			rootNodes = append(rootNodes, v)
		}
	}
	reayNodes := ts.done(&rootNodes)
	for len(reayNodes) > 0 {
		reayNodes = ts.done(&reayNodes)
	}
}

func graphOrder[V comparable](g *Graph[V]) []V {
	nodeLen := len(g.node2info)

	ts := topologicalOrder[V]{
		graph:         g,
		traversed:     make(map[V]uint, nodeLen),
		notReadyNodes: make(map[V]int, nodeLen),
	}
	ts.traverse()

	set := make([]V, ts.traversedNums)
	for node, idx := range ts.traversed {
		set[idx] = node
	}
	return set
}
