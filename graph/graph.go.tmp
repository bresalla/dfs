package graph

import (
	"fmt"
	"slices"
)

type Edge struct {
	From string
	To   string
}
type Graph struct {
	Edges        []*Edge
	indexEnabled bool
}

func (g *Graph) AddEdges(edges []*Edge) {
	g.Edges = append(g.Edges, edges...)
}

type Option func(*Graph)

func WithIndex() Option {
	return func(g *Graph) {
		g.indexEnabled = true
	}
}

func (g *Graph) DFS(options ...Option) []Edge {
	for _, o := range options {
		o(g)
	}
	var stack *Stack = NewStack()
	visited := make(map[string][]string)
	path := []Edge{}
	var from string
	for g.unprocessedEdgesAvailible(path) {
		stack = NewStack()
		from = ""
		current := g.selectStartNode(path)
		stack.Push(current)
		fmt.Println("Push:", current)
		for !stack.IsEmpty() {
			current, ok := stack.Pop()
			fmt.Println("Pop:", current)
			if !ok {
				break
			}
			path = g.checkAndAddToPath(path, current, from)
			fmt.Println("Path:", path)
			from = current
			// Assuming GetFirstUnvisitedChild returns the first unvisited child
			if child, ok := g.getNextUnvisitedChild(current, visited); ok {

				stack.Push(current) // Push current node back to stack
				fmt.Println("Push:", current)
				stack.Push(child) // Push child to stack
				fmt.Println("Push:", child)
				visited[current] = append(visited[current], child)
			}
		}
	}
	if g.indexEnabled {
		path = g.markLeaves(path)
	}

	return path
}

func (g *Graph) getNextUnvisitedChild(from string, visited map[string][]string) (string, bool) {
	fmt.Println("getNextUnvisitedChild for:", from)
	for _, edge := range g.Edges {
		listOfVisitedChildren, ok := visited[from]
		if !ok {
			listOfVisitedChildren = make([]string, 0)
		}
		if edge.From == from && !slices.Contains(listOfVisitedChildren, edge.To) {
			listOfVisitedChildren = append(listOfVisitedChildren, edge.To)
			visited[from] = listOfVisitedChildren
			fmt.Println("Found unvisited child:", edge.To)
			return edge.To, true
		}
	}
	fmt.Println("No unvisited child found")
	return "", false
}

func (g *Graph) selectStartNode(path []Edge) string {
	for _, edge := range g.Edges {
		if slices.Contains(path, *edge) {
			continue
		}
		var hasparent bool
		for _, c := range g.Edges {
			if edge.From == c.To {
				hasparent = true
				break
			}
		}
		if !hasparent {
			fmt.Println("Found start node:", edge.From)
			return edge.From
		}
	}
	return g.Edges[0].From
}

func (g *Graph) checkAndAddToPath(path []Edge, current string, from string) []Edge {
	if from == "" {
		fmt.Println("Adding start node:", current)
		path = append(path, Edge{
			From: from,
			To:   current,
		})
		return path
	}
	fmt.Println("Checking edge:", from, "->", current)
	for _, edge := range g.Edges {
		if edge.From == from && edge.To == current {
			fmt.Println("edge found:", edge.From, "->", edge.To)
			if !slices.Contains(path, *edge) {
				fmt.Println("and edge is not in path, adding it")
				path = append(path, Edge{
					From: from,
					To:   current,
				})
			}
			return path
		}
	}
	return path
}

func (g *Graph) unprocessedEdgesAvailible(path []Edge) bool {
	for _, edge := range g.Edges {
		if !slices.Contains(path, *edge) {
			return true
		}
	}
	return false
}

// for tree like this:
// {To: "A"},
// {From: "A", To: "B"},
// {To: "C"},
// {From: "C", To: "D"},
// find all leaves and mark them with *
func (g *Graph) markLeaves(path []Edge) []Edge {
	// Build a map: node -> first index where it appears as 'From'
	fromIndex := make(map[string]int)

	for i, e := range path {
		if e.From != "" {
			// Only record the first appearance
			if _, exists := fromIndex[e.From]; !exists {
				fromIndex[e.From] = i
			}
		}
	}

	for i, edge := range path {
		firstFromIdx, exists := fromIndex[edge.To]
		if !exists || firstFromIdx <= i {
			// If 'To' is never used as 'From' or only used before (or at) current position -> mark as leaf
			path[i].To = edge.To + "*"
		}
	}

	return path
}
