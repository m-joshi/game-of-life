package main

import (
	"fmt"
	"game-of-life/world"
)

func main() {
	fmt.Println("Game Of Life")

	world.InitiateWorld()

	world.World[1][2].Resurrect()
	world.World[1][3].Resurrect()
	world.World[2][3].Resurrect()
	world.DisplayWorld()

	world.NextGenWorld()
	world.DisplayWorld()
}
