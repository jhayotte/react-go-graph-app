package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Node struct {
	ID        string                 `json:"id"`
	Type      string                 `json:"type,omitempty"`
	Data      map[string]interface{} `json:"data"`
	Position  map[string]float64     `json:"position"`
	ClassName string                 `json:"className"`
}

type Edge struct {
	ID       string `json:"id"`
	Source   string `json:"source"`
	Target   string `json:"target"`
	Animated bool   `json:"animated,omitempty"`
}

type GraphData struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

func graphHandler(w http.ResponseWriter, r *http.Request) {
	nodes := []Node{
		{ID: "1", Type: "input", Data: map[string]interface{}{"label": "Input Node"}, Position: map[string]float64{"x": 250, "y": 5}, ClassName: "react-flow__node-input"},
		{ID: "2", Data: map[string]interface{}{"label": "Default Node"}, Position: map[string]float64{"x": 100, "y": 100}, ClassName: "react-flow__node-default"},
		{ID: "3", Data: map[string]interface{}{"label": "Output Node"}, Position: map[string]float64{"x": 400, "y": 100}, ClassName: "react-flow__node-output"},
	}
	edges := []Edge{
		{ID: "e1-2", Source: "1", Target: "2", Animated: true},
		{ID: "e2-3", Source: "2", Target: "3"},
	}

	graphData := GraphData{Nodes: nodes, Edges: edges}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(graphData)
}

func main() {
	fs := http.FileServer(http.Dir("./frontend/build"))
	http.Handle("/", fs)
	http.HandleFunc("/api/graph", graphHandler)

	log.Println("Server starting on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
