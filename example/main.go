package main

import (
	"fmt"

	"github.com/aio-arch/graphlib"
)

func main() {

	/*
	                                ┌────► A-1-1
	                                │
	   ┌────────────────────────► A-1
	   │                           ▲│
	   A       ┌─────► A-2-1       │└────► A-1-2
	   │       │                   │
	   └────► A-2                  │
	           │                   │
	           └─────► A-2-2       │
	                      │        │
	                      ├─────► C-1
	                      │
	   B ───► B-1 ───► B-1-1
	*/
	g := graphlib.NewGraph[string]()

	// add node
	g.AddNode("A-1-1")
	g.AddNode("A-1")

	//add edge
	g.AddEdge("A-1", "A-1-1") // edge: A-1 -> A-1-1

	g.AddNode("A-1-2")
	g.AddEdge("A-1", "A-1-2") // edge: A-1 -> A-1-2

	// add multiple edges and add node inline
	g.Add("A-1", "A", "C-1") // edge: A -> A-1 and C-1 -> A-1
	g.Add("A-2-1", "A-2")    // edge: A-2 -> A-2-1
	g.Add("A-2-2", "A-2")
	g.Add("A-2", "A")
	g.Add("B-1-1", "B-1")
	g.Add("B-1", "B")
	g.Add("A") // add node A
	g.Add("B") // add node B
	g.Add("C-1", "A-2-2", "B-1-1")

	topo, err := graphlib.TopologicalOrder(g)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Topological Order:%v\n", topo)
	//Topological Order:[A B A-2 B-1 A-2-1 A-2-2 B-1-1 C-1 A-1 A-1-1 A-1-2]

	g2, err := graphlib.TopologicalPrune(g, []string{"A-1-1"})
	if err != nil {
		panic(err)
	}
	topo2, err := graphlib.TopologicalOrder(g2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Pruned Topological Order:%v\n", topo2)
	//Pruned Topological Order:[A B A-2 B-1 A-2-2 B-1-1 C-1 A-1 A-1-1]
}
