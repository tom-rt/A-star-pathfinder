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
	cost int
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
		cost = 1
	} else {
		cost = parentNode.cost + 1
	}

	var node node = node {
		cost: cost,
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
	fmt.Println(point)
	_, check = total[point]
	fmt.Println(point, check)
	if !check {
		newNode := createNode(maze, parent, total, point)
		total[point] = newNode
	}
	*newPoints = append(*newPoints, point)
}

func checkNeighbors(maze t.Maze, total list, lastAdded []t.Point, cost int) (list, []t.Point, int) {
	var newPoints []t.Point
	for _, currPoint := range lastAdded {
		top, right, bottom, left := getNeighbors(maze, currPoint)

		// On regarde tous ses nœuds voisins.
		analyzePoint(maze, total, currPoint, &newPoints, top)
		analyzePoint(maze, total, currPoint, &newPoints, right)
		analyzePoint(maze, total, currPoint, &newPoints, bottom)
		analyzePoint(maze, total, currPoint, &newPoints, left)

	}
	fmt.Println("######################################################")
	fmt.Println(newPoints)
	return total, newPoints, cost + 1
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
	var total list = make(list)
	var lastAdded []t.Point
	var cost int
	var start time.Time
	var elapsed time.Duration

	start = time.Now()
	// On commence par le nœud de départ, c'est le nœud courant.
	var currPoint t.Point = maze.Start
	lastAdded = append(lastAdded, currPoint)

	total, lastAdded, cost = checkNeighbors(maze, total, lastAdded, cost)
	for isOver == false {
		// On cherche le meilleur nœud de toute la liste ouverte. Si la liste ouverte est vide, il n'y a pas de solution, fin de l'algorithme.
		if len(lastAdded) == 0 {
			isOver = true
			fmt.Println("Pas de solution.")
			continue
		}

		// bestPoint, bestNode := getBestNode(list)

		// On le met dans la liste fermée et on le retire de la liste ouverte.
		// closedList[bestPoint] = bestNode
		// delete(openList, bestPoint)

		// On réitère avec ce nœud comme nœud courant jusqu'à ce que le nœud courant soit le nœud de destination.
		for _, point := range lastAdded {

			if point.X == maze.End.X && point.Y == maze.End.Y {
				elapsed = time.Since(start)
				// cost = drawPath(maze, bestNode, closedList)
				fmt.Println("DIJKSTRA: Chemin trouvé en " + strconv.Itoa(cost) + " coups et " + elapsed.String())
				// bestPoint.X == maze.End.X && bestPoint.Y == maze.End.Y {
				// isOver = true
				continue
			} else {
				total, lastAdded, cost = checkNeighbors(maze, total, lastAdded, cost)
			}
		}
	}
}