package main

import rl "github.com/gen2brain/raylib-go/raylib"

func createGame(width int32, height int32, title string, fps int32) {
	rl.InitWindow(width, height, title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(fps)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		draw()
		rl.EndDrawing()
	}
}

func draw() {
	rl.ClearBackground(rl.RayWhite)
	rl.DrawText("pong", 400, 225, 20, rl.Black)
}

func main() {
	createGame(800, 450, "pong", 120)
}
