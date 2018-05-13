package graph

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddVertex(t *testing.T) {
	g := NewGraph(false)

	for i := 0; i < 5; i++ {
		g.AddVertex(i)
	}

	vertices := g.GetVertices()

	assert.Equal(t, len(vertices), 5, "there should be 5 vertices")

	for i := 0; i < len(vertices); i++ {
		assert.Equal(t, vertices[i], i, "vertices in the array should be correct")
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph(false)

	for i := 0; i < 3; i++ {
		g.AddVertex(i)
	}

	g.AddEdge(0, 2)
	g.AddEdge(3, 0)
	g.AddEdge(2, 0)

	vertices := g.GetVertices()

	assert.Equal(t, len(vertices), 4, "there should be 5 vertices")

	for i := 0; i < len(vertices); i++ {
		assert.Equal(t, vertices[i], i, "vertices in the array should be correct")
	}

	edges := g.GetAdjList(0)

	assert.Equal(t, len(edges), 2, "0 has 2 edges")

	assert.Equal(t, edges[0], 2, "0 <-> 2 exists in the graph")
	assert.Equal(t, edges[1], 3, "0 <-> 3 exists in the graph")

	edges = g.GetAdjList(2)

	assert.Equal(t, len(edges), 1, "2 has 1 edge")

	assert.Equal(t, edges[0], 0, "2 <-> 0 exists in the graph")

	edges = g.GetAdjList(3)

	assert.Equal(t, len(edges), 1, "3 has 1 edge")

	assert.Equal(t, edges[0], 0, "3 <-> 0 exists in the graph")

	edges = g.GetAdjList(1)

	assert.Equal(t, len(edges), 0, "1 has 0 edges")
}
