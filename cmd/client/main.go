package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/matheusrb95/desert-survival/internal/game"
)

const (
	screenWidth  = 800
	screenHeight = 600
	gameTitle    = "Desert Survival"
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, gameTitle)
	rl.SetTargetFPS(60)

	game := game.New()

	for !rl.WindowShouldClose() {
		game.Update()
		game.Draw()
	}

	game.Unload()
	rl.CloseWindow()
}
