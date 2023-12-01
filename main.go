package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "pong")
	defer rl.CloseWindow()

	rl.SetTargetFPS(120)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("pong", 400, 225, 20, rl.Black)

		rl.EndDrawing()
	}
}
