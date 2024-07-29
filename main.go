package main

import (
	"bytes"
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
)

const (
	screenWidth  = 320
	screenHeight = 240

	frameWidth  = 99
	frameHeight = 75
)

var (
	assets embed.FS
	//go:embed assets/PNG/playerShip1_blue.png
	playerShipData []byte
	PlayerSprite   = loadImage(playerShipData)
)

func loadImage(data []byte) *ebiten.Image {

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func (v *Vector) Add(dx, dy float64) {
	v.X += dx
	v.Y += dy
}

type Game struct {
	playerPosition Vector
}

func (g *Game) Update() error {
	speedX := float64(200 / ebiten.TPS())
	speedY := float64(200 / ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.playerPosition.Y -= speedY
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.playerPosition.Y += speedY
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {

		g.playerPosition.X -= speedX
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {

		g.playerPosition.X += speedX
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(-frameWidth/2, -frameHeight/2)
	//op.GeoM.Translate(screenWidth/2, screenHeight/2)
	op.GeoM.Translate(g.playerPosition.X, g.playerPosition.Y)
	//width := PlayerSprite.Bounds().Dx()
	//height := PlayerSprite.Bounds().Dy()
	//
	//halfW := float64(width / 2)
	//halfH := float64(height / 2)
	//op.GeoM.Translate(-halfW, -halfH)
	//op.GeoM.Rotate(45.0 * math.Pi / 180.0)
	//op.GeoM.Translate(halfW, halfH)

	//ops := &colorm.DrawImageOptions{}
	//cm := colorm.ColorM{}
	//cm.Translate(1.0, 1.0, 1.0, 0.0)
	//colorm.DrawImage(screen, PlayerSprite, cm, ops)
	screen.DrawImage(PlayerSprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {

	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{
		playerPosition: Vector{X: 100, Y: 100},
	}); err != nil {
		log.Fatal(err)
	}
}

type Vector struct {
	X float64
	Y float64
}
