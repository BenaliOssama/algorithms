package graphs

import (
	"lem-in/queue"
)

type Node struct {
	Vx    *Vertex
	Panic bool
}

func NewNode(v *Vertex, panic bool) Node {
	return Node{
		Vx:    v,
		Panic: panic,
	}
}

/*
this algorithm is a modification of the edmonds-carp algorith to adapte it to the lem-in project.
the idea is based on forcing the bfs function to take a reverse path at some moment if it incounter an already used vertex
in other word. as long as we cant use a node for two path. the path that incounter the vertex that is in use is useless
anless that path took a reverse way and then found it's way to the exit again.
in that we can break the lenk of the path and we get two ways instead of one.
*/

// BFS to find an augmenting path
func (g *Graph) bfs(parent map[*Vertex]*Vertex) bool {
	visited := make(map[*Vertex]bool)
	q := queue.New() // Using a simple slice as a queue

	q.Enqueue(NewNode(g.Start, false))
	visited[g.Start] = true

	for !q.IsEmpty() {
		current := q.Dequeue().Item.(Node)
		// check if the current node is in panic mode
		for i, neighbor := range current.Vx.adjacentVerteces {
			if !visited[neighbor] { // Not visited
				// if i'm in panic. i'm looking for exite
				// only look for visited nodes
				if current.Panic { // panic mode
					switch neighbor.visited {
					case true: // found a potential exit
						// exit panic mode
						// i found an exit i don't have to panic in this loop
						// pc = false
						q.Enqueue(NewNode(neighbor, false)) // Enqueue
					case false: // no exit found
						if i == len(current.Vx.adjacentVerteces)-1 {
							visited[current.Vx] = false // free this node for latter use
						}
						continue
					}
				} else { // normal mode
					// enter panic mode. node is visited . or inhreting panic
					if neighbor.visited { // panic can be inhereted
						q.Enqueue(NewNode(neighbor, true))
					} else {
						q.Enqueue(NewNode(neighbor, false)) // Enqueue
					}
				}
				visited[neighbor] = true
				parent[neighbor] = current.Vx
				if neighbor == g.End {
					return true
				}
			}
		}
	}
	return false
}

// Edmonds-Karp algorithm to find maximum flow
func (g *Graph) EdmondsKarp() int {
	totalFlow := 0
	parent := make(map[*Vertex]*Vertex)

	for g.bfs(parent) {
		// Found an augmenting path
		totalFlow++
		// Update the edges in the path
		for v := g.End; v != g.Start; v = parent[v] {

			u := parent[v]
			u.visited = true
			// Remove the forward edge
			for i, adjacent := range u.adjacentVerteces {
				if adjacent == v {
					u.adjacentVerteces = append(u.adjacentVerteces[:i], u.adjacentVerteces[i+1:]...)
					break
				}
			}
		}
	}
	return totalFlow
}
