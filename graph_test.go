package graphlib

import (
	"strconv"
	"testing"

	"golang.org/x/exp/slices"
)

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
			_ = cycle
			if ok != tt.want {
				t.Errorf("expect[%d] %v, got %v", i, tt.want, ok)
			}
			if !slices.Equal(cycle, tt.cycle) {
				t.Errorf("expect[%d] %v, got %v", i, tt.cycle, cycle)
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
			_ = cycle
			if ok != tt.want {
				t.Errorf("expect[%d] %v, got %v", i, tt.want, ok)
			}
			if !slices.Equal(cycle, tt.cycle) {
				t.Errorf("expect[%d] %v, got %v", i, tt.cycle, cycle)
			}
		})
	}
}
