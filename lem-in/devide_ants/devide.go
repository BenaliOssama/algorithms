package devide

import (
	"errors"
	"fmt"

	"strings"
)

type Path struct {
	Rooms     []string
	Passenger int
}

func NewPath() *Path {
	return &Path{Rooms: []string{}, Passenger: 0}
}

func (p *Path) Add(room string) {
	(*p).Rooms = append((*p).Rooms, room)
}

func (p *Path) Pass() {
	p.Passenger++
}

// ////////////////////////////////////////////////////////
type Paths []*Path

// add to the path ?
// declare Paths
func NewPaths() *Paths {
	return &Paths{}
}

func (p *Paths) Append(path *Path) {
	*p = append(*p, path)
}

func (p *Paths) Shortest() (*Path, error) {
	if len(*p) == 0 {
		return nil, errors.New("no path found")
	}
	// var shortest *Path
	shortest := (*p)[0]
	for _, path := range *p {
		if path.Passenger < shortest.Passenger {
			shortest = path
		}
	}
	return shortest, nil
}

func Devide(ways [][]string, ants int, end string) ([][]string, int, error) {
	// make new paths
	paths := NewPaths()
	// fil paths with every path and how many passengers it takes
	for _, p := range ways {
		// make new path
		path := NewPath()
		for _, room := range p {
			path.Add(room)
		}
		path.Passenger = len(path.Rooms) - 2
		paths.Append(path)
	}
	// show is what i show to the user
	show := [][]string{}
	special := []string{}
	took := []string{}
	// give all the paths their ants
	for i := 1; i <= ants; i++ {
		short, err := paths.Shortest()
		if err != nil {
			return nil, 0, err
		}
		if len(short.Rooms) == 2 {
			special = append(special, antpath(i, short.Rooms)...)
		} else {
			took = antpath(i, short.Rooms)
			show = append(show, took)
		}
		short.Pass()
	}
	show = append(show, special)
	max := MaxSteps(paths)

	mat := Retate(show, end)

	return mat, max, nil
}

func antpath(ant int, path []string) []string {
	took := []string{}
	// why not printing the first room
	for i := len(path) - 2; i >= 0; i-- {
		room := path[i]
		took = append(took, fmt.Sprintf("L%d-%s", ant, room))
	}
	return took
}

// change the view from horizontal to vertical
func Retate(matrix [][]string, end string) [][]string {

	result := [][]string{}
	stop := true
	y := 0
	for stop {
		stop = false
		line := []string{}
		for i := 0; i < len(matrix); i++ {
			// make a special case for branches with len 1
			branch := matrix[i]

			if len(branch) > y {
				stop = true
				if y == len(branch)-1 {
					line = append(line, branch[y])
				} else if !Check(line, branch[y], end) {
					line = append(line, branch[y])
				} else {
					matrix[i] = append([]string{""}, branch...)
				}
			}
		}
		result = append(result, line)
		y++
	}
	return result
}

// to not use the same room at the same time
func Check(words []string, word string, end string) bool {
	// Ensure the input word is not empty
	if len(word) == 0 {
		return false
	}

	// Get the last character of the word
	lastTerm := strings.Split(word, "-")[1]

	// Compare with the last characters of words in the slice
	for _, w := range words {
		// Ensure the word in the slice is not empty
		if len(w) == 0 {
			continue
		}
		// Get the last character of the word in the slice
		lastTermInSlice := strings.Split(w, "-")[1]

		// Compare the last characters
		if lastTerm == lastTermInSlice && lastTermInSlice != end {
			return true
		}
	}

	return false
}

func MaxSteps(paths *Paths) int {
	max := 0
	for _, v := range *paths {
		if v.Passenger >= max {
			max = v.Passenger
		}
	}
	return max
}

// desplay the result to on the terminal
func Print(mat [][]string) {

	for _, line := range mat {
		if len(line) == 0 {
			continue
		}
		for i, word := range line {
			if word != "" {
				fmt.Print(word)
				if i != len(line)-1 {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}
