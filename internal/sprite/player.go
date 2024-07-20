package sprite

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	*Sprite

	Walking bool
	Hit     bool

	Speed  int
	Health float32
}

func NewPlayer(tx *rl.Texture2D) *Player {
	return &Player{
		Sprite: New(tx),
		Speed:  50,
		Health: 1,
	}
}

func (p *Player) Update() {
	p.Walking = p.Direction != rl.NewVector2(0, 0)

	p.FrameCounter++
	if p.FrameCounter >= frameDelay {
		p.FrameCounter = 0

		if p.Hit {
			if p.FrameIndex%2 == 0 {
				p.Color = rl.NewColor(255, 255, 255, 150)
			} else {
				p.Color = rl.White
			}

			if p.FrameIndex == numFrames-1 {
				p.Hit = false
			}
		}

		p.FrameIndex++
		p.FrameIndex %= numFrames
		p.FrameRec.Y = spriteSize

		p.FrameRec.X = float32(spriteSize * p.FrameIndex)
	}
}

func (p *Player) MoveHitbox() rl.Rectangle {
	return rl.NewRectangle(p.Position.X+2, p.Position.Y+12, 8, 4)
}

func (p *Player) DamageHitbox() rl.Rectangle {
	return rl.NewRectangle(p.Position.X+2, p.Position.Y+4, 8, 12)
}

func (p *Player) Draw() {
	rl.DrawTextureRec(*p.Texture, p.FrameRec, p.Position, p.Color)
}
