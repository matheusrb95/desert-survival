package projectile

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	frameDelay = 10
	numFrames  = 4
	spriteSize = 32
)

type Projectile struct {
	Pos    rl.Vector2
	Width  int32
	Height int32
	Speed  float32
	Angle  float32
	Radius float32

	texture      *rl.Texture2D
	frameRec     rl.Rectangle
	frameCounter int
	frameIndex   int
}

func New(tx *rl.Texture2D) *Projectile {
	return &Projectile{
		texture:  tx,
		Pos:      rl.Vector2Zero(),
		frameRec: rl.NewRectangle(0, 0, float32(spriteSize), float32(spriteSize)),
		Width:    25,
		Height:   25,
		Speed:    90,
		Angle:    0,
		Radius:   50,
	}
}

func (p *Projectile) Update(center rl.Vector2) {
	p.frameCounter++
	if p.frameCounter >= frameDelay {
		p.frameCounter = 0

		p.frameIndex++
		p.frameIndex %= numFrames
		p.frameRec.Y = spriteSize

		p.frameRec.X = float32(spriteSize * p.frameIndex)
	}

	deltaTime := rl.GetFrameTime()
	p.Angle += p.Speed * deltaTime

	p.Pos.X = center.X + float32(math.Cos(float64(p.Angle)*(math.Pi/180))*float64(p.Radius)) - float32(p.Width)/2
	p.Pos.Y = center.Y + float32(math.Sin(float64(p.Angle)*(math.Pi/180))*float64(p.Radius)) - float32(p.Height)/2
}

func (p *Projectile) Hitbox() rl.Rectangle {
	return rl.NewRectangle(p.Pos.X+spriteSize/4, p.Pos.Y+spriteSize/4, spriteSize/2, spriteSize/2)
}

func (p *Projectile) Draw() {
	rl.DrawTextureRec(*p.texture, p.frameRec, p.Pos, rl.White)
}
