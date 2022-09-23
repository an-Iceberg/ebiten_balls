package balls

import (
	"ebiten_balls/utils"
	"ebiten_balls/vec2d"
	"image/color"
	"math"
	"math/rand"
)

type Ball struct {
	Position vec2d.Vec2D
	Velocity vec2d.Vec2D
	Radius   float64
	Color    color.RGBA
}

func AddBalls(balls []Ball, screenWidth, screenHeight int) []Ball {
	for i := 0; i < 10; i++ {
		radius := 10. + rand.Float64()*(40-10)
		speed := rand.Intn(300) + 100
		angle := rand.Intn(628)

		newBall := Ball{
			Position: vec2d.Vec2D{
				// Places the circle within the screen space (regardless of radius)
				X: utils.RandomFloat(radius, float64(screenWidth)-radius),
				Y: utils.RandomFloat(radius, float64(screenHeight)-radius),
			},
			// TODO: random velocity
			Velocity: vec2d.Vec2D{
				X: float64(speed) * math.Cos(float64(angle)/100),
				Y: float64(speed) * math.Sin(float64(angle)/100),
			},
			// Random radius size between 10 and 40
			Radius: radius,
			Color: color.RGBA{
				R: uint8(rand.Intn(255)),
				G: uint8(rand.Intn(255)),
				B: uint8(rand.Intn(255)),
				A: 255,
			},
		}

		balls = append(balls, newBall)
	}

	return balls
}

func RemoveBalls(balls []Ball) []Ball {
	if len(balls) > 9 {
		balls = balls[10:]
	} else {
		balls = nil
	}

	return balls
}

func UpdateBalls(balls []Ball, screenWidth, screenHeight int) []Ball {
	deltaTime := 1. / 60.

	// ! A for range loop only returns copies of the slice elements

	for i := range balls {
		// Updates the position of the balls
		balls[i].Position.X += balls[i].Velocity.X * deltaTime
		balls[i].Position.Y += balls[i].Velocity.Y * deltaTime

		// Keeping the balls inside screen space on the x axsis
		if balls[i].Position.X < float64(balls[i].Radius) || balls[i].Position.X > (float64(screenWidth)-float64(balls[i].Radius)) {
			balls[i].Position.X = utils.Clamp(balls[i].Position.X, float64(balls[i].Radius), float64(screenWidth)-float64(balls[i].Radius))
			balls[i].Velocity.X *= -1
		}

		// Keeping the balls inside screen space on the y axis
		if balls[i].Position.Y < float64(balls[i].Radius) || balls[i].Position.Y > (float64(screenHeight)-float64(balls[i].Radius)) {
			balls[i].Position.Y = utils.Clamp(balls[i].Position.Y, float64(balls[i].Radius), float64(screenHeight)-float64(balls[i].Radius))
			balls[i].Velocity.Y *= -1
		}
	}

	return balls
}
