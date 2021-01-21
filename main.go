package main

import (
	"fmt"

	. "pathfinder/types"
	maze "pathfinder/maze"
)

func analyzePoint(point Point) {

}


// On commence par le nœud de départ, c'est le nœud courant.
// On regarde tous ses nœuds voisins.
// Si un nœud voisin est un obstacle, on l'oublie.
// Si un nœud voisin est déjà dans la liste fermée, on l'oublie.
// Si un nœud voisin est déjà dans la liste ouverte, on met à jour la liste ouverte si le nœud dans la liste ouverte a une moins bonne qualité (et on n'oublie pas de mettre à jour son parent).
// Sinon, on ajoute le nœud voisin dans la liste ouverte avec comme parent le nœud courant.
// On cherche le meilleur nœud de toute la liste ouverte. Si la liste ouverte est vide, il n'y a pas de solution, fin de l'algorithme.
// On le met dans la liste fermée et on le retire de la liste ouverte.
// On réitère avec ce nœud comme nœud courant jusqu'à ce que le nœud courant soit le nœud de destination.

func solve(maze Maze) {
	var openList map[Point]Node = make(map[Point]Node)
	var closedList map[Point]Node = make(map[Point]Node)
	var p Point = Point{1,2}
	var n Node = Node{1, 3, 4, Point{1,2}}

	var currPoint Point = maze.Start
	var topPoint Point = Point{currPoint.X, currPoint.Y - 1}
	var rightPoint Point = Point{currPoint.X + 1, currPoint.Y}
	var bottomPoint Point = Point{currPoint.X, currPoint.Y - 1}
	var leftPoint Point = Point{currPoint.X - 1, currPoint.Y}

	openList[p] = n
	closedList[p] = n
	fmt.Println(openList, closedList)
	analyzePoint(maze, topPoint)
	analyzePoint(maze, rightPoint)
	analyzePoint(maze, bottomPoint)
	analyzePoint(maze, leftPoint)
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