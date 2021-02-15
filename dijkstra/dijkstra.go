package dijkstra

import (
	"fmt"
	"math"
	t "pathfinder/types"
	utils "pathfinder/utils"
	"strconv"
	"time"
)

type node struct {
	cost   int
	parent t.Point
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

func createNode(maze t.Maze, parent t.Point, total list, point t.Point) *node {
	var cost int

	parentNode, check := total[parent]
	if !check {
		cost = 0
	} else {
		cost = parentNode.cost + 1
	}

	var node node = node{
		cost:   cost,
		parent: parent,
	}

	return &node
}

func analyzePoint(maze t.Maze, total list, parent t.Point, newPoints *[]t.Point, point t.Point) {
	var check bool
	// Est-ce un obstacle ? Si oui, on oublie ce nœud ;
	if maze.Maze[point.Y][point.X] == '*' {
		return
	}

	// Est-il dans la liste ? Si oui, ce nœud a déjà été étudié ou bien est en cours d'étude, on ne fait rien;
	_, check = total[point]
	if !check {
		newNode := createNode(maze, parent, total, point)
		total[point] = newNode
		*newPoints = append(*newPoints, point)
	}
}

func checkNeighbors(maze t.Maze, total list, lastAdded []t.Point) (list, []t.Point) {
	var newPoints []t.Point
	for _, currPoint := range lastAdded {
		top, right, bottom, left := getNeighbors(maze, currPoint)

		// On regarde tous ses nœuds voisins.
		analyzePoint(maze, total, currPoint, &newPoints, top)
		analyzePoint(maze, total, currPoint, &newPoints, right)
		analyzePoint(maze, total, currPoint, &newPoints, bottom)
		analyzePoint(maze, total, currPoint, &newPoints, left)

	}
	return total, newPoints
}

func drawPath(maze t.Maze, lastNode *node, closedList list) {
	var currNode *node = lastNode
	var currParent t.Point = lastNode.parent

	for currNode.parent.X != maze.Start.X || currNode.parent.Y != maze.Start.Y {
		maze.Maze[currParent.Y] = utils.ReplaceAtIndex(maze.Maze[currParent.Y], 'o', currParent.X)
		currNode = closedList[currNode.parent]
		currParent = currNode.parent
	}
}

// FindPath solves the maze
func FindPath(maze t.Maze) {
	var isOver bool = false
	var total list = make(list)
	var lastAdded []t.Point
	var start time.Time
	var elapsed time.Duration

	start = time.Now()
	var currPoint t.Point = maze.Start
	baseNode := createNode(maze, t.Point{X: 0, Y: 0}, total, currPoint)
	total[currPoint] = baseNode
	lastAdded = append(lastAdded, currPoint)

	total, lastAdded = checkNeighbors(maze, total, lastAdded)

	for isOver == false {
		if len(lastAdded) == 0 {
			isOver = true
			fmt.Println("Pas de solution.")
			continue
		}

		// On cherche le meilleur nœud de toute la liste ouverte. Si la liste ouverte est vide, il n'y a pas de solution, fin de l'algorithme.
		for _, point := range lastAdded {
			if point.X == maze.End.X && point.Y == maze.End.Y {
				elapsed = time.Since(start)
				drawPath(maze, total[point], total)
				totalCost := strconv.Itoa(total[total[point].parent].cost)
				fmt.Println("DIJKSTRA: Chemin trouvé en " + totalCost + " coups et " + elapsed.String())
				isOver = true
				continue
			}
		}

		total, lastAdded = checkNeighbors(maze, total, lastAdded)

	}

}
