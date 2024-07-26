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

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-frameWidth/2, -frameHeight/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)

	screen.DrawImage(PlayerSprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
