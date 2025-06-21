package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func (k *KeyPad) TryMove(dx, dy int) {
	newX := k.currentPosition.x + dx
	newY := k.currentPosition.y + dy

	if newX >= 0 && newX < k.numColumns && newY >= 0 && newY < k.numRows {
		if k.board[newY][newX] > 0 {
			k.currentPosition.x = newX
			k.currentPosition.y = newY
		}
	}
}

func (k *KeyPad) Move(s string) {
	switch s {
	case "U":
		k.TryMove(0, -1)
	case "D":
		k.TryMove(0, 1)
	case "L":
		k.TryMove(-1, 0)
	case "R":
		k.TryMove(1, 0)
	}
}

func (k KeyPad) Combination() string {
	var builder strings.Builder
	for _, d := range k.combination {
		switch d {
		case 10:
			builder.WriteString("A")
		case 11:
			builder.WriteString("B")
		case 12:
			builder.WriteString("C")
		case 13:
			builder.WriteString("D")
		default:
			builder.WriteString(fmt.Sprintf("%d", d))
		}
	}

	return builder.String()
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

	fmt.Printf("Bathroom combination: %s\n", keypad2.Combination())
}
