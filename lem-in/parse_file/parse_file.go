package parse_file

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Nest struct {
	Rooms  []string
	Tunels [][2]string
	Start  string
	End    string
	Ants   int
}

// return the file name
func GetFileName(args []string) (string, error) {
	if len(args) != 1 { // check the arg of program
		return "", errors.New("usage: go run . [file.txt] or lem-in [file.txt]")
	}
	arg := args[0]
	if !strings.HasSuffix(arg, ".txt") { // the the extension of arg program
		return "", errors.New("format: only txt extension files")
	}
	// return the file name
	return arg, nil
}

// the function that read the file and call the parsing to parse the file
func FillTheNest(filename string) (Nest, string, error) {
	var nest Nest
	var log string
	file, err := os.Open(filename)
	if err != nil {
		return nest, log, errors.New("file: check your file")
	}
	defer file.Close()

	result := []string{}
	Scanner := bufio.NewScanner(file)
	for Scanner.Scan() {
		line := strings.TrimSpace(Scanner.Text())
		if len(line) == 0 { // any empty line call an empty comment so ingore it
			continue
		}
		result = append(result, line)
	}
	nest, log, err = Parse(result)
	return nest, log, err
}

// the parsing function that fild a Nest struct and handle the error
func Parse(result []string) (Nest, string, error) {
	var nest Nest
	var log string
	var err error
	if len(result) == 0 { // check the file is not empty
		return nest, log, errors.New("file: empty file")
	}
	// search for ants in file
	if len(result) > 1 {
		nest.Ants, err = strconv.Atoi(result[0])
		if err != nil || nest.Ants <= 0 {
			return nest, log, errors.New("parsing: invalid number of ant")
		}
		log += fmt.Sprintf("%d\n", nest.Ants)

	}

	// filde all rooms by reading the result or lines
	for i := 1; i < len(result); i++ {
		arg := result[i]
		// Check for start and end indicators
		if strings.HasPrefix(arg, "##") {
			// Ensure there's a next element is valid
			if i == len(result)-1 {
				continue
			}
			nextArg := result[i+1] // Store next argument for easier access
			switch arg {
			case "##start": // in case of finding start tag
				tunel, err := GetRoom(strings.Fields(nextArg))
				if err {
					continue
				} else {
					// filed the start room
					nest.Start = tunel
					nest.Rooms = append(nest.Rooms, tunel)
					// Remove the start indicator and its following element
					result = append(result[:i], result[i+2:]...)
					i-- // Adjust index after modification
				}
				log += arg + "\n" + nextArg + "\n"
			case "##end": // in case of finding end tag
				tunel, err := GetRoom(strings.Fields(nextArg))
				if err {
					continue
				} else {
					// filed the end room
					nest.End = tunel
					nest.Rooms = append(nest.Rooms, tunel)
					// Remove the start indicator and its following element
					result = append(result[:i], result[i+2:]...)
					i-- // Adjust index after modification
				}
				log += arg + "\n" + nextArg + "\n"
			}
		}

		// tor wich is the normal line to test or to read
		Tor := strings.Fields(arg)
		switch len(Tor) {
		case 3: // requirement for room param is 3 word
			tunel, err := GetRoom(strings.Fields(arg))
			if err {
				return nest, log, fmt.Errorf("parsing: invalid flag [%s]", arg)
			}
			// case room
			nest.Rooms = append(nest.Rooms, tunel)
			log += arg + "\n"
		case 1: // requirement for tunel param is 1 word
			tunel, err := GetTunel(arg)
			if !err && len(tunel[0]) != 0 && len(tunel[1]) != 0 {
				nest.Tunels = append(nest.Tunels, tunel)
				log += arg + "\n"
			} else if err && nest.Ants != 0 && !strings.HasPrefix(arg, "#") {
				return nest, log, fmt.Errorf("parsing: invalid flag [%s]", arg)
			} else if strings.HasPrefix(arg, "#") && arg != "##start" && arg != "##end" {
				log += arg + "\n"
			}
		}
	}

	err = ValidData(nest)
	if err != nil {
		return nest, log, err
	} else {
		return nest, log, nil
	}
}

func GetRoom(room []string) (string, bool) {
	// requirement for a room is to not start with "L" or "#" and got 3 word
	if len(room) != 3 || len(room[0]) < 1 {
		return "", true
	}
	// the second and third word must be cordinat wich is int
	_, err := strconv.Atoi(room[1])
	if err != nil {
		return "", true
	}
	_, err = strconv.Atoi(room[2])
	if err != nil {
		return "", true
	}
	// the name wich is neccesary must be the first word
	return room[0], false
}

func GetTunel(tunel string) ([2]string, bool) {
	// a tunels must be compose by two word separed by "-"
	t := strings.Split(tunel, "-")
	if len(t) != 2 {
		return [2]string{}, true
	}
	return [2]string{t[0], t[1]}, false
}

func ValidData(data Nest) error {
	if data.Start == "" || data.End == "" { // check if start or end room not found
		return errors.New("parsing: missing starting or ending room")
	}
	if len(data.Tunels) < 1 { // check if ther is at leas one tunel
		return errors.New("parsing: no tunel found")
	}
	for _, value := range data.Rooms {
		if strings.HasPrefix(value, "L") || strings.HasPrefix(value, "#") {
			return errors.New("parsing: invalid room name")
		}
	}
	for _, value := range data.Tunels {
		if value[0] == value[1] {
			return errors.New("parsing: room linked to him self")
		}
	}
	return nil
}
