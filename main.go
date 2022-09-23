package main

import (
	"ebiten_balls/balls"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct{}

// Global variables
var (
	black        color.Color = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	Balls        []balls.Ball
	ScreenWidth  = 800
	ScreenHeight = 600
)

func (game *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		Balls = balls.AddBalls(Balls, ScreenWidth, ScreenHeight)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) && len(Balls) > 0 {
		Balls = balls.RemoveBalls(Balls)
	}

	Balls = balls.UpdateBalls(Balls, ScreenWidth, ScreenHeight)

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	screen.Fill(black)

	// mouseX, mouseY := ebiten.CursorPosition()

	for _, ball := range Balls {
		ebitenutil.DrawCircle(screen, float64(ball.Position.X), float64(ball.Position.Y), float64(ball.Radius), ball.Color)
	}
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())

	game := &Game{}

	ebiten.SetWindowPosition(600, 200)
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Balls!")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
