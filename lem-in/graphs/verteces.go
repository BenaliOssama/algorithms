package graphs

import (
	"errors"
	"fmt"
)

type Vertex struct {
	Name             string
	adjacentVerteces []*Vertex
	visited          bool
}

func NewVertex(name string) *Vertex {
	return &Vertex{
		Name:             name,
		adjacentVerteces: []*Vertex{},
		visited:          false,
	}
}

// //////////////////////////////////////////////////////////////////
// Add_adjacent_vertex adds an adjacent vertex to the current vertex.
// It ensures that each vertex is only added once in both directions.
func (v *Vertex) AddAdjacentVertex(vertex *Vertex) error {

	if vertex == nil {
		return errors.New("vertex pointer is nil")
	}

	// Check if the current vertex already includes the adjacent vertex
	alreadyConnected, err := v.include(vertex)
	if err != nil {
		return err
	}

	if !alreadyConnected {
		v.adjacentVerteces = append(v.adjacentVerteces, vertex)
	}

	// Ensure the adjacent vertex points back to the current vertex
	alreadyConnected, err = vertex.include(v)
	if err != nil {
		return err
	}
	if !alreadyConnected {
		vertex.adjacentVerteces = append(vertex.adjacentVerteces, v)
	}

	return nil
}

// include checks if the vertex is already related to the given vertex.
func (v *Vertex) include(vertex *Vertex) (bool, error) {
	if v == nil {
		return false, errors.New("source pointer is nil")
	}
	if vertex == nil {
		return false, errors.New("destination pointer is nil")
	}

	// Check if the vertex is in the adjacent vertices list
	for _, e := range v.adjacentVerteces {
		if e == vertex {
			return true, nil
		}
	}
	return false, nil
}

// the name of a vertex can not be repeated
func (g *Graph) NewVerteces(names []string) {
	for _, name := range names {
		g.Verteces[name] = NewVertex(name)
	}
}

// this is used to connect all rooms of the graph at onece
// it can be used after we add all rooms to the graph
func (g *Graph) ConnectRooms(edges [][2]string) error {
	for _, cols := range edges {
		// Ensure both vertices exist in the snap map
		if vertex1, ok1 := g.Verteces[cols[0]]; ok1 {
			if vertex2, ok2 := g.Verteces[cols[1]]; ok2 {
				if err := vertex1.AddAdjacentVertex(vertex2); err != nil {
					return fmt.Errorf("%w: %s -> %s", err, cols[0], cols[1])
				}
			} else {
				return fmt.Errorf("vertex %s not found in snap", cols[1])
			}
		} else {
			return fmt.Errorf("vertex %s not found in snap", cols[0])
		}
	}
	return nil
}
