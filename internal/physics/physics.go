package physics

import (
	"github.com/matheusrb95/desert-survival/pkg/sprite"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func CollideWithRects(rect rl.Rectangle, rects []rl.Rectangle) []rl.Rectangle {
	var hitSlice []rl.Rectangle
	for _, rec := range rects {
		if rl.CheckCollisionRecs(rect, rec) {
			hitSlice = append(hitSlice, rec)
		}
	}
	return hitSlice
}

func Move(s *sprite.Sprite, rects []rl.Rectangle) {
	norm := rl.Vector2Scale(rl.Vector2Normalize(s.Direction), rl.GetFrameTime()*float32(s.Speed))

	s.Position.X += norm.X
	hitList := CollideWithRects(s.MoveHitbox(), rects)
	for _, tile := range hitList {
		if norm.X > 0 {
			s.Position.X = tile.X - (s.MoveHitbox().Width + s.MoveHitbox().X - s.Position.X)
		} else if norm.X < 0 {
			s.Position.X = tile.X + tile.Width - (s.MoveHitbox().X - s.Position.X)
		}
	}

	s.Position.Y += norm.Y
	hitList = CollideWithRects(s.MoveHitbox(), rects)
	for _, tile := range hitList {
		if norm.Y > 0 {
			s.Position.Y = tile.Y - (s.MoveHitbox().Height + s.MoveHitbox().Y - s.Position.Y)
		} else if norm.Y < 0 {
			s.Position.Y = tile.Y + tile.Height - (s.MoveHitbox().Y - s.Position.Y)
		}
	}
}
