package types

type Point struct {
	X int
	Y int
}

type NodeInt struct {
	CostStart int
	CostEnd int
	Cost int
	Parent Point
}

type NodeFloat struct {
	CostStart float64 
	CostEnd float64
	Cost float64
	Parent Point
}

type Maze struct {
	Maze []string
	Start Point
	End Point
}

type ListInt map[Point]*NodeInt
type ListFloat map[Point]*NodeFloat