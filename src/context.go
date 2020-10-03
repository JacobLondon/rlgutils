package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)



func ContextLoop() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")

	rl.SetTargetFPS(60)

	TextureInit()
	t := TextureGet("missing.png")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawTexture(*t, rl.GetMouseX(), rl.GetMouseY(), rl.White)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}

	TextureCleanup()
	rl.CloseWindow()
}
