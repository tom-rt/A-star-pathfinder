package main

import (
	"fmt"

	. "pathfinder/types"
	maze "pathfinder/maze"
)

func main() {
	var fileName string = maze.GetFileName()
	if len(fileName) == 0 {
		return
	}

	var maze Maze = maze.CreateMaze(fileName)
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