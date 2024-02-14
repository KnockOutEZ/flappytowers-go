package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"strconv"
	"time"
)

type Apple struct {
	posX   int32
	posY   int32
	width  int32
	height int32
	Color  rl.Color
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(400)
	rl.InitAudioDevice()
	eat_noise := rl.LoadSound("sound/eat.wav")
	rl.InitWindow(screenWidth, screenHeight, "FlappyTowers")
	rl.SetTargetFPS(120)
	plane_down := rl.LoadImage("assets/plane-down.png")
	plane_up := rl.LoadImage("assets/plane-up.png")
	texture := rl.LoadTextureFromImage(plane_up)
	rand.Seed(time.Now().UnixNano())
	var apple_loc int = rand.Intn(400-2+1) - 2
	Apples := []Apple{}
	current_apple := Apple{screenWidth, int32(apple_loc), 50, 50, rl.Red}
	Apples = append(Apples, current_apple)
	var x_coords int32 = screenWidth/2 - texture.Width/2
	var y_coords int32 = screenHeight/2 - texture.Height/2 - 40
	var score int = 0
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

        rl.DrawTexture(rl.LoadTextureFromImage(rl.LoadImage("assets/background.png")), 0, 0, rl.White)
		rl.DrawTexture(texture, x_coords, y_coords, rl.White)
		rl.DrawText("Current Score: "+strconv.Itoa(score), 0, 0, 30, rl.LightGray)
		rl.ClearBackground(rl.RayWhite)
		if rl.IsKeyDown(rl.KeySpace) {
			texture = rl.LoadTextureFromImage(plane_up)
			y_coords -= 5
		} else {
			texture = rl.LoadTextureFromImage(plane_down)
			y_coords += 5
		}
		for io, current_apple := range Apples {
			rl.DrawTexture(rl.LoadTextureFromImage(rl.LoadImage("assets/tower.png")), current_apple.posX, current_apple.posY, rl.White)
			Apples[io].posX = Apples[io].posX - 5
			if current_apple.posX < 0 {
				Apples[io].posX = 800
				Apples[io].posY = int32(rand.Intn(400-2+1) - 2)
				score--
			}
			if rl.CheckCollisionRecs(rl.NewRectangle(float32(x_coords), float32(y_coords), float32(50), float32(80)), rl.NewRectangle(float32(current_apple.posX), float32(current_apple.posY), float32(current_apple.width), float32(current_apple.height))) {
				Apples[io].posX = 800
				Apples[io].posY = int32(rand.Intn(400-2+1) - 2)
				score++
				rl.PlaySound(eat_noise)
			}
		}
		if y_coords > 400 {
			rl.UnloadTexture(texture)
			Apples = nil
			rl.DrawText("Your final score is: "+strconv.Itoa(score), 30, 40, 30, rl.Red)
		}
		rl.EndDrawing()
		time.Sleep(50000000)
	}
	// rl.StopSoundMulti()
	rl.UnloadSound(eat_noise)
	rl.UnloadTexture(texture)
	rl.CloseWindow()
}
