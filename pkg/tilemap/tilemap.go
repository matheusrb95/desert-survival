package tilemap

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	tileSize = 16
)

type Map struct {
	texture    *rl.Texture2D
	textureTop *rl.Texture2D
	frameRec   rl.Rectangle
}

func New(tx *rl.Texture2D, txTop *rl.Texture2D) *Map {

	return &Map{
		texture:    tx,
		textureTop: txTop,
		frameRec:   rl.NewRectangle(0, 0, float32(tx.Width), float32(tx.Height)),
	}
}

func (m *Map) Draw() {
	rl.DrawTexture(*m.texture, 0, 0, rl.White)
	rl.DrawTextureRec(*m.texture, m.frameRec, rl.NewVector2(0, 0), rl.White)
}

func (m *Map) DrawTop() {
	rl.DrawTextureRec(*m.textureTop, m.frameRec, rl.NewVector2(0, 0), rl.White)
}

func (m *Map) Obstacles() []rl.Rectangle {
	var result []rl.Rectangle

	result = append(result, m.borderCollision()...)
	result = append(result, m.objectCollision()...)

	return result
}

func (m *Map) DrawObstacles() {
	for _, rec := range m.Obstacles() {
		rl.DrawRectangleLinesEx(rec, 1, rl.Red)
	}
}

func (m *Map) borderCollision() []rl.Rectangle {
	var result []rl.Rectangle

	leftBorder := rl.NewRectangle(0, 0, tileSize, float32(m.texture.Height))
	rightBorder := rl.NewRectangle(float32(m.texture.Width)-tileSize, 0, tileSize, float32(m.texture.Height))

	result = append(result, leftBorder)
	result = append(result, rightBorder)

	return result
}

func (m *Map) objectCollision() []rl.Rectangle {
	var result []rl.Rectangle

	tree := rl.NewRectangle(tileSize*11+2, tileSize*5+6, 10, 10)
	obj := rl.NewRectangle(tileSize*3, tileSize*4, 10, 10)

	result = append(result, tree)
	result = append(result, obj)

	return result
}
