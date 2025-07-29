package graphlib

type nodeInfo[V comparable] struct {
	SortIdx         uint       //Sort idx
	PredecessorNums uint       //Number of parent nodes
	SuccessorNums   uint       //Record next Successor append index
	Successor       map[V]uint //Child Nodes
}

func (n *nodeInfo[V]) SuccessorSortSet() []V {
	set := make([]V, n.SuccessorNums)
	for node, idx := range n.Successor {
		set[idx] = node
	}
	return set
}

type Graph[V comparable] struct {
	nodeNums  uint               //Record next node append index
	node2info map[V]*nodeInfo[V] //Node parent/child info
}

func NewGraph[V comparable]() *Graph[V] {
	return &Graph[V]{
		node2info: make(map[V]*nodeInfo[V]),
	}
}

func (g *Graph[V]) AddNode(node V) *nodeInfo[V] {
	if result, has := g.node2info[node]; has {
		return result
	}
	result := nodeInfo[V]{
		SortIdx:   g.nodeNums,
		Successor: make(map[V]uint, 2),
	}
	g.node2info[node] = &result
	g.nodeNums++
	return &result
}

func (g *Graph[V]) AddEdge(from, to V) {
	var f, t *nodeInfo[V]
	var has bool
	if f, has = g.node2info[from]; !has {
		panic(ErrUnknownNode[V]{node: from})
	}
	if t, has = g.node2info[to]; !has {
		panic(ErrUnknownNode[V]{node: to})
	}
	if _, has := f.Successor[to]; !has {
		f.Successor[to] = f.SuccessorNums
		f.SuccessorNums++
		t.PredecessorNums++
	}
}

func (g *Graph[V]) NodeSortSet() []V {
	set := make([]V, g.nodeNums)
	for node, info := range g.node2info {
		set[info.SortIdx] = node
	}
	return set
}

func (g *Graph[V]) Add(node V, predecessors ...V) {
	g.AddNode(node)
	if len(predecessors) > 0 {
		for _, predecessor := range predecessors {
			g.AddNode(predecessor)
			g.AddEdge(predecessor, node)
		}
	}
}

// IsAcyclic checks if the graph is acyclic. If not, return the first detected cycle.
// it using https://github.com/python/cpython/blob/3.14/Lib/graphlib.py#L202 _find_cycle method
func (g *Graph[V]) IsAcyclic() ([]V, bool) {
	type iterItem[V comparable] struct {
		val   V
		isEnd bool
	}
	stack := make([]V, 0, len(g.node2info))
	itStack := make([]iterItem[V], 0, len(g.node2info))
	seen := make(map[V]struct{}, len(g.node2info))
	node2stacki := make(map[V]int, len(g.node2info))
	for _, node := range g.NodeSortSet() {
		if _, has := seen[node]; has {
			continue
		}

		for {
			if _, has := seen[node]; has {
				if _, has := node2stacki[node]; has {
					// cycle
					return append(stack[node2stacki[node]:], node), false
				}
			} else {
				seen[node] = struct{}{}
				itStack = append(itStack, iterItem[V]{isEnd: true})
				for _, successor := range g.node2info[node].SuccessorSortSet() {
					itStack = append(itStack, iterItem[V]{val: successor})
				}
				node2stacki[node] = len(stack)
				stack = append(stack, node)
			}

			for len(stack) > 0 {
				//itStack pop
				iter := itStack[len(itStack)-1]
				itStack = itStack[:len(itStack)-1]
				if iter.isEnd {
					cleanNode := stack[len(stack)-1]
					//stack pop
					stack = stack[:len(stack)-1]
					//node2stacki pop
					delete(node2stacki, cleanNode)
				} else {
					node = iter.val
					break
				}
			}
			if len(stack) == 0 {
				break
			}
		}
	}
	return nil, true
}
