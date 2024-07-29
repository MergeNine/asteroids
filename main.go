package main

import (
	"bytes"
	"embed"
	"fmt"
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
		if g.playerPosition.Y > 0 {
			g.playerPosition.Y -= speedY
		}
		fmt.Println(g.playerPosition.Y, g.playerPosition.X)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if g.playerPosition.Y < screenHeight-frameHeight {
			g.playerPosition.Y += speedY
		}
		fmt.Println(g.playerPosition.Y, g.playerPosition.X)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if g.playerPosition.X > 0 {
			g.playerPosition.X -= speedX
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if g.playerPosition.X < screenWidth-frameWidth {
			g.playerPosition.X += speedX
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.playerPosition.X, g.playerPosition.Y)

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
