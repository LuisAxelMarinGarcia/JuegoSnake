package views

import (
	"image"
	_ "image/png"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"os"
	
)

type GameState int

const (
	StartMenu GameState = iota
	Playing
)


var StartScreenPic pixel.Picture
var StartButtonSprite *pixel.Sprite
var StartButtonPic pixel.Picture
var ButtonScaleFactor float64 = 0.4
var ButtonPulseDirection float64 = 0.0009
var CurrentState GameState = StartMenu

const (
	Width     = 700
	Height    = 600
	CellSize  = 20
	NumCellsX = Width / CellSize
	NumCellsY = Height / CellSize
)

func LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func LoadAssets() {
	var err error
	StartScreenPic, err = LoadPicture("assets/main.png")
	if err != nil {
		panic(err)
	}
	StartButtonPic, err = LoadPicture("assets/start_button.png")
	if err != nil {
		panic(err)
	}
	StartButtonSprite = pixel.NewSprite(StartButtonPic, StartButtonPic.Bounds())
}

func DrawStartScreen(win *pixelgl.Window) {
	win.Clear(pixel.RGB(0, 0, 0))
	startScreen := pixel.NewSprite(StartScreenPic, StartScreenPic.Bounds())
	startScreen.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	ButtonScaleFactor += ButtonPulseDirection
	if ButtonScaleFactor >= 0.5 || ButtonScaleFactor <= 0.4 {
		ButtonPulseDirection = -ButtonPulseDirection
	}

	StartButtonSprite.Draw(win, pixel.IM.Scaled(pixel.ZV, ButtonScaleFactor).Moved(win.Bounds().Center().Sub(pixel.V(0, 150))))
}

func HandleStartScreenInput(win *pixelgl.Window) {
	if win.JustPressed(pixelgl.MouseButtonLeft) {
		if StartButtonSprite.Frame().Contains(win.MousePosition()) {
			CurrentState = Playing
		}
	}
}


