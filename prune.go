package graphlib

type topologicalPrune[V comparable] struct {
	graph  *Graph[V]
	target *Graph[V]
}

// func (tp *topologicalPrune[V]) sortWaitingNodes(nodes []V) {
// 	sort.Slice(nodes, func(i, j int) bool {
// 		return tp.graph.node2info[nodes[i]].SortIdx < tp.graph.node2info[nodes[j]].SortIdx
// 	})
// }

func (tp *topologicalPrune[V]) prune(nodes []V) *Graph[V] {

	traversed := make(map[V]int, len(nodes))
	readyNodes := make(map[V]uint, len(nodes))
	childNodes := make(map[V]uint, len(nodes))
	for _, node := range nodes {
		nodeInfo := tp.graph.node2info[node]
		readyNodes[node] = nodeInfo.PredecessorNums
	}

	for len(readyNodes) > 0 {
		for node, predecessorNums := range readyNodes {

			delete(readyNodes, node)
			traversed[node] = 0
			if predecessorNums == 0 {
				tp.target.AddNode(node)
			} else {
				childNodes[node] = predecessorNums
			}
		}

		for node, info := range tp.graph.node2info {
			traversedSuccessorNums, ok := traversed[node]
			if ok && len(info.Successors) == traversedSuccessorNums {
				continue
			}

			var ready bool
			for _, successor := range info.Successors {
				if successorPredecessorNums, has := childNodes[successor]; has {
					ready = true
					tp.target.Add(successor, node)
					traversed[node] = traversedSuccessorNums + 1
					if successorPredecessorNums == 1 {
						delete(childNodes, successor)
					} else {
						childNodes[successor] = successorPredecessorNums - 1
					}
				}

			}
			if ready && info.PredecessorNums > 0 {
				readyNodes[node] = info.PredecessorNums
			}
			if len(childNodes) == 0 {
				break
			}
		}
	}

	// sort by graph order
	var idx uint
	for _, node := range tp.graph.NodeSortSet() {
		if info, has := tp.target.node2info[node]; has {
			info.SortIdx = idx
			idx++
		}
		if idx == tp.target.nodeNums {
			break
		}
	}

	return tp.target
}

func graphPrune[V comparable](g *Graph[V], nodes []V) *Graph[V] {
	tp := topologicalPrune[V]{
		graph:  g,
		target: NewGraph[V](),
	}
	tp.prune(nodes)
	return tp.target
}
