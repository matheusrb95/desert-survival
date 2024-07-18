package sprite

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	spriteSize = 16
	frameDelay = 8
	numFrames  = 4
)

type Sprite struct {
	Position   rl.Vector2
	Direction  rl.Vector2
	SpriteSize int

	Walking   bool
	Attacking bool
	Hit       bool

	Speed  int
	Health float32

	texture      *rl.Texture2D
	frameRec     rl.Rectangle
	frameCounter int
	frameIndex   int
	color        rl.Color
}

func New(tx *rl.Texture2D) *Sprite {
	return &Sprite{
		texture:    tx,
		Position:   rl.NewVector2(100, 100),
		frameRec:   rl.NewRectangle(0, 0, float32(spriteSize), float32(spriteSize)),
		Speed:      50,
		Health:     1,
		color:      rl.White,
		SpriteSize: 16,
	}
}

func (s *Sprite) Update() {
	s.Walking = s.Direction != rl.NewVector2(0, 0)

	s.frameCounter++
	if s.frameCounter >= frameDelay {
		s.frameCounter = 0

		if s.Hit {
			if s.frameIndex%2 == 0 {
				s.color = rl.NewColor(255, 255, 255, 150)
			} else {
				s.color = rl.White
			}

			if s.frameIndex == 5 {
				s.Hit = false
			}
		}

		s.frameIndex++
		s.frameIndex %= numFrames
		s.frameRec.Y = spriteSize

		s.frameRec.X = float32(spriteSize * s.frameIndex)
	}
}

func (s *Sprite) GoUp() {
	s.Direction.Y = -1
}

func (s *Sprite) GoDown() {
	s.Direction.Y = 1
}

func (s *Sprite) GoLeft() {
	if s.frameRec.Width > 0 {
		s.frameRec.Width *= -1
	}

	s.Direction.X = -1
}

func (s *Sprite) GoRight() {
	if s.frameRec.Width < 0 {
		s.frameRec.Width *= -1
	}

	s.Direction.X = 1
}

func (s *Sprite) MoveHitbox() rl.Rectangle {
	return rl.NewRectangle(s.Position.X+2, s.Position.Y+12, 8, 4)
}

func (s *Sprite) DamageHitbox() rl.Rectangle {
	return rl.NewRectangle(s.Position.X+2, s.Position.Y+4, 8, 12)
}

func (s *Sprite) Draw() {
	rl.DrawTextureRec(*s.texture, s.frameRec, s.Position, s.color)
}

func (s *Sprite) DrawMoveHitbox() {
	rl.DrawRectangleLinesEx(s.MoveHitbox(), 1, rl.Green)
}

func (s *Sprite) DrawDamageHitbox() {
	rl.DrawRectangleLinesEx(s.DamageHitbox(), 1, rl.Blue)
}
