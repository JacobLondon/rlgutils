package main

import (
	"math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PropMovementFunc func(prop *Prop, amount float32)

type PropMovement struct {
	amount float32
	movement PropMovementFunc
}

type Prop struct {
	pos rl.Vector2
	bobrate float32
	bobdelta float32
	rotation float32
	scale float32
	animation *Animation
	movements []PropMovement
}

func propMovementNew(amount float32, movement PropMovementFunc) PropMovement {
	var self PropMovement

	self.amount = amount
	self.movement = movement

	return self
}

func (self PropMovement) propMovementUpdate(prop *Prop) {
	self.movement(prop, self.amount)
}

func PropNew(png string, tileWidth int32, tileHeight int32) *Prop {
	return PropNewFromAnimation(AnimationNew(png, tileWidth, tileHeight))
}

func PropNewFromAnimation(animation *Animation) *Prop {
	var self Prop

	self.pos = rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: 0}
	self.bobrate = 0.001
	self.bobdelta = 0
	self.rotation = 0
	self.scale = 0
	self.animation = animation
	self.movements = make([]PropMovement, 0, 4)

	return &self
}

func (self *Prop) PropCopy() *Prop {
	var copy Prop

	copy = *self
	copy.animation = self.animation.AnimationCopy()
	return &copy
}

func (self *Prop) PropUpdate() {
	for i := 0; i < len(self.movements); i += 1 {
		self.movements[i].propMovementUpdate(self)
	}

	if self.bobdelta > 2 * math.Pi {
		self.bobdelta = 0
	}
}

func (self *Prop) PropDraw() {
	if !self.PropIsOffscreen() {
		self.animation.AnimationDraw(self.pos, self.scale, self.rotation)
	}
}

func (self *Prop) PropIsOffscreen() bool {
	return self.pos.X + float32(self.animation.AnimationGetWidth())  < 0 ||
	       self.pos.X > float32(rl.GetScreenWidth()) ||
	       self.pos.Y + float32(self.animation.AnimationGetHeight()) < 0 ||
	       self.pos.Y > float32(rl.GetScreenHeight())
}

func (self *Prop) PropAddMovement(movement PropMovementFunc, amount float32) {
	self.movements = append(self.movements, propMovementNew(amount, movement))
}

/*
 * A bunch of pre-loaded PropMovementFuncs
 */

func PropMovementFunc_RotateClockwise(prop *Prop, amount float32) {
	prop.rotation += amount
}

func PropMovementFunc_RotateCounterClockwise(prop *Prop, amount float32) {
	prop.rotation -= amount
}

func PropMovementFunc_Scale(prop *Prop, amount float32) {
	if prop.PropIsOffscreen() {
		prop.scale = 1
	} else {
		prop.scale *= amount
	}
}

func PropMovementFunc_TrackMouseVertical(prop *Prop, amount float32) {
	y := float32(rl.GetMouseY())

	if prop.pos.Y > y {
		prop.pos.Y -= float32(math.Abs(float64(prop.pos.Y - y))) / amount
	} else {
		prop.pos.Y += float32(math.Abs(float64(prop.pos.Y - y))) / amount
	}
}

func PropMovementFunc_TrackMouseHorizontal(prop *Prop, amount float32) {
	x := float32(rl.GetMouseX())

	if prop.pos.X > x {
		prop.pos.X -= float32(math.Abs(float64(prop.pos.X - x))) / amount
	} else {
		prop.pos.X += float32(math.Abs(float64(prop.pos.X - x))) / amount
	}
}

func PropMovementFunc_Left(prop *Prop, amount float32) {
	prop.pos.X -= amount
}

func PropMovementFunc_Right(prop *Prop, amount float32) {
	prop.pos.X += amount
}

func PropMovementFunc_Up(prop *Prop, amount float32) {
	prop.pos.Y -= amount
}

func PropMovementFunc_Down(prop *Prop, amount float32) {
	prop.pos.Y += amount
}

func PropMovementFunc_BobVertical(prop *Prop, amount float32) {
	prop.pos.Y += amount * float32(math.Sin(float64(prop.bobdelta)))
	prop.bobdelta += prop.bobrate
}

func PropMovementFunc_BobHorizontal(prop *Prop, amount float32) {
	prop.pos.X += amount * float32(math.Cos(float64(prop.bobdelta)))
	prop.bobdelta += prop.bobrate
}

func PropMovementFunc_LoopLeft(prop *Prop, amount float32) {
	PropMovementFunc_Left(prop, amount)
	if prop.PropIsOffscreen() {
		prop.pos.X = float32(rl.GetScreenWidth())
	}
}

func PropMovementFunc_LoopRight(prop *Prop, amount float32) {
	PropMovementFunc_Right(prop, amount)
	if prop.PropIsOffscreen() {
		prop.pos.X = 0 - float32(prop.animation.AnimationGetWidth())
	}
}

func PropMovementFunc_LoopUp(prop *Prop, amount float32) {
	PropMovementFunc_Up(prop, amount)
	if prop.PropIsOffscreen() {
		prop.pos.Y = float32(rl.GetScreenHeight())
	}
}

func PropMovementFunc_LoopDown(prop *Prop, amount float32) {
	PropMovementFunc_Down(prop, amount)
	if prop.PropIsOffscreen() {
		prop.pos.Y = 0 - float32(prop.animation.AnimationGetHeight())
	}
}
