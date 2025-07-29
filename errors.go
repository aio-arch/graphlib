package graphlib

import (
	"fmt"
	"strings"
)

type ErrUnknownNode[V comparable] struct {
	node V
}

func (e *ErrUnknownNode[V]) Error() string {
	return fmt.Sprintf("unknown node: %v", e.node)
}

type ErrCycle[V comparable] struct {
	nodes []V
}

func (e *ErrCycle[V]) Error() string {
	ss := make([]string, len(e.nodes))
	for i, n := range e.nodes {
		ss[i] = fmt.Sprintf("%v", n)
	}
	return fmt.Sprintf("graph contains a cycle: %s", strings.Join(ss, ", "))
}
