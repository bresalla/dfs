package graph

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGraph_AddEdges(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		edges    []*Edge
		expected []Edge
	}{
		{
			name: "Simple path A -> B -> C",
			edges: []*Edge{
				{From: "A", To: "B"},
				{From: "B", To: "C"},
			},
			expected: []Edge{{To: "A"}, {From: "A", To: "B"}, {From: "B", To: "C"}},
		},
		{
			name: "Two sided",
			edges: []*Edge{
				{From: "A", To: "B"},
				{From: "B", To: "A"},
			},
			expected: []Edge{{To: "A"}, {From: "A", To: "B"}, {From: "B", To: "A"}},
		},
		{
			name: "Disconnected graph",
			edges: []*Edge{
				{From: "A", To: "B"},
				{From: "C", To: "D"},
			},
			expected: []Edge{{To: "A"}, {From: "A", To: "B"}, {To: "C"}, {From: "C", To: "D"}}, // Assuming dfs only visits connected nodes
		},
		{
			name: "Long graph without cycles",
			edges: []*Edge{
				{From: "A", To: "B"},
				{From: "B", To: "C"},
				{From: "C", To: "D"},
				{From: "D", To: "E"},
				{From: "B", To: "D"},
				{From: "C", To: "E"},
			},
			expected: []Edge{{To: "A"}, {From: "A", To: "B"}, {From: "B", To: "C"}, {From: "C", To: "D"}, {From: "D", To: "E"}, {From: "C", To: "E"}, {From: "B", To: "D"}}, // DFS should visit all nodes
		},
		{
			name: "Long graph without cycles(inverted order)",
			edges: []*Edge{
				{From: "A", To: "D"},
				{From: "B", To: "C"},
				{From: "A", To: "B"},
			},
			expected: []Edge{{To: "A"}, {From: "A", To: "D"}, {From: "A", To: "B"}, {From: "B", To: "C"}}, // DFS should visit all nodes
		},
		{
			name: "Long graph with cycles",
			edges: []*Edge{
				{From: "A", To: "B"},
				{From: "B", To: "C"},
				{From: "C", To: "A"},
			},
			expected: []Edge{{To: "A"}, {From: "A", To: "B"}, {From: "B", To: "C"}, {From: "C", To: "A"}},
		},
		{
			name: "Long graph with cycles(inverted order)",
			edges: []*Edge{
				{From: "A", To: "C"},
				{From: "B", To: "A"},
				{From: "C", To: "B"},
			},
			expected: []Edge{{To: "A"}, {From: "A", To: "C"}, {From: "C", To: "B"}, {From: "B", To: "A"}},
		},
		{
			name: "Multiple graphs",
			edges: []*Edge{
				{From: "A", To: "B"},
				{From: "B", To: "C"},
				{From: "D", To: "E"},
				{From: "E", To: "F"},
			},
			expected: []Edge{{To: "A"}, {From: "A", To: "B"}, {From: "B", To: "C"}, {To: "D"}, {From: "D", To: "E"}, {From: "E", To: "F"}},
		},
		{
			name: "Multiple graphs with cycles",
			edges: []*Edge{
				{From: "A", To: "B"},
				{From: "B", To: "C"},
				{From: "B", To: "D"},
				{From: "B", To: "E"},
				{From: "C", To: "A"},
				{From: "D", To: "A"},
				{From: "E", To: "A"},
			},
			expected: []Edge{{To: "A"}, {From: "A", To: "B"}, {From: "B", To: "C"}, {From: "C", To: "A"}, {From: "B", To: "D"}, {From: "D", To: "A"}, {From: "B", To: "E"}, {From: "E", To: "A"}},
		}, {
			name: "Two input one output",
			edges: []*Edge{
				{From: "A", To: "B"},
				{From: "C", To: "B"},
			},
			expected: []Edge{{To: "A"}, {From: "A", To: "B"}, {To: "C"}, {From: "C", To: "B"}},
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			graph := &Graph{}
			graph.AddEdges(tc.edges)
			res := graph.DFS()
			if diff := cmp.Diff(tc.expected, res); diff != "" {
				t.Errorf("dfs() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGraph_selectStartNode(t *testing.T) {
	type fields struct {
		Edges []*Edge
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Simple start node",
			fields: fields{
				Edges: []*Edge{
					{From: "A", To: "B"},
					{From: "B", To: "C"},
				},
			},
			want: "A",
		},
		{
			name: "Start node with disconnected graph",
			fields: fields{
				Edges: []*Edge{
					{From: "A", To: "B"},
					{From: "C", To: "D"},
				},
			},
			want: "A",
		},
		{
			name: "Cyclic graph",
			fields: fields{
				Edges: []*Edge{
					{From: "A", To: "B"},
					{From: "B", To: "C"},
					{From: "C", To: "A"},
				},
			},
			want: "A",
		},
		{
			name: "Single edge graph",
			fields: fields{
				Edges: []*Edge{
					{From: "A", To: "A"},
				},
			},
			want: "A",
		},
		{
			name: "Multiple edges with shared nodes",
			fields: fields{
				Edges: []*Edge{
					{From: "A", To: "B"},
					{From: "B", To: "C"},
					{From: "C", To: "D"},
					{From: "D", To: "E"},
				},
			},
			want: "A",
		}, {
			name: "Multiple edges with shared nodes (inverted order)",
			fields: fields{
				Edges: []*Edge{
					{From: "C", To: "D"},
					{From: "B", To: "C"},
					{From: "A", To: "B"},
				},
			},
			want: "A",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Edges: tt.fields.Edges,
			}
			if got := g.selectStartNode(nil); got != tt.want {
				t.Errorf("selectStartNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
