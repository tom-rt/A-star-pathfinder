package main

import (
	"fmt"
	_ "math"
	"pathfinder/astar"
	. "pathfinder/types"
	utils "pathfinder/utils"
)

func printMaze(maze []string) {
	for i:=0; i< len(maze); i++ {
		fmt.Println(maze[i])
	}
}

func main() {
	var maze Maze
	var fileName string = utils.GetFileName()
	if len(fileName) == 0 {
		return
	}

	maze = utils.CreateMazeFromFile(fileName)
	printMaze(maze.Maze)
	astar.FindPathInteger(maze)
	printMaze(maze.Maze)

	maze = utils.CreateMazeFromFile(fileName)
	printMaze(maze.Maze)

	return
}