package main

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
	_ "image/png"
	"log"
	"math"
)

const (
	screenWidth  = 320 * 4
	screenHeight = 240 * 4

	frameWidth  = 99
	frameHeight = 75
)

var (
	assets embed.FS
	//go:embed assets/PNG/playerShip1_blue.png
	playerShipData []byte
	PlayerSprite   = loadImage(playerShipData)
)

type Game struct {
	playerPosition Vector
	thrust         Thrust
}

type Vector struct {
	X float64
	Y float64
}

type Thrust struct {
	time      int
	magnitude int
}

func (g *Game) Update() error {

	expGro := float64(g.thrust.magnitude*2) * (math.Pow(float64(1+(g.thrust.magnitude)/100), float64(g.thrust.time)))
	fmt.Println(expGro)
	speedX := expGro
	speedY := float64(200 / ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyUp) {

		if g.playerPosition.Y > 0 {
			g.playerPosition.Y -= speedY
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if g.playerPosition.Y < screenHeight-frameHeight {
			g.playerPosition.Y += speedY
		}
		//fmt.Println(g.playerPosition.Y, g.playerPosition.X)
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		pressDur := inpututil.KeyPressDuration(ebiten.KeyLeft)
		if g.playerPosition.X > 0 {
			g.thrust.time = pressDur / 10
			g.thrust.magnitude = pressDur / 10
			g.playerPosition.X -= speedX
		}

		//if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		//	t.keyHeld()
		//	fmt.Println(t.time, t.magnitude)
		//}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		pressDur := inpututil.KeyPressDuration(ebiten.KeyRight)
		if g.playerPosition.X < screenWidth-frameWidth {
			g.thrust.time = pressDur / 10
			g.thrust.magnitude = pressDur / 10
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
	return 320 * 4, 240 * 4
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{
		playerPosition: Vector{X: 100, Y: 100},
		thrust:         Thrust{time: 1, magnitude: 1},
	}); err != nil {
		log.Fatal(err)
	}
}

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
