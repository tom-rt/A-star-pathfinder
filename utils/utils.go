package utils

import (
	"bufio"
	"log"
	"os"
	"strings"

	t "pathfinder/types"
)

// CreateMazeFromFile parses a map in a file and returns a Maze structure
func CreateMazeFromFile(fileName string) t.Maze {
	var m t.Maze
	var startX int
	var endX int

	file, err := os.Open(fileName)
	var line string = ""
	if err != nil {
		log.Fatal(fileName, " File does not exist:", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line = scanner.Text()
		startX = strings.Index(line, "D")
		endX = strings.Index(line, "A")

		if startX != -1 {
			m.Start = t.Point{X: startX, Y: len(m.Maze)}
		}

		if endX != -1 {
			m.End = t.Point{X: endX, Y: len(m.Maze)}
		}

		m.Maze = append(m.Maze, line)
	}
	return m
}

// GetParams parses parameters
func GetParams() (string, bool) {
	len := len(os.Args)
	if len == 2 {
		fileName := os.Args[1]
		return fileName, false
	} else if len == 3 && os.Args[1] == "-p" {
		fileName := os.Args[2]
		return fileName, true
	} else {
		return "", false
	}
}

// ReplaceAtIndex replaces a char at the given index of a string
func ReplaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}