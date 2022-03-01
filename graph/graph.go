package graph

// Graph representation using adj. list
type Graph struct {
	vertices []*Vertex
	edges    map[Vertex][]*Vertex
}

type Vertex struct {
	value interface{}
}

// Graph representation using adj. matrix
type GraphMatrix struct {
	numNodes int
	edges    [][]int
}

// DGraphMatrix representation using adj. matrix
type DGraphMatrix struct {
	numNodes int
	edges    [][]edge
}

type edge struct {
	from   int
	to     int
	weight int
}
