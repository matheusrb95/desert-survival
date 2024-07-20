package sprite

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	spriteSize = 16
	frameDelay = 8
	numFrames  = 4
)

type Sprite struct {
	Texture      *rl.Texture2D
	Color        rl.Color
	Position     rl.Vector2
	Direction    rl.Vector2
	FrameRec     rl.Rectangle
	FrameCounter uint
	FrameIndex   uint
}

func New(tx *rl.Texture2D) *Sprite {
	return &Sprite{
		Texture:  tx,
		Color:    rl.White,
		Position: rl.NewVector2(100, 100),
		FrameRec: rl.NewRectangle(0, 0, spriteSize, spriteSize),
	}
}

func (s *Sprite) GoUp() {
	s.Direction.Y = -1
}

func (s *Sprite) GoDown() {
	s.Direction.Y = 1
}

func (s *Sprite) GoLeft() {
	if s.FrameRec.Width > 0 {
		s.FrameRec.Width *= -1
	}

	s.Direction.X = -1
}

func (s *Sprite) GoRight() {
	if s.FrameRec.Width < 0 {
		s.FrameRec.Width *= -1
	}

	s.Direction.X = 1
}
