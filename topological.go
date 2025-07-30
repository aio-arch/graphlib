package graphlib

func TopologicalOrder[V comparable](digraph *Graph[V]) ([]V, error) {
	cycleNodes, isAcyclic := digraph.IsAcyclic()
	if !isAcyclic {
		return nil, &ErrCycle[V]{cycleNodes}
	}

	set := graphOrder(digraph)
	return set, nil
}

func TopologicalPrune[V comparable](digraph *Graph[V], nodes []V) (*Graph[V], error) {
	//check node is in graph
	for _, node := range nodes {
		if _, has := digraph.node2info[node]; !has {
			return nil, &ErrUnknownNode[V]{node}
		}
	}
	cycleNodes, isAcyclic := digraph.IsAcyclic()
	if !isAcyclic {
		return nil, &ErrCycle[V]{cycleNodes}
	}

	g := graphPrune(digraph, nodes)
	return g, nil
}
