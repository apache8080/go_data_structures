package graph

import (
	"errors"

	"github.com/apache8080/go_data_structures/list"
)

type Graph struct {
	adjLists map[int]*list.List
	vertices []int
	directed bool
}

func NewGraph(directed bool) *Graph {
	graph := new(Graph)
	graph.directed = directed
	graph.adjLists = make(map[int]*list.List)
	return graph
}

func (this *Graph) IsVertex(vertex int) bool {
	for i := 0; i < len(this.vertices); i++ {
		if this.vertices[i] == vertex {
			return true
		}
	}

	return false
}

func (this *Graph) GetAdjList(vertex int) []int {
	if this.IsVertex(vertex) {
		return this.adjLists[vertex].GetArray()
	}

	return []int{}
}

func (this *Graph) GetVertices() []int {
	return this.vertices
}

func (this *Graph) AddVertex(vertex int) error {
	if !this.IsVertex(vertex) {
		this.vertices = append(this.vertices, vertex)
		this.adjLists[vertex] = list.NewList()
		return nil
	} else {
		return errors.New("Vertex already exists in the graph.")
	}
}

func (this *Graph) AddEdge(start int, end int) {
	this.AddVertex(start)
	this.AddVertex(end)

	if !this.IsEdge(start, end) {
		this.adjLists[start].Append(end)
		if !this.directed {
			this.adjLists[end].Append(start)
		}
	}
}

func (this *Graph) IsEdge(start int, end int) bool {
	adj := this.adjLists[start].GetArray()

	exists := false
	
	for i := 0; i < len(adj); i++ {
		if (adj[i] == end) {
			exists = true
		}
	}
	
	if exists && !this.directed {
		adj := this.adjLists[start].GetArray()

		for i := 0; i < len(adj); i++ {
			if adj[i] == end {
				return true
			}
		}
	}

	return exists
}

