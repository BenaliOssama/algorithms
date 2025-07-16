package main

import (
	"fmt"
	"os"

	devide "lem-in/devide_ants"
	graphs "lem-in/graphs"
	fl "lem-in/parse_file"
)

var err_msg = "\033[31mERROR\033[0m: \033[33minvalid data format\033[0m"

func main() {
	//##############################################
	/*############ Parsing the File ###############*/

	// declare a new graph
	file_name, err := fl.GetFileName(os.Args[1:])
	if err != nil {
		fmt.Printf("%s\n%v\n", err_msg, err)
		return
	}
	// get data from file
	nest, parsing_log, err := fl.FillTheNest(file_name)
	if err != nil {
		fmt.Printf("%s\n%v\n", err_msg, err)
		return
	}
	//##############################################
	/*############ Making the Graph ##############*/

	graph := graphs.NewGraph()
	graph.NewVerteces(nest.Rooms)
	graph.Start = graph.Verteces[nest.Start]
	graph.End = graph.Verteces[nest.End]
	// creat edges relations betwen vertexes
	err = graph.ConnectRooms(nest.Tunels)
	if err != nil {
		fmt.Printf("%s\n%v\n", err_msg, err)
		return
	}
	//##################################################
	/*############# Find the Best Paths ##############*/

	/*######## Use BFS ########*/
	simple_paths := graph.AllPaths(graph.Start, graph.End, false)
	simple, fsteps, err := devide.Devide(simple_paths, nest.Ants, graph.End.Name)
	fmt.Println(simple_paths)
	if err != nil {
		fmt.Printf("%s\n%v\n", err_msg, err)
		return
	}
	/*####### Use B-Edmonds-Carp ########*/
	graph.EdmondsKarp()
	edmonds := graph.AllPaths(graph.End, graph.Start, true)
	carp, lsteps, err := devide.Devide(edmonds, nest.Ants, graph.End.Name)
	fmt.Println(edmonds)
	if err != nil {
		fmt.Printf("%s\n%v\n", err_msg, err)
		return
	}
	//####################################################################
	/*############# Compare paths And Devide The Ants ##################*/

	if fsteps <= lsteps {
		fmt.Println(parsing_log)
		devide.Print(simple)
	} else {
		fmt.Println(parsing_log)
		devide.Print(carp)
	}
}
