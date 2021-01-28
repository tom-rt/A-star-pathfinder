package astarint

import (
	"fmt"
	_ "math"
	"strconv"
	. "pathfinder/types"
)

func getNeighbors(maze Maze, point Point) (Point, Point, Point, Point) {
	var top Point = Point{X: point.X, Y: point.Y - 1}
	var right Point = Point{X: point.X + 1, Y: point.Y}
	var bottom Point = Point{X: point.X, Y: point.Y + 1}
	var left Point = Point{X: point.X - 1, Y: point.Y}
	return top, right, bottom, left
}

func getBestNode(list ListInt) (Point, *NodeInt) {
	var bestNode *NodeInt
	var bestPoint Point
	var found bool = false

	for key, val := range list {
		if !found || val.Cost <= bestNode.Cost {
			bestNode = val
			bestPoint = key
			found = true
		}
	}

	return bestPoint, bestNode
}

// func calcDistanceFloat(pointA Point, pointB Point) float64 {
// 	return math.Sqrt(((pointA.X - pointB.X) * (pointA.X - pointB.X)) + ((pointA.Y - pointB.Y) * (pointA.Y - pointB.Y)))
// }

func calcDistance(pointA Point, pointB Point) int {
	var abs int
	dist := ((pointA.X - pointB.X) * (pointA.X - pointB.X)) + ((pointA.Y - pointB.Y) * (pointA.Y - pointB.Y))
	if dist < 0 {
		abs = dist * -1
	} else {
		abs = dist
	}
	return abs
}

func createNode(maze Maze, parent Point, parentNode *NodeInt, point Point) *NodeInt {
	fmt.Println(parentNode)
	var distStart int
	if (parentNode == nil){
		distStart = 1
	} else {
		distStart = parentNode.CostStart + 1
	}
	var costStart = distStart
	var costEnd = calcDistance(maze.End, point)
	var node NodeInt = NodeInt {
		CostStart: costStart,
		CostEnd: costEnd,
		Cost: costStart + costEnd,
		Parent: parent,
	}
	return &node
}

func analyzePoint(maze Maze, openList ListInt, closedList ListInt, parent Point, point Point) {
	var check bool
	// Est-ce un obstacle ? Si oui, on oublie ce nœud ;
	if maze.Maze[point.Y][point.X] == '*' {
		return
	}

	// Est-il dans la liste fermée ? Si oui, ce nœud a déjà été étudié ou bien est en cours d'étude, on ne fait rien ;
	_, check = closedList[point]
	if !check {
		// Est-il dans la liste ouverte ? Si oui, on calcule la qualité de ce nœud, et si elle est meilleure que celle de son homologue dans la liste ouverte, on modifie le nœud présent dans la liste ouverte ;
		node, check := openList[point]
		newNode := createNode(maze, parent, closedList[parent], point)
		if check {
			if (node.Cost >= newNode.Cost) {
				openList[point] = newNode
			}
		} else { // sinon, on l'ajoute dans la liste ouverte avec comme parent le noed courant, et on calcule sa qualité.
			openList[point] = newNode
		}
	}
}

func tmp(maze Maze, openList ListInt, closedList ListInt, currPoint Point) (ListInt, ListInt) {
	top, right, bottom, left := getNeighbors(maze, currPoint)

	// On regarde tous ses nœuds voisins.
	analyzePoint(maze, openList, closedList, currPoint, top)
	analyzePoint(maze, openList, closedList, currPoint, right)
	analyzePoint(maze, openList, closedList, currPoint, bottom)
	analyzePoint(maze, openList, closedList, currPoint, left)

	return openList, closedList
}

func drawPath(maze Maze, lastNode *NodeInt, closedList ListInt) int {
	var currNode *NodeInt = lastNode
	var currParent Point = lastNode.Parent
	var cost int = 0

	for currNode.Parent.X != maze.Start.X || currNode.Parent.Y != maze.Start.Y {
		maze.Maze[currParent.Y] = replaceAtIndex(maze.Maze[currParent.Y], 'o', currParent.X)
		currNode = closedList[currNode.Parent]
		currParent = currNode.Parent
		cost = cost + 1
	}
	return cost
}

func FindPath(maze Maze) {
	var isOver bool = false
	var openList ListInt = make(ListInt)
	var closedList ListInt = make(ListInt)
	var cost int

	// On commence par le nœud de départ, c'est le nœud courant.
	var currPoint Point = maze.Start
	openList, closedList = tmp(maze, openList, closedList, currPoint)

	for isOver == false {
		// On cherche le meilleur nœud de toute la liste ouverte. Si la liste ouverte est vide, il n'y a pas de solution, fin de l'algorithme.
		if len(openList) == 0 {
			isOver = true
			fmt.Println("Pas de solution.")
			continue
		}
		bestPoint, bestNode := getBestNode(openList)

		// On le met dans la liste fermée et on le retire de la liste ouverte.
		closedList[bestPoint] = bestNode
		delete(openList, bestPoint)

		// On réitère avec ce nœud comme nœud courant jusqu'à ce que le nœud courant soit le nœud de destination.
		if bestPoint.X == maze.End.X && bestPoint.Y == maze.End.Y {
			cost = drawPath(maze, bestNode, closedList)
			isOver = true
			fmt.Println("INT: Chemin trouvé en " + strconv.Itoa(cost) + " coups.")
			continue
		} else {
			openList, closedList = tmp(maze, openList, closedList, bestPoint)
		}
	}
}


func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}