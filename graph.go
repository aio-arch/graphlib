package graphlib

import "fmt"

type Graph[V comparable] struct {
	sorter    []V
	node2info map[V]*nodeInfo[V]
}

func NewGraph[V comparable]() *Graph[V] {
	return &Graph[V]{
		sorter:    make([]V, 0, 8),
		node2info: make(map[V]*nodeInfo[V]),
	}
}

func (g *Graph[V]) AddNode(node V) *nodeInfo[V] {
	if result, has := g.node2info[node]; has {
		return result
	}
	result := nodeInfo[V]{
		Node:      node,
		Successor: make([]V, 0, 2),
	}
	g.node2info[node] = &result
	g.sorter = append(g.sorter, node)
	return &result
}

func (g *Graph[V]) AddEdge(from, to V) {
	var f, t *nodeInfo[V]
	var has bool
	if f, has = g.node2info[from]; !has {
		panic(fmt.Sprintf("Add Edge err,from node[%v] is not exist", from))
	}
	if t, has = g.node2info[to]; !has {
		panic(fmt.Sprintf("Add Edge err,to node[%v] is not exist", to))
	}
	f.Successor = append(f.Successor, to)
	t.PredecessorNums += 1
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
	for _, node := range g.sorter {
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
				for _, successor := range g.node2info[node].Successor {
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

type nodeInfo[V comparable] struct {
	Node            V
	Successor       []V
	PredecessorNums int
}
