package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"dfs/graph"
)

func main() {
	// Enable CORS
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", corsMiddleware(fs))

	// Handle DFS calculation
	http.HandleFunc("/calculate-dfs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var edges []graph.Edge
		if err := json.NewDecoder(r.Body).Decode(&edges); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Convert to pointer slice for Graph
		edgePtrs := make([]*graph.Edge, len(edges))
		for i := range edges {
			edgePtrs[i] = &edges[i]
		}

		g := &graph.Graph{}
		g.AddEdges(edgePtrs)
		result := g.DFS()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Server starting at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
