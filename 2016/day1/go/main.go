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

func Abs(i int) int {
	if i < 0 {
		i = -i
	}
	return i
}

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case South:
		return "South"
	case East:
		return "East"
	case West:
		return "West"
	}
	return "Unknown"
}

type Position struct {
	x int
	y int
}

type Walker struct {
	Facing           Direction
	Position         Position
	VisitedLocations map[Position]struct{}
}

func (w *Walker) TurnLeft() {
	w.Facing = (w.Facing + 3) % 4
}

func (w *Walker) TurnRight() {
	w.Facing = (w.Facing + 1) % 4
}

func (w *Walker) StepNorth() {
	w.Position.y += 1
	if _, seen := w.VisitedLocations[w.Position]; seen {
		fmt.Println("Already seen", w.Position)
	} else {
		w.VisitedLocations[w.Position] = struct{}{}
	}
}

func (w *Walker) StepEast() {
	w.Position.x += 1
	if _, seen := w.VisitedLocations[w.Position]; seen {
		fmt.Println("Already seen", w.Position)
	} else {
		w.VisitedLocations[w.Position] = struct{}{}
	}
}

func (w *Walker) StepSouth() {
	w.Position.y -= 1
	if _, seen := w.VisitedLocations[w.Position]; seen {
		fmt.Println("Already seen", w.Position)
	} else {
		w.VisitedLocations[w.Position] = struct{}{}
	}
}

func (w *Walker) StepWest() {
	w.Position.x -= 1
	if _, seen := w.VisitedLocations[w.Position]; seen {
		fmt.Println("Already seen", w.Position)
	} else {
		w.VisitedLocations[w.Position] = struct{}{}
	}
}

func (w *Walker) MoveForward(steps int) {
	for i := 0; i < steps; i++ {
		switch w.Facing {
		case North:
			w.StepNorth()
		case East:
			w.StepEast()
		case South:
			w.StepSouth()
		case West:
			w.StepWest()
		}
	}
}

func (w *Walker) HandleMove(move string) {
	cleanMove := strings.ReplaceAll(move, ",", "")
	direction := string(cleanMove[0])
	steps := cleanMove[1:]
	numSteps, err := strconv.Atoi(steps)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("Handling Move '%s'. Rotate: '%s'. Move: '%s'\n", cleanMove, direction, steps)
	switch direction {
	case "L":
		w.TurnLeft()
	case "R":
		w.TurnRight()
	}
	w.MoveForward(numSteps)
}

func (w Walker) BlocksFromOrigin() int {
	return Abs(w.Position.x) + Abs(w.Position.y)
}

func NewWalker() *Walker {
	return &Walker{
		Position:         Position{0, 0},
		Facing:           North,
		VisitedLocations: make(map[Position]struct{}),
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
		w.HandleMove(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Final Position: %d,%d. Distance from Spawn: %d blocks", w.Position.x, w.Position.y, w.BlocksFromOrigin())
}
