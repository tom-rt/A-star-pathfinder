package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x int
	y int
}

type Node struct {
	costStart int
	costEnd int
	costSum int
	parent Point
}

func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	var line string = ""
	var maze []string
	if err != nil {
		log.Fatal(fileName, " File does not exist:", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line = scanner.Text()
		maze = append(maze, line)
	}
	return maze
}

func getFileName() string {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a file name.")
		return ""
	}
	fileName := os.Args[1]
	return fileName
}

func main() {
	var fileName string = getFileName()
	if len(fileName) == 0 {
		return
	}
	var maze []string = readFile(fileName)
	fmt.Println(maze)

	var openList map[Point]Node = make(map[Point]Node)
	var closedList map[Point]Node = make(map[Point]Node)
	var p Point = Point{1,2}
	var n Node = Node{1, 3, 4, Point{1,2}}

	openList[p] = n
	closedList[p] = n
	fmt.Println(openList, closedList)
	return
}