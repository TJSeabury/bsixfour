package main

import (
	"bytes"
	"encoding/base64"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
)

func main() {

	a := app.New()
	w := a.NewWindow("bsixfour")
	w.Resize(fyne.NewSize(1024, 768))

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	b64Text := canvas.NewText("", green)
	clock := canvas.NewText("", green)

	content := container.New(
		layout.NewVBoxLayout(),
		container.New(layout.NewCenterLayout(), clock),
		container.New(layout.NewPaddedLayout(), b64Text),
	)

	var imgBytes []byte
	var ext string
	var err error
	var b64Data string
	d := dialog.NewFileOpen(func(f fyne.URIReadCloser, er error) {
		if er != nil {
			err = er
			return
		}
		ext = f.URI().Extension()
		ext = strings.Replace(ext, ".", "", -1)
		if !(ext == "jpg" || ext == "jpeg" || ext == "png") {
			err = er
			return
		}
		imgBytes, er = os.ReadFile(f.URI().Path())
		err = er
		if err != nil {
			panic(err)
		}
		b64Data, er = imgToBase64(imgBytes, ext)
		updateText(b64Text, b64Data)
	}, w)
	d.Show()

	updateTime(clock)
	w.SetContent(content)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
			//updateText(b64Text, b64Data)
		}
	}()

	w.Show()

	a.Run()
}

func updateTime(clock *canvas.Text) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.Text = formatted
	clock.Refresh()
}

func updateText(t *canvas.Text, b string) {
	t.Text = b
	t.Refresh()
}

func imgToBase64(b []byte, ext string) (string, error) {
	var err error = nil
	switch ext {
	case "png":
	case "jpeg":
	case "jpg":
		img, err := jpeg.Decode(bytes.NewReader(b))
		if err != nil {
			return "", err
		}
		var buf bytes.Buffer
		if err := png.Encode(&buf, img); err != nil {
			return "", err
		}
		b = buf.Bytes()
	default:
		return "", err
	}
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(b), err
}
