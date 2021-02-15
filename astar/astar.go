package astar

import (
	"fmt"
	"math"
	t "pathfinder/types"
	utils "pathfinder/utils"
	"strconv"
	"time"
)

type node struct {
	costStart int
	costEnd   float64
	cost      float64
	parent    t.Point
}

type list map[t.Point]*node

func getNeighbors(maze t.Maze, point t.Point) (t.Point, t.Point, t.Point, t.Point) {
	var top t.Point = t.Point{X: point.X, Y: point.Y - 1}
	var right t.Point = t.Point{X: point.X + 1, Y: point.Y}
	var bottom t.Point = t.Point{X: point.X, Y: point.Y + 1}
	var left t.Point = t.Point{X: point.X - 1, Y: point.Y}
	return top, right, bottom, left
}

func getBestNode(list list) (t.Point, *node) {
	var bestNode *node
	var bestPoint t.Point
	var found bool = false

	for key, val := range list {
		if !found || val.cost <= bestNode.cost {
			bestNode = val
			bestPoint = key
			found = true
		}
	}

	return bestPoint, bestNode
}

func calcDistance(pointA t.Point, pointB t.Point) float64 {
	var ax float64 = float64(pointA.X)
	var ay float64 = float64(pointA.Y)
	var bx float64 = float64(pointB.X)
	var by float64 = float64(pointB.Y)
	ret := math.Sqrt(((ax - bx) * (ax - bx)) + ((ay - by) * (ay - by)))
	return ret
}

func createNode(maze t.Maze, parent t.Point, closedList list, point t.Point) *node {
	var costStart int
	parentNode, check := closedList[parent]
	if !check {
		costStart = 1
	} else {
		costStart = parentNode.costStart + 1
	}
	var costEnd = calcDistance(maze.End, point)
	var node node = node{
		costStart: costStart,
		costEnd:   costEnd,
		cost:      float64(costStart) + costEnd,
		parent:    parent,
	}
	return &node
}

func analyzePoint(maze t.Maze, openList list, closedList list, parent t.Point, point t.Point) {
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
		newNode := createNode(maze, parent, closedList, point)
		if check {
			if node.cost > newNode.cost {
				openList[point] = newNode
			}
		} else { // sinon, on l'ajoute dans la liste ouverte avec comme parent le noed courant, et on calcule sa qualité.
			openList[point] = newNode
		}
	}
}

func checkNeighbors(maze t.Maze, openList list, closedList list, currPoint t.Point) (list, list) {
	top, right, bottom, left := getNeighbors(maze, currPoint)

	// On regarde tous ses nœuds voisins.
	analyzePoint(maze, openList, closedList, currPoint, top)
	analyzePoint(maze, openList, closedList, currPoint, right)
	analyzePoint(maze, openList, closedList, currPoint, bottom)
	analyzePoint(maze, openList, closedList, currPoint, left)

	return openList, closedList
}

func drawPath(maze t.Maze, lastNode *node, closedList list) int {
	var currNode *node = lastNode
	var currParent t.Point = lastNode.parent
	var cost int = 0

	for currNode.parent.X != maze.Start.X || currNode.parent.Y != maze.Start.Y {
		maze.Maze[currParent.Y] = utils.ReplaceAtIndex(maze.Maze[currParent.Y], 'o', currParent.X)
		currNode = closedList[currNode.parent]
		currParent = currNode.parent
		cost = cost + 1
	}
	return cost
}

// FindPath solves the maze
func FindPath(maze t.Maze) {
	var isOver bool = false
	var openList list = make(list)
	var closedList list = make(list)
	var cost int
	var start time.Time
	var elapsed time.Duration

	start = time.Now()
	// On commence par le nœud de départ, c'est le nœud courant.
	var currPoint t.Point = maze.Start
	openList, closedList = checkNeighbors(maze, openList, closedList, currPoint)

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
			elapsed = time.Since(start)
			cost = drawPath(maze, bestNode, closedList)
			fmt.Println("ASTAR: Chemin trouvé en " + strconv.Itoa(cost) + " coups et " + elapsed.String())
			isOver = true
			continue
		} else {
			openList, closedList = checkNeighbors(maze, openList, closedList, bestPoint)
		}
	}
}
