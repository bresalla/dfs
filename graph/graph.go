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
	Edges []*Edge
}

func (g *Graph) AddEdges(edges []*Edge) {
	for _, edge := range edges {
		g.Edges = append(g.Edges, edge)
	}
}

func (g *Graph) dfs() []Edge {
	var stack *Stack = NewStack()
	visited := make(map[string][]string)
	path := []Edge{}
	var from string
	for g.UnprocessedEdgesAvailible(path) {
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
			if child, ok := g.GetFirstUnvisitedChild(current, visited); ok {

				stack.Push(current) // Push current node back to stack
				fmt.Println("Push:", current)
				stack.Push(child) // Push child to stack
				fmt.Println("Push:", child)
				visited[current] = append(visited[current], child)
			}
		}
	}

	return path
}

func (g *Graph) GetFirstUnvisitedChild(from string, visited map[string][]string) (string, bool) {
	fmt.Println("GetFirstUnvisitedChild:", from)
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
			fmt.Println("Edge found:", edge.From, "->", edge.To)
			path = append(path, Edge{
				From: from,
				To:   current,
			})
			return path
		}
	}
	return path
}

func (g *Graph) UnprocessedEdgesAvailible(path []Edge) bool {
	for _, edge := range g.Edges {
		if !slices.Contains(path, *edge) {
			return true
		}
	}
	return false
}
