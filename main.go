package main

import (
	"fmt"
	"pathfinder/dijkstra"
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
	fileName, print := utils.GetParams()
	if len(fileName) == 0 {
		return
	}

	maze = utils.CreateMazeFromFile(fileName)
	// astar.FindPath(maze)
	dijkstra.FindPath(maze)
	if print {
		printMaze(maze.Maze)
	}

	return
}