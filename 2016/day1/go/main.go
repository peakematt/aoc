package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Position struct {
	x int
	y int
}

func (p *Position) Add(other Position) {
	p.x += other.x
	p.y += other.y
}

type Vector struct {
	magnitude int
	direction Direction
}

func (v Vector) EndPosition() Position {
	switch v.direction {
	case North:
		return Position{0, v.magnitude}
	case South:
		return Position{0, -v.magnitude}
	case East:
		return Position{v.magnitude, 0}
	case West:
		return Position{-v.magnitude, 0}
	}
	return Position{x: 0, y: 0}
}

type Walker struct {
	Facing                     Direction
	currentPosition            Position
	visitedPositions           map[Position]bool
	firstRepeatedPosition      Position
	foundFirstRepeatedPosition bool
}

func (w *Walker) TurnLeft() {
	w.Facing = (w.Facing + 3) % 4
}

func (w *Walker) TurnRight() {
	w.Facing = (w.Facing + 1) % 4
}

func (w *Walker) WalkTo(p Position) {
	w.currentPosition.Add(p)

	if w.visitedPositions[w.currentPosition] && !w.foundFirstRepeatedPosition {
		w.firstRepeatedPosition = w.currentPosition
		w.foundFirstRepeatedPosition = true
	} else {
		w.visitedPositions[w.currentPosition] = true
	}
}

func (w *Walker) HandleVector(move string) {
	cleanMove := strings.ReplaceAll(move, ",", "")
	direction := string(cleanMove[0])
	steps := cleanMove[1:]
	numSteps, err := strconv.Atoi(steps)
	if err != nil {
		fmt.Println(err)
	}

	switch direction {
	case "L":
		w.TurnLeft()
	case "R":
		w.TurnRight()
	}

	for range numSteps {
		w.WalkTo(Vector{direction: w.Facing, magnitude: 1}.EndPosition())
	}
}

func NewWalker() *Walker {
	return &Walker{
		Facing:                     North,
		visitedPositions:           make(map[Position]bool),
		foundFirstRepeatedPosition: false,
	}
}

func main() {
	f, err := os.Open("input.txt")
	w := NewWalker()
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		w.HandleVector(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("First Overlapping Location: {%d,%d}\n", w.firstRepeatedPosition.x, w.firstRepeatedPosition.y)
	fmt.Printf("Final Location: {%d,%d}\n", w.currentPosition.x, w.currentPosition.y)
}
