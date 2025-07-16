package graphs

import (
	"errors"
	"fmt"

	"lem-in/queue"
	"lem-in/utils"
)

type Graph struct {
	Verteces map[string]*Vertex
	Start    *Vertex // the starting room
	End      *Vertex // the ending room
}

func NewGraph() *Graph {
	return &Graph{Verteces: make(map[string]*Vertex), Start: nil, End: nil}
}

// add a vertex to the graph
func (g *Graph) Add(v *Vertex) {
	g.Verteces[v.Name] = v
}

// BFS to find an augmenting path between from and to
func (g *Graph) BFS(from, to *Vertex, visited map[string]bool, reverse bool) []string {
	// parent := make(map[*Vertex]*Vertex)
	// parent := [][2]string{}
	parent := make(map[*Vertex]*Vertex)
	q := queue.New() // Using a simple slice as a queue
	q.Enqueue(from)
	visited[from.Name] = true

	// If we reach the Start node, we can construct the path

	for !q.IsEmpty() {

		current := q.Dequeue().Item.(*Vertex)
		if current == to {
			// return assemble(parent, to.Name)
			path := constructPath(parent, to, reverse)
			// case of the start connected to the end
			if len(path) == 2 {
				// break the connection forward
				err := g.breakEndStart()
				if err != nil {
					fmt.Println(err)
					return nil
				}
				return path
			}
			return path

		}

		for _, neighbor := range current.adjacentVerteces {
			if !visited[neighbor.Name] { // Not visited
				q.Enqueue(neighbor) // Enqueue
				visited[neighbor.Name] = true
				// parent[current] = neighbor
				parent[neighbor] = current

			}
		}
	}

	return nil
}

func (g *Graph) AllPaths(from, to *Vertex, reverse bool) [][]string {
	visited := make(map[string]bool)
	paths := [][]string{}
	broke := false
	for {
		path := g.BFS(from, to, utils.CopyMap(visited), reverse)
		if len(path) < 1 {
			break
		}
		paths = append(paths, path)
		if len(path) != 2 {
			for _, v := range path {
				if v != to.Name {
					visited[v] = true
				}
			}
		} else {
			broke = true
		}

	}
	if broke {
		g.Start.AddAdjacentVertex(g.End)
	}
	return paths
}

// Helper function to reconstruct the path from parent map
func constructPath(parent map[*Vertex]*Vertex, from *Vertex, reverse bool) []string {
	var path []string
	for v := from; v != nil; v = parent[v] { // Fix here: start from end
		if reverse {
			path = append([]string{v.Name}, path...) // Prepend the node
		} else {
			path = append(path, v.Name) // Prepend the node
		}
	}
	return path
}

func (g *Graph) breakEndStart() error {
	for i := 0; i < len(g.Start.adjacentVerteces); i++ {
		if g.Start.adjacentVerteces[i] == g.End {
			g.Start.adjacentVerteces = append(g.Start.adjacentVerteces[:i], g.Start.adjacentVerteces[i+1:]...)
		}
	}
	for i := 0; i < len(g.End.adjacentVerteces); i++ {
		if g.End.adjacentVerteces[i] == g.Start {
			g.End.adjacentVerteces = append(g.End.adjacentVerteces[:i], g.End.adjacentVerteces[i+1:]...)
			return nil
		}
	}

	return errors.New("not start end connection found")
}
