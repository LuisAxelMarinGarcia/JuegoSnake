package models

import (
	"time"
	"math/rand"
	_ "image/png"
	"github.com/faiface/pixel/pixelgl"
	"sync"
	"myproject/views"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel"

)

type Position struct {
	x, y int
}

var Snake = []Position{{x: views.NumCellsX / 2, y: views.NumCellsY / 2}}
var Direction = Position{x: 1, y: 0}
var Food = NewFood()
var Mu sync.Mutex
var Score int

func HandleInput(win *pixelgl.Window) {
	for !win.Closed() {
		if win.JustPressed(pixelgl.KeyUp) && Direction.y == 0 {
			Direction = Position{x: 0, y: 1}
		}
		if win.JustPressed(pixelgl.KeyDown) && Direction.y == 0 {
			Direction = Position{x: 0, y: -1}
		}
		if win.JustPressed(pixelgl.KeyLeft) && Direction.x == 0 {
			Direction = Position{x: -1, y: 0}
		}
		if win.JustPressed(pixelgl.KeyRight) && Direction.x == 0 {
			Direction = Position{x: 1, y: 0}
		}
		time.Sleep(time.Millisecond * 10)
	}
}

func DrawRectangle(win *pixelgl.Window, pos Position, color pixel.RGBA) {
	imd := imdraw.New(nil)
	imd.Color = color
	imd.Push(pixel.V(float64(pos.x*views.CellSize), float64(pos.y*views.CellSize)))
	imd.Push(pixel.V(float64((pos.x+1)*views.CellSize), float64((pos.y+1)*views.CellSize)))
	imd.Rectangle(0)
	imd.Draw(win)
}

func MoveSnake() {
	for {
		Mu.Lock()
		head := Snake[0]
		next := Position{x: head.x + Direction.x, y: head.y + Direction.y}
		Snake = append([]Position{next}, Snake[:len(Snake)-1]...)

		if Snake[0] == Food {
			Snake = append(Snake, Position{})
			Food = NewFood()
			Score++
		}

		if Snake[0].x < 0 || Snake[0].x >= views.NumCellsX || Snake[0].y < 0 || Snake[0].y >= views.NumCellsY {
			ResetGame()
		}
		for i := 1; i < len(Snake); i++ {
			if Snake[0] == Snake[i] {
				ResetGame()
			}
		}

		Mu.Unlock()
		time.Sleep(time.Second / 5)
	}
}

func ResetGame() {
	Snake = []Position{{x: views.NumCellsX / 2, y: views.NumCellsY / 2}}
	Direction = Position{x: 1, y: 0}
	Food = NewFood()
	Score = 0
}

func DropFood() {
	for {
		Mu.Lock()
		if rand.Float32() < 0.05 {
			Food = NewFood()
		}
		Mu.Unlock()
		time.Sleep(time.Second)
	}
}

func NewFood() Position {
	return Position{x: rand.Intn(views.NumCellsX), y: rand.Intn(views.NumCellsY)}
}

