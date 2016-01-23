package main

type Point struct {
	X int
	Y int
}

type Barrel struct {
	Pos Point
}

type GameState struct {
	Level     int
	Score     int
	Bonus     int
	PlayerPos Point
	Barrels   []Barrel
}

type History []GameState
