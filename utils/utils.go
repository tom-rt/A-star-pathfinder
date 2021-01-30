package utils

import (
	"bufio"
	"log"
	"os"
	"strings"

	. "pathfinder/types"
)

func CreateMazeFromFile(fileName string) Maze {
	var m Maze
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
			m.Start = Point{startX, len(m.Maze)}
		}

		if endX != -1 {
			m.End = Point{endX, len(m.Maze)}
		}

		m.Maze = append(m.Maze, line)
	}
	return m
}

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