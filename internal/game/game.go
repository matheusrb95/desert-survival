package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/matheusrb95/desert-survival/internal/physics"
	"github.com/matheusrb95/desert-survival/pkg/sprite"
	"github.com/matheusrb95/desert-survival/pkg/tilemap"
)

type Textures struct {
	Player rl.Texture2D
	Enemy  rl.Texture2D
	Map    rl.Texture2D
	MapTop rl.Texture2D
}

type Game struct {
	Player    *sprite.Sprite
	Enemy     *sprite.Sprite
	Map       *tilemap.Map
	Textures  Textures
	Camera    rl.Camera2D
	Obstacles []rl.Rectangle
}

func New() (g Game) {
	g.Init()
	return
}

func (g *Game) Init() {
	g.Load()

	g.Player = sprite.New(&g.Textures.Player)
	g.Enemy = sprite.New(&g.Textures.Enemy)
	g.Map = tilemap.New(&g.Textures.Map, &g.Textures.MapTop)

	g.Obstacles = g.Map.Obstacles()
	g.Obstacles = append(g.Obstacles, g.Player.MoveHitbox())
	g.Obstacles = append(g.Obstacles, g.Enemy.MoveHitbox())

	g.Camera = rl.Camera2D{
		Target:   rl.NewVector2(0, 0),
		Offset:   rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenWidth())/2),
		Rotation: 0,
		Zoom:     3,
	}
}

func (g *Game) Load() {
	g.Textures.Player = rl.LoadTexture("assets/HalflingRogue.png")
	g.Textures.Enemy = rl.LoadTexture("assets/EarthElemental.png")
	g.Textures.Map = rl.LoadTexture("assets/ground.png")
	g.Textures.MapTop = rl.LoadTexture("assets/top.png")
}

func (g *Game) Unload() {
	rl.UnloadTexture(g.Textures.Player)
	rl.UnloadTexture(g.Textures.Enemy)
	rl.UnloadTexture(g.Textures.Map)
	rl.UnloadTexture(g.Textures.MapTop)
}

func (g *Game) Update() {
	g.Player.Direction = rl.NewVector2(0, 0)

	if rl.IsKeyDown(rl.KeyD) {
		g.Player.GoRight()
	} else if rl.IsKeyDown(rl.KeyA) {
		g.Player.GoLeft()
	}

	if rl.IsKeyDown(rl.KeyW) {
		g.Player.GoUp()
	} else if rl.IsKeyDown(rl.KeyS) {
		g.Player.GoDown()
	}

	physics.Move(g.Player, g.Obstacles)

	g.Player.Update()
	g.Enemy.Update()

	g.Camera.Target = rl.Vector2AddValue(g.Player.Position, float32(g.Player.SpriteSize)/2)
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode2D(g.Camera)

	g.Map.Draw()
	g.Player.Draw()
	g.Enemy.Draw()
	g.Map.DrawTop()

	g.Map.DrawObstacles()
	//g.Player.DrawMoveHitbox()
	//g.Player.DrawDamageHitbox()

	rl.EndMode2D()
	rl.EndDrawing()
}
