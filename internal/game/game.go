package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/matheusrb95/desert-survival/internal/physics"
	"github.com/matheusrb95/desert-survival/internal/sprite"
	"github.com/matheusrb95/desert-survival/internal/sprite/projectile"
	"github.com/matheusrb95/desert-survival/internal/sprite/tilemap"
)

const spriteSize = 16

type Textures struct {
	Player     rl.Texture2D
	Projectile rl.Texture2D
	Enemy      rl.Texture2D
	Map        rl.Texture2D
	MapTop     rl.Texture2D
}

type Game struct {
	Player     *sprite.Player
	Enemy      *sprite.Enemy
	Projectile *projectile.Projectile
	Map        *tilemap.Map
	Textures   Textures
	Camera     rl.Camera2D
	Obstacles  []rl.Rectangle
}

func New() (g Game) {
	g.Init()
	return
}

func (g *Game) Init() {
	g.Load()

	g.Player = sprite.NewPlayer(&g.Textures.Player)
	g.Projectile = projectile.New(&g.Textures.Projectile)
	g.Enemy = sprite.NewEnemy(&g.Textures.Enemy)
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
	g.Textures.Projectile = rl.LoadTexture("assets/projectile.png")
	g.Textures.Enemy = rl.LoadTexture("assets/EarthElemental.png")
	g.Textures.Map = rl.LoadTexture("assets/ground.png")
	g.Textures.MapTop = rl.LoadTexture("assets/top.png")
}

func (g *Game) Unload() {
	rl.UnloadTexture(g.Textures.Player)
	rl.UnloadTexture(g.Textures.Projectile)
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

	centerX := g.Player.Position.X + float32(spriteSize)/2
	centerY := g.Player.Position.Y + float32(spriteSize)/2
	center := rl.NewVector2(centerX, centerY)
	g.Projectile.Update(center)
	if !g.Enemy.Hit && rl.CheckCollisionRecs(g.Projectile.Hitbox(), g.Enemy.DamageHitbox()) {
		g.Enemy.Health -= 25
		g.Enemy.Hit = true
	}

	if g.Player.Position.X > g.Enemy.Position.X {
		g.Enemy.GoRight()
	}
	if g.Player.Position.X < g.Enemy.Position.X {
		g.Enemy.GoLeft()
	}
	if g.Player.Position.Y > g.Enemy.Position.Y {
		g.Enemy.GoDown()
	}
	if g.Player.Position.Y < g.Enemy.Position.Y {
		g.Enemy.GoUp()
	}

	physics.MoveEnemy(g.Enemy, g.Obstacles)
	g.Enemy.Update()

	g.Camera.Target = rl.Vector2AddValue(g.Player.Position, float32(spriteSize)/2)
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode2D(g.Camera)

	g.Map.Draw()
	g.Player.Draw()
	g.Projectile.Draw()
	g.Enemy.Draw()
	g.Map.DrawTop()

	rl.EndMode2D()
	rl.EndDrawing()
}
