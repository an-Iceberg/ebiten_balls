package vec2d

import (
	"fmt"
	"math"
)

type Vec2D struct {
	X float64
	Y float64
}

func (vector *Vec2D) AddVec(vec Vec2D) {
	vector.X += vec.X
	vector.Y += vec.Y
}

func (vector *Vec2D) AddScalar(scalar float64) {
	vector.X += scalar
	vector.Y += scalar
}

func (vector *Vec2D) SubVec(vec Vec2D) {
	vector.X -= vec.X
	vector.Y -= vec.Y
}

func (vector *Vec2D) SubScalar(scalar float64) {
	vector.X -= scalar
	vector.Y -= scalar
}

func (vector *Vec2D) MulVec(vec Vec2D) {
	vector.X *= vec.X
	vector.Y *= vec.Y
}

func (vector *Vec2D) MulScalar(scalar float64) {
	vector.X *= scalar
	vector.Y *= scalar
}

func (vector *Vec2D) DivVec(vec Vec2D) {
	vector.X /= vec.X
	vector.Y /= vec.Y
}

func (vector *Vec2D) DivScalar(scalar float64) {
	vector.X /= scalar
	vector.Y /= scalar
}

func (vector *Vec2D) IsEqualTo(vec Vec2D) bool {
	return vector.X == vec.X && vector.Y == vec.Y
}

func (vector *Vec2D) Mag() float64 {
	return math.Sqrt(math.Pow(vector.X, 2) + math.Pow(vector.Y, 2))
}

func (vector *Vec2D) Mag2() float64 {
	return math.Pow(vector.X, 2) + math.Pow(vector.Y, 2)
}

func (vector *Vec2D) Norm() Vec2D {
	mag := vector.Mag()
	return Vec2D{X: vector.X * mag, Y: vector.Y * mag}
}

func (vector *Vec2D) Perp() Vec2D {
	return Vec2D{X: -vector.Y, Y: vector.X}
}

func (vector *Vec2D) Floor() Vec2D {
	return Vec2D{X: math.Floor(vector.X), Y: math.Floor(vector.Y)}
}

func (vector *Vec2D) Ceil() Vec2D {
	return Vec2D{X: math.Ceil(vector.X), Y: math.Ceil(vector.Y)}
}

func (vector *Vec2D) Max(vec Vec2D) Vec2D {
	return Vec2D{X: math.Max(vector.X, vec.X), Y: math.Max(vector.Y, vec.Y)}
}

func (vector *Vec2D) Min(vec Vec2D) Vec2D {
	return Vec2D{X: math.Min(vector.X, vec.X), Y: math.Min(vector.Y, vec.Y)}
}

func (vector *Vec2D) Cart() Vec2D {
	return Vec2D{X: math.Cos(vector.Y) * vector.X, Y: math.Sin(vector.Y) * vector.X}
}

func (vector *Vec2D) Polar() Vec2D {
	return Vec2D{X: vector.Mag(), Y: math.Atan2(vector.Y, vector.X)}
}

func (vector *Vec2D) Dot(vec Vec2D) Vec2D {
	return Vec2D{X: vector.X * vec.X, Y: vector.Y * vec.Y}
}

func (vector *Vec2D) Cross(vec Vec2D) Vec2D {
	return Vec2D{X: vector.X * vec.Y, Y: vector.Y * vec.X}
}

func (vector *Vec2D) Print() {
	fmt.Printf("(X: %.2f, Y: %.2f)\n", vector.X, vector.Y)
}

func (vector *Vec2D) ToString() string {
	return fmt.Sprintf("(X: %.2f, Y: %.2f)", vector.X, vector.Y)
}
