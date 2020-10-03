package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

/**
 * Given a texture and how it is tiled, * traverse the tiles to animate the
 * single texture by changing to each tile. For example, if an image is has
 * 4x4 tiles, then width = height = 4, update will handle switching.
 */
type Animation struct {
	texture *rl.Texture2D
	tileHeight int32 // subsection height in units
	tileWidth int32  // subsection width in units
	ci int32         // current height index
	cj int32         // current width index
}

func AnimationNew(png string, tileWidth int32, tileHeight int32) *Animation {
	var self Animation

	self.texture = TextureGet(png)
	self.tileHeight = tileHeight
	self.tileWidth = tileWidth
	self.ci = 0
	self.cj = 0

	return &self
}

func (self *Animation) AnimationCopy() *Animation {
	var copy Animation
	copy = *self
	return &copy
}

func (self *Animation) AnimationUpdate() {
	self.cj += 1

	if self.cj >= self.tileWidth {
		self.cj = 0
		self.ci += 1
	}

	if self.ci >= self.tileHeight {
		self.ci = 0
	}
}

func (self *Animation) AnimationDraw(pos rl.Vector2, scale float32, rotation float32) {
	rl.DrawTexturePro(
		*self.texture,
		rl.Rectangle{
			X: float32(self.texture.Width / self.tileWidth * self.cj),
			Y: float32(self.texture.Height / self.tileHeight * self.ci),
			Width: float32(self.texture.Width / self.tileWidth),
			Height: float32(self.texture.Height / self.tileHeight),
		},
		rl.Rectangle{
			X: pos.X,
			Y: pos.Y,
			Width: float32(self.texture.Width / self.tileWidth) * scale,
			Height: float32(self.texture.Height / self.tileHeight) * scale,
		},
		rl.Vector2{X: 0, Y: 0},
		rotation,
		rl.White,
	)
}

func (self *Animation) AnimationReset() {
	self.ci = 0
	self.cj = 0
}

func (self *Animation) AnimationGetWidth() int32 {
	return self.texture.Width / self.tileWidth
}

func (self *Animation) AnimationGetHeight() int32 {
	return self.texture.Height / self.tileHeight
}

func (self *Animation) AnimationGetFrames() int32 {
	return self.tileWidth * self.tileHeight
}
