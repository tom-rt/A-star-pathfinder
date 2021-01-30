package types

type Point struct {
	X int
	Y int
}

type Maze struct {
	Maze []string
	Start Point
	End Point
}