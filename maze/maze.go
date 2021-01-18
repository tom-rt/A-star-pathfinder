package maze

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	. "pathfinder/types"
)

func CreateMaze(fileName string) Maze {
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
		startX = strings.Index(line, "x")
		endX = strings.Index(line, "o")

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

func GetFileName() string {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a file name.")
		return ""
	}
	fileName := os.Args[1]
	return fileName
}