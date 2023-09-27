package scenes

import (
	"os"
	"strconv"
	"time"
	_ "image/png"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"myproject/models"
	"myproject/views"
)


func DrawScore(win *pixelgl.Window) {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(10, views.Height-20), atlas)
	txt.Color = pixel.RGB(1, 1, 1)
	txt.WriteString("Puntuacion: " + strconv.Itoa(models.Score))
	txt.Draw(win, pixel.IM)
}

func InitBackgroundMusic(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		panic(err)
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	loopedStreamer := beep.Loop(-1, streamer)
	speaker.Play(loopedStreamer)
}


func Run() {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Snake",
		Bounds: pixel.R(0, 0, views.Width, views.Height),
	})
	if err != nil {
		panic(err)
	}
	backgroundPic, err := views.LoadPicture("assets/background.png")
	if err != nil {
		panic(err)
	}
	backgroundSprite := pixel.NewSprite(backgroundPic, backgroundPic.Bounds())
	InitBackgroundMusic("assets/music.mp3")
	views.LoadAssets()
	go models.MoveSnake()
	go models.DropFood()
	go models.HandleInput(win)
	for !win.Closed() {
		switch views.CurrentState {
		case views.StartMenu:
			views.DrawStartScreen(win)
			views.HandleStartScreenInput(win)
		case views.Playing:
			models.Mu.Lock()
			win.Clear(pixel.RGB(0, 0, 0))
			backgroundSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
			for _, pos := range models.Snake {
				models.DrawRectangle(win, pos, pixel.RGB(0, 1, 0))
			}
			models.DrawRectangle(win, models.Food, pixel.RGB(1, 0, 0))
			DrawScore(win)
			models.Mu.Unlock()
			time.Sleep(time.Millisecond * 100)
		}
		win.Update()
	}
}
