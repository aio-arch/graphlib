# graphlib
[![Go Report Card](https://goreportcard.com/badge/github.com/aio-arch/graphlib)](https://goreportcard.com/report/github.com/aio-arch/graphlib)
[![Codecov](https://img.shields.io/codecov/c/github/aio-arch/graphlib?style=flat-square&logo=codecov)](https://app.codecov.io/gh/aio-arch/graphlib)
[![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/aio-arch/graphlib/go.yml)](https://github.com/aio-arch/graphlib/actions)
![Minimum Go Version](https://img.shields.io/badge/go-%3E%3D1.18-30dff3?style=flat-square&logo=go)

A Topological sort lib.

Sorting and pruning of DAG graphs.

Ideas borrowed from [python graphlib](https://github.com/python/cpython/blob/3.14/Lib/graphlib.py)

# How to install
```bash
go get -u github.com/aio-arch/graphlib
```

# How to use
```golang
    import "github.com/aio-arch/graphlib"

    // New A graph

	// for string type
	g1 := graphlib.NewGraph[string]()

	// for int type
	g2 := graphlib.NewGraph[int]()

	// for add node and add edge
	g1.AddNode("A1")
	g1.AddNode("B2")
	g1.AddEdge("A1", "B2") // edge: A1 -> B2

	// for add mulit edge,add node inline
	g2.Add(10, 1, 9)    // edge: 1 -> 10 and 9 -> 10
	g2.Add(100, 10, 90) // edge: 10 -> 100 and 90 -> 100

	// topological order
	topo, err := graphlib.TopologicalOrder(g1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Topological Order:%v\n", topo)

	// topological prune
	g3, err := graphlib.TopologicalPrune(g2, []int{10, 90})
	if err != nil {
		fmt.Println(err.Error())
	}
	_ = g3

```
