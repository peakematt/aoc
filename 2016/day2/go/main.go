package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	x int
	y int
}

type KeyPad struct {
	board           [][]int
	currentPosition Position
	numRows         int
	numColumns      int
	combination     []int
}

func (k KeyPad) CurrentNumber() int {
	return k.board[k.currentPosition.y][k.currentPosition.x]
}

func (k *KeyPad) Capture() {
	k.combination = append(k.combination, k.CurrentNumber())
}

func (k *KeyPad) MoveUp() {
	k.currentPosition.y -= 1
	if k.currentPosition.y < 0 {
		// We clamp currentPosition.y to 0
		k.currentPosition.y = 0
	}

	if k.CurrentNumber() < 0 {
		k.currentPosition.y += 1
	}

	fmt.Printf("Moved Up. New Position: %d\n", k.CurrentNumber())
}

func (k *KeyPad) MoveDown() {
	k.currentPosition.y += 1
	if k.currentPosition.y > k.numRows-1 {
		// We clamp currentPosition.y to numRows-1 (handling off-by-ones)
		k.currentPosition.y = k.numRows - 1
	}

	if k.CurrentNumber() < 0 {
		k.currentPosition.y -= 1
	}

	fmt.Printf("Moved Down. New Position: %d\n", k.CurrentNumber())
}

func (k *KeyPad) MoveLeft() {
	k.currentPosition.x -= 1
	if k.currentPosition.x < 0 {
		// We clamp currentPosition.x to 0
		k.currentPosition.x = 0
	}

	if k.CurrentNumber() < 0 {
		k.currentPosition.x += 1
	}

	fmt.Printf("Moved Left. New Position: %d\n", k.CurrentNumber())
}

func (k *KeyPad) MoveRight() {
	k.currentPosition.x += 1
	if k.currentPosition.x > k.numColumns-1 {
		// We clamp currentPosition.x to numColumns-1 (handling off-by-ones)
		k.currentPosition.x = k.numColumns - 1
	}

	if k.CurrentNumber() < 0 {
		k.currentPosition.x -= 1
	}

	fmt.Printf("Moved Right. New Position: %d\n", k.CurrentNumber())
}

func (k *KeyPad) Move(s string) {
	switch s {
	case "U":
		k.MoveUp()
	case "D":
		k.MoveDown()
	case "L":
		k.MoveLeft()
	case "R":
		k.MoveRight()
	}
}

func (k KeyPad) Combination() string {
	combo := ""
	for _, d := range k.combination {
		if d < 10 {
			combo += fmt.Sprintf("%d", d)
		} else if d == 10 {
			combo += "A"
		} else if d == 11 {
			combo += "B"
		} else if d == 12 {
			combo += "C"
		} else if d == 13 {
			combo += "D"
		}
	}

	return combo
}

func main() {
	/*keypad := KeyPad{
		currentPosition: Position{x: 1, y: 1},
		numRows:         3,
		numColumns:      3,
		board: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	}*/

	keypad2 := KeyPad{
		numRows:         5,
		numColumns:      5,
		currentPosition: Position{x: 0, y: 2},
		board: [][]int{
			{-1, -1, 1, -1, -1},
			{-1, 2, 3, 4, -1},
			{5, 6, 7, 8, 9},
			{-1, 10, 11, 12, -1},
			{-1, -1, 13, -1, -1},
		},
	}

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		for _, c := range line {
			keypad2.Move(string(c))
		}
		keypad2.Capture()
	}

	fmt.Print("Bathroom combination:")
	fmt.Print(keypad2.Combination())
	fmt.Print("\n")
}
