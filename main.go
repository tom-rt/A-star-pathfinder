package main

import (
	"fmt"
	maze "pathfinder/maze"
	. "pathfinder/types"
)

func getNeighbors(maze Maze, point Point) (Point, Point, Point, Point) {
	var top Point = Point{X: point.X, Y: point.Y - 1}
	var right Point = Point{X: point.X + 1, Y: point.Y}
	var bottom Point = Point{X: point.X, Y: point.Y + 1}
	var left Point = Point{X: point.X - 1, Y: point.Y}
	return top, right, bottom, left
}

func getBestNode(list List) (Point, *Node) {
	var bestNode *Node
	var bestPoint Point
	var found bool = false

	for key, val := range list {
		if !found || val.Cost < bestNode.Cost {
			bestNode = val
			bestPoint = key
			found = true
		}
	}

	return bestPoint, bestNode
}

func calcDistance(pointA Point, pointB Point) int {
	var abs int
	dist := (pointA.X - pointB.X) + (pointA.X - pointB.X)
	if dist < 0 {
		abs = dist * -1
	} else {
		abs = dist
	}
	return abs
}

func createNode(maze Maze, parent Point, point Point) Node {
	var costStart = calcDistance(maze.Start, point)
	var costEnd = calcDistance(maze.End, point)
	var n Node = Node {
		CostStart: costStart,
		CostEnd: costEnd,
		Cost: costStart + costEnd,
		Parent: parent,
	}
	return n
}

func analyzePoint(maze Maze, openList List, closedList List, parent Point, point Point) {
	var check bool
	// Est-ce un obstacle ? Si oui, on oublie ce nœud ;
	if maze.Maze[point.Y][point.X] == '*' {
		return
	}

	// Est-il dans la liste fermée ? Si oui, ce nœud a déjà été étudié ou bien est en cours d'étude, on ne fait rien ;
	_, check = closedList[point]
	if check {
		return
	}

	// Est-il dans la liste ouverte ? Si oui, on calcule la qualité de ce nœud, et si elle est meilleure que celle de son homologue dans la liste ouverte, on modifie le nœud présent dans la liste ouverte ;
	node, check := openList[point]
	newNode := createNode(maze, parent, point)
	if check {
		if (node.Cost > newNode.Cost) {
			openList[point] = &newNode
		}
	} else { // sinon, on l'ajoute dans la liste ouverte avec comme parent le noed courant, et on calcule sa qualité.
		openList[point] = &newNode
	}
}

func tmp(maze Maze, openList List, closedList List, currPoint Point) (List, List) {
	top, right, bottom, left := getNeighbors(maze, currPoint)
	
	// On regarde tous ses nœuds voisins.
	analyzePoint(maze, openList, closedList, currPoint, top)
	analyzePoint(maze, openList, closedList, currPoint, right)
	analyzePoint(maze, openList, closedList, currPoint, bottom)
	analyzePoint(maze, openList, closedList, currPoint, left)

	return openList, closedList
}

func findPath(maze Maze) {
	var isOver bool = false
	var outcome string = ""
	var openList map[Point]*Node = make(List)
	var closedList map[Point]*Node = make(List)
	// var openList map[Point]*Node = map[Point]Node
	// var closedList map[Point]*Node = map[Point]Node

	// On commence par le nœud de départ, c'est le nœud courant.
	var currPoint Point = maze.Start
	openList, closedList = tmp(maze, openList, closedList, currPoint)

	for isOver == false {
		// On cherche le meilleur nœud de toute la liste ouverte. Si la liste ouverte est vide, il n'y a pas de solution, fin de l'algorithme.
		if len(openList) == 0 {
			outcome = "Pas de solution."
			isOver = true
			continue
		}
		bestPoint, bestNode := getBestNode(openList)

		// On le met dans la liste fermée et on le retire de la liste ouverte.
		closedList[bestPoint] = bestNode

		// On réitère avec ce nœud comme nœud courant jusqu'à ce que le nœud courant soit le nœud de destination.
		if bestPoint.X == maze.End.X && bestPoint.Y == maze.End.Y {
			outcome = "Chemin trouvé !"
			isOver = true
			continue
		} else {
			openList, closedList = tmp(maze, openList, closedList, bestPoint)
		}
	}
	fmt.Println(outcome)
	for key, _ := range closedList {
		fmt.Println(key)
		// maze.Maze[key.Y][key.X] = '$'
		maze.Maze[key.Y] = replaceAtIndex(maze.Maze[key.Y], 'o', key.X)
		// fmt.Println(maze.Maze[key.Y][key.X])
	}

	printMaze(maze)
}

func printMaze(maze Maze) {
	for i:=0; i< len(maze.Maze); i++ {
		fmt.Println(maze.Maze[i])
	}
}

func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}

func main() {
	var fileName string = maze.GetFileName()
	if len(fileName) == 0 {
		return
	}

	var maze Maze = maze.CreateMaze(fileName)
	printMaze(maze)
	findPath(maze)
	
	return
}