package physics

import (
	"github.com/matheusrb95/desert-survival/internal/sprite"

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

func Move(p *sprite.Player, rects []rl.Rectangle) {
	norm := rl.Vector2Scale(rl.Vector2Normalize(p.Direction), rl.GetFrameTime()*float32(p.Speed))

	p.Position.X += norm.X
	hitList := CollideWithRects(p.MoveHitbox(), rects)
	for _, tile := range hitList {
		if norm.X > 0 {
			p.Position.X = tile.X - (p.MoveHitbox().Width + p.MoveHitbox().X - p.Position.X)
		} else if norm.X < 0 {
			p.Position.X = tile.X + tile.Width - (p.MoveHitbox().X - p.Position.X)
		}
	}

	p.Position.Y += norm.Y
	hitList = CollideWithRects(p.MoveHitbox(), rects)
	for _, tile := range hitList {
		if norm.Y > 0 {
			p.Position.Y = tile.Y - (p.MoveHitbox().Height + p.MoveHitbox().Y - p.Position.Y)
		} else if norm.Y < 0 {
			p.Position.Y = tile.Y + tile.Height - (p.MoveHitbox().Y - p.Position.Y)
		}
	}
}

func MoveEnemy(e *sprite.Enemy, rects []rl.Rectangle) {
	norm := rl.Vector2Scale(rl.Vector2Normalize(e.Direction), rl.GetFrameTime()*float32(e.Speed))

	e.Position.X += norm.X
	hitList := CollideWithRects(e.MoveHitbox(), rects)
	for _, tile := range hitList {
		if norm.X > 0 {
			e.Position.X = tile.X - (e.MoveHitbox().Width + e.MoveHitbox().X - e.Position.X)
		} else if norm.X < 0 {
			e.Position.X = tile.X + tile.Width - (e.MoveHitbox().X - e.Position.X)
		}
	}

	e.Position.Y += norm.Y
	hitList = CollideWithRects(e.MoveHitbox(), rects)
	for _, tile := range hitList {
		if norm.Y > 0 {
			e.Position.Y = tile.Y - (e.MoveHitbox().Height + e.MoveHitbox().Y - e.Position.Y)
		} else if norm.Y < 0 {
			e.Position.Y = tile.Y + tile.Height - (e.MoveHitbox().Y - e.Position.Y)
		}
	}
}
