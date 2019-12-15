package model

// GameObj is the basic game element
type GameObj struct {
	ID string
	X  int
	Y  int
}

// Wall is a piece of wall
type Wall struct {
	GameObj
}

// Head is a head of a snake
type Head struct {
	GameObj
	Color string
	Type  string
}
