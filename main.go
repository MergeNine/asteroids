package main

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
	"math"
	"time"
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
	keyHoldTime    = make(map[ebiten.Key]time.Time)
	keyState       = make(map[ebiten.Key]bool)
	initGroAmt     = 200.0
	frameRate      = 60.0
	start          = time.Now()
	keyPressTime   time.Time
	keyHeld        bool
	totalDuration  time.Duration
)

type Game struct {
	pPos Vector
	t    Thrust
}

type Vector struct {
	X float64
	Y float64
}

type Thrust struct {
	accel float64
}

func (g *Game) Update() error {
	currentTime := time.Now()
	expGro := initGroAmt * (math.Pow(1+g.t.accel/250, float64(2)))
	expGro = math.Min(expGro, 2000)

	//fmt.Println(expGro)
	speedX := expGro / frameRate
	speedY := expGro / frameRate
	// Calculate elapsed time since the last frame
	elapsedSinceLastFrame := currentTime.Sub(start)
	start = currentTime
	for _, key := range []ebiten.Key{ebiten.KeyUp, ebiten.KeyDown, ebiten.KeyLeft, ebiten.KeyRight} {

		if ebiten.IsKeyPressed(key) {

			if !keyState[key] {
				// Key was just pressed
				keyPressTime = currentTime
				keyState[key] = true
			}

			// Calculate the duration the key has been held this frame
			if keyState[key] {
				totalDuration += elapsedSinceLastFrame
			}
			switch key {

			case ebiten.KeyUp:
				if g.pPos.Y > 0 {
					g.t.accel = float64(totalDuration.Milliseconds())
					g.pPos.Y -= speedY
				} else {
					start = time.Now()
				}
			case ebiten.KeyDown:
				if g.pPos.Y < screenHeight-frameHeight {
					g.t.accel = float64(totalDuration.Milliseconds())
					g.pPos.Y += speedY
				} else {
					start = time.Now()
				}
			case ebiten.KeyLeft:
				if g.pPos.X > 0 {
					g.t.accel = float64(totalDuration.Milliseconds())
					g.pPos.X -= speedX
				} else {
					start = time.Now()
				}
			case ebiten.KeyRight:
				if g.pPos.X < screenWidth-frameWidth {
					g.t.accel = float64(totalDuration.Milliseconds())
					g.pPos.X += speedX
					fmt.Println(totalDuration.Milliseconds())
				} else {
					start = time.Now()
				}
			default:
				fmt.Println("no key pressed")
			}

		} else {
			if keyState[key] {
				// Key was released, calculate duration
				keyState[key] = false
				duration := time.Since(keyPressTime)
				fmt.Printf("Key %v held for: %v\n", key, duration)
				totalDuration = 0 // Reset total duration after key release

			}
		}

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.pPos.X, g.pPos.Y)

	screen.DrawImage(PlayerSprite, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320 * 4, 240 * 4
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{
		pPos: Vector{X: 100, Y: 100},
		t:    Thrust{accel: 1},
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
