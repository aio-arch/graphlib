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
	waitNodes := make(map[V]struct{}, len(nodes))
	traversed := make(map[V]struct{}, len(nodes))
	for _, node := range nodes {
		waitNodes[node] = struct{}{}
	}
	for len(waitNodes) > 0 {
		for node := range waitNodes {
			delete(waitNodes, node)
			if _, has := traversed[node]; has {
				continue
			}
			traversed[node] = struct{}{}
			nodeInfo := tp.graph.node2info[node]
			if len(nodeInfo.Predecessors) == 0 {
				tp.target.AddNode(node)
				continue
			}
			for _, predecessor := range nodeInfo.Predecessors {
				tp.target.Add(node, predecessor)
				waitNodes[predecessor] = struct{}{}
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
