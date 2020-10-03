package main

import (
	"fmt"
	"os"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var textureLookup = map[string]*rl.Texture2D{}

var defaultImage *rl.Image
var defaultTexture rl.Texture2D

func TextureInit() {
	defaultImage = rl.GenImageColor(100, 100, rl.Magenta)
	defaultTexture = rl.LoadTextureFromImage(defaultImage)
}

func TextureGet(png string) *rl.Texture2D {
	// ref if we already loaded it

	if val, ok := textureLookup[png]; ok {
		return val
	}

	// png doesn't exist
	if _, err := os.Stat(png); os.IsNotExist(err) {
		fmt.Printf("Warning: TextureGet: Could not load %s. Utilizing default\n",
			png)
		return &defaultTexture
	}

	// all is well
	texture := rl.LoadTexture(png)
	textureLookup[png] = &texture
	return &texture
}

func TextureCleanup() {
	for key, _ := range textureLookup {
		rl.UnloadTexture(*textureLookup[key])
	}

	rl.UnloadTexture(defaultTexture)
	rl.UnloadImage(defaultImage)
}
