package graph

type Graph[V comparable] struct {
	graph     map[V][]V
	isVisited map[V]struct{}
	buf       []V
	isBFS     bool
}

func NewGraph[V comparable]() *Graph[V] {
	//
	return &Graph[V]{
		graph:     make(map[V][]V),
		isVisited: make(map[V]struct{}),
		buf:       make([]V, 0, 8),
	}
}

func (g *Graph[V]) bufin(val V) {
	g.buf = append(g.buf, val)
}

func (g *Graph[V]) bufout() *V {
	if len(g.buf) > 0 {
		if g.isBFS {
			head := g.buf[0]
			g.buf = g.buf[1:]
			return &head
		} else {
			last := len(g.buf) - 1
			tail := g.buf[last]
			g.buf = g.buf[:last]
			return &tail
		}
	}
	return nil
}

func (g *Graph[V]) StartSearch(startVert, endVert V, isBFS bool) bool {
	var void struct{}
	g.isBFS = isBFS

	g.isVisited[startVert] = void
	g.bufin(startVert)
	for len(g.buf) != 0 {
		vert := g.bufout()
		if vert == nil {
			continue
		}
		edges := g.graph[*vert]
		for _, svert := range edges {
			if _, ok := g.isVisited[svert]; !ok {
				g.isVisited[svert] = void
				g.bufin(svert)
			}
		}
		if _, ok := g.isVisited[endVert]; ok {
			return true
		}
	}
	return false
}
