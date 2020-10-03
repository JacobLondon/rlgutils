package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SceneInitFunc func(scene *Scene)

type Scene struct {
	props []*Prop
	init SceneInitFunc
	name string
}
