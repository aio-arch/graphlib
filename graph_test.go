package graphlib

import (
	"strconv"
	"testing"

	"golang.org/x/exp/slices"
)

func TestAddEdge(t *testing.T) {
	cases := []struct {
		node []string
		edge []string
		want error
	}{
		{
			node: []string{"A", "B"},
			edge: []string{"A", "B"},
			want: nil,
		},
		{
			node: []string{"A", "B"},
			edge: []string{"A", "C"},
			want: &ErrUnknownNode[string]{node: "C"},
		},
		{
			node: []string{"A", "B"},
			edge: []string{"C", "B"},
			want: &ErrUnknownNode[string]{node: "C"},
		},
	}
	for i, tt := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				r := recover()
				if tt.want != nil || r != nil {
					err, ok := r.(ErrUnknownNode[string])
					if !ok {
						t.Errorf("expected error, got %v,not a error", r)
					}
					if err.Error() != tt.want.Error() {
						t.Errorf("expected error %v, got %v", tt.want, err)
					}
				}
			}()
			g := NewGraph[string]()
			for _, node := range tt.node {
				g.AddNode(node)
			}

			g.AddEdge(tt.edge[0], tt.edge[1])
		})
	}
}

func TestIsAcyclic_string(t *testing.T) {
	cases := []struct {
		graph [][]string
		want  bool
		cycle []string
	}{
		{
			graph: [][]string{
				{"A"},
				{"B", "A"},
				{"C", "B", "A"},
			},
			want:  true,
			cycle: []string{},
		},
		{
			graph: [][]string{
				{"A", "C"},
				{"B", "A"},
				{"C", "B", "A"},
			},
			want:  false,
			cycle: []string{"A", "C", "A"},
		},
	}

	for i, tt := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := NewGraph[string]()
			for _, vals := range tt.graph {
				g.Add(vals[0], vals[1:]...)
			}
			cycle, ok := g.IsAcyclic()
			if ok != tt.want {
				t.Errorf("expect want=%v, got=%v", tt.want, ok)
			}
			if !slices.Equal(cycle, tt.cycle) {
				t.Errorf("expect want=%v, got=%v", tt.cycle, cycle)
			}
		})
	}
}

func TestIsAcyclic_int(t *testing.T) {
	cases := []struct {
		graph [][]int
		want  bool
		cycle []int
	}{
		{
			graph: [][]int{
				{1},
				{2, 1},
				{3, 2, 1},
			},
			want:  true,
			cycle: []int{},
		},
		{
			graph: [][]int{
				{1, 3},
				{2, 1},
				{3, 2, 1},
			},
			want:  false,
			cycle: []int{1, 3, 1},
		},
	}

	for i, tt := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := NewGraph[int]()
			for _, vals := range tt.graph {
				g.Add(vals[0], vals[1:]...)
			}
			cycle, ok := g.IsAcyclic()
			if ok != tt.want {
				t.Errorf("expect want=%v, got=%v", tt.want, ok)
			}
			if !slices.Equal(cycle, tt.cycle) {
				t.Errorf("expect want=%v, got=%v", tt.cycle, cycle)
			}
		})
	}
}
