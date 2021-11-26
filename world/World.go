package world

import (
	"fmt"
)

var World [25][25]Cell

func InitiateWorld() {
	for row := 0; row < len(World); row++ {
		for column := 0; column < len(World[row]); column++ {
			World[row][column] = NewCell(row, column)
		}
	}
}

func DisplayWorld() {
	for row := 0; row < len(World); row++ {
		for column := 0; column < len(World[row]); column++ {
			if  World[row][column].GetState() {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func NextGenWorld()  {
	fmt.Println("Next Gen")
	for row := 1; row < len(World)-1; row++ {
		for column := 1; column < len(World[row])-1; column++ {
			World[row][column].NextGen()
		}
	}
}

