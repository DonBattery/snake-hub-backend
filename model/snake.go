package model

// BodyPart ...
type BodyPart struct {
	GameObj
	Color string
	Type  string
}

// Snake ...
type Snake struct {
	Head Head
	Body []BodyPart
}
