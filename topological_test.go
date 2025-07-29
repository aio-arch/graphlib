package graphlib

import (
	"errors"
	"strconv"
	"testing"

	"golang.org/x/exp/slices"
)

func TestTopologicalOrder(t *testing.T) {
	cases := []struct {
		graph [][]string
		want  []string
	}{
		{
			graph: [][]string{
				{"A"},
				{"B", "A"},
				{"C", "B", "A"},
			},
			want: []string{"A", "B", "C"},
		},
		{
			graph: [][]string{
				{"A-1-1", "A-1"},
				{"A-1-2", "A-1"},
				{"A-1", "A"},
				{"A-2-1", "A-2"},
				{"A-2-2", "A-2"},
				{"A-2", "A"},
				{"B-1-1", "B-1"},
				{"B-1", "B"},
				{"A"},
				{"B"},
			},
			want: []string{"A", "B", "A-1", "A-2", "B-1", "A-1-1", "A-1-2", "A-2-1", "A-2-2", "B-1-1"},
		},
		{
			graph: [][]string{
				{"A", "C"},
				{"B", "A"},
				{"C", "B", "A"},
			},
			want: []string{},
		},
	}

	for i, tt := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := NewGraph[string]()
			for _, vals := range tt.graph {
				g.Add(vals[0], vals[1:]...)
			}
			topo, err := TopologicalOrder(g)
			var errCycle *ErrCycle[string]
			if err != nil && !errors.As(err, &errCycle) {
				t.Errorf("expect wantError=ErrCycle, got=%v", err.Error())
			}
			if !slices.Equal(topo, tt.want) {
				t.Errorf("expect want=%v, got=%v", tt.want, topo)
			}
		})
	}
}
