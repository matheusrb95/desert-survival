package sprite

import rl "github.com/gen2brain/raylib-go/raylib"

type Enemy struct {
	*Sprite

	Walking bool
	Hit     bool

	Speed  int
	Health float32
}

func NewEnemy(tx *rl.Texture2D) *Enemy {
	return &Enemy{
		Sprite: New(tx),
		Speed:  20,
		Health: 1,
	}
}

func (e *Enemy) Update() {
	e.Walking = e.Direction != rl.NewVector2(0, 0)

	e.FrameCounter++
	if e.FrameCounter >= frameDelay {
		e.FrameCounter = 0

		if e.Hit {
			if e.FrameIndex%2 == 0 {
				e.Color = rl.NewColor(255, 255, 255, 150)
			} else {
				e.Color = rl.White
			}

			if e.FrameIndex == numFrames-1 {
				e.Hit = false
			}
		}

		e.FrameIndex++
		e.FrameIndex %= numFrames
		e.FrameRec.Y = spriteSize

		e.FrameRec.X = float32(spriteSize * e.FrameIndex)
	}
}

func (e *Enemy) MoveHitbox() rl.Rectangle {
	return rl.NewRectangle(e.Position.X+2, e.Position.Y+12, 8, 4)
}

func (e *Enemy) DamageHitbox() rl.Rectangle {
	return rl.NewRectangle(e.Position.X+2, e.Position.Y+4, 8, 12)
}

func (e *Enemy) Draw() {
	rl.DrawTextureRec(*e.Texture, e.FrameRec, e.Position, e.Color)
}
