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

func calcNode(maze Maze, point Point) Node {
	var n Node = Node{1, 3, 4, Point{1,2}}
	return n
}

func analyzePoint(maze Maze, openList map[Point]Node, closedList map[Point]Node, point Point) {
	var check bool

	// est-ce un obstacle ? Si oui, on oublie ce nœud ;
	if maze.Maze[point.Y][point.X] == '*' {
		return
	}

	// est-il dans la liste fermée ? Si oui, ce nœud a déjà été étudié ou bien est en cours d'étude, on ne fait rien ;
	_, check = closedList[point]
	if check {
		return
	}

	_, ok := openList[point]
	fmt.Println("ici", ok)
	var n Node = Node{1, 3, 4, Point{1,2}}
	openList[point] = n
	_, ok = openList[point]
	fmt.Println("la", ok)


	// if closed{
	// 	fmt.Println("Wall")
	// }

	// est-il dans la liste ouverte ? Si oui, on calcule la qualité de ce nœud, et si elle est meilleure que celle de son homologue dans la liste ouverte, on modifie le nœud présent dans la liste ouverte ;
	node, check := openList[point]
	if check {
		newNode := calcNode(maze, point) 
		if (node.CostSum > newNode.CostSum) {
			return
		}
		return
	}

	// sinon, on l'ajoute dans la liste ouverte avec comme parent le noed courant, et on calcule sa qualité.

	// var p Point = Point{1,2}
	// closedList[p] = n

}


func solve(maze Maze) {
	var openList map[Point]Node = make(map[Point]Node)
	var closedList map[Point]Node = make(map[Point]Node)

	// On commence par le nœud de départ, c'est le nœud courant.
	var currPoint Point = maze.Start
	top, right, bottom, left := getNeighbors(maze, currPoint)
	
	// On regarde tous ses nœuds voisins.
	analyzePoint(maze, openList, closedList, top)
	analyzePoint(maze, openList, closedList, right)
	analyzePoint(maze, openList, closedList, bottom)
	analyzePoint(maze, openList, closedList, left)

	fmt.Println(openList, closedList)
	// On cherche le meilleur nœud de toute la liste ouverte. Si la liste ouverte est vide, il n'y a pas de solution, fin de l'algorithme.
	// On le met dans la liste fermée et on le retire de la liste ouverte.
	// On réitère avec ce nœud comme nœud courant jusqu'à ce que le nœud courant soit le nœud de destination.

}

func main() {
	var fileName string = maze.GetFileName()
	if len(fileName) == 0 {
		return
	}

	var maze Maze = maze.CreateMaze(fileName)
	solve(maze)
	
	return
}