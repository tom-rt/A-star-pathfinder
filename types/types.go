package types

type Point struct {
	X int
	Y int
}

type Node struct {
	CostStart int
	CostEnd int
	CostSum int
	Parent Point
}

type Maze struct {
	Maze []string
	Start Point
	End Point
}