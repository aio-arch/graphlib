package graphlib

import (
	"errors"
	"strconv"
	"testing"
)

// golang.org/x/exp/slices.Equal
// support golang version < 1.21
func slicesEqual[S ~[]E, E comparable](s1, s2 S) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

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
			graph: [][]string{
				{"A-1-1", "A-1"},
				{"A-1-2", "A-1"},
				{"A-1", "A", "C-1"},
				{"A-2-1", "A-2"},
				{"A-2-2", "A-2"},
				{"A-2", "A"},
				{"B-1-1", "B-1"},
				{"B-1", "B"},
				{"A"},
				{"B"},
				{"C-1", "A-2-2", "B-1-1"},
			},
			want: []string{"A", "B", "A-2", "B-1", "A-2-1", "A-2-2", "B-1-1", "C-1", "A-1", "A-1-1", "A-1-2"},
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
			if !slicesEqual(topo, tt.want) {
				t.Errorf("expect want=%v, got=%v", tt.want, topo)
			}
		})
	}
}

func TestTopologicalPrune(t *testing.T) {
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
	graph := [][]string{
		{"A-1-1", "A-1"},
		{"A-1-2", "A-1"},
		{"A-1", "A", "C-1"},
		{"A-2-1", "A-2"},
		{"A-2-2", "A-2"},
		{"A-2", "A"},
		{"B-1-1", "B-1"},
		{"B-1", "B"},
		{"A"},
		{"B"},
		{"C-1", "A-2-2", "B-1-1"},
	}
	g := NewGraph[string]()
	for _, vals := range graph {
		g.Add(vals[0], vals[1:]...)
	}

	cases := map[string]struct {
		nodes   []string
		want    []string
		wantErr error
	}{
		"prune unknown node": {
			nodes: []string{
				"unknown",
			},
			want:    []string{},
			wantErr: &ErrUnknownNode[string]{node: "unknown"},
		},
		"prune A,B": {
			nodes: []string{
				"B",
				"A",
			},
			want: []string{"A", "B"},
		},
		"prune A,B-1": {
			nodes: []string{
				"B-1",
				"A",
			},
			want: []string{"A", "B", "B-1"},
		},
		"prune A-2,C-1": {
			nodes: []string{
				"A-2",
				"C-1",
			},
			want: []string{"A", "B", "A-2", "B-1", "A-2-2", "B-1-1", "C-1"},
		},
		"prune A-1-1": {
			nodes: []string{
				"A-1-1",
			},
			want: []string{"A", "B", "A-2", "B-1", "A-2-2", "B-1-1", "C-1", "A-1", "A-1-1"},
		},
	}

	for i, tt := range cases {
		t.Run(i, func(t *testing.T) {
			target, err := TopologicalPrune(g, tt.nodes)
			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("expected error %v, got %v", tt.want, err)
				}
			} else {
				topo, _ := TopologicalOrder(target)
				if !slicesEqual(topo, tt.want) {
					t.Errorf("expect want=%v, got=%v", tt.want, topo)
				}
			}
		})
	}
}
