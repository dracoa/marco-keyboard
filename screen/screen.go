package screen

import (
	"bytes"
	"github.com/kbinani/screenshot"
	"image"
	"image/jpeg"
)

func bounds(screen int) image.Rectangle {
	if screen < 2 || screen > screenshot.NumActiveDisplays() { // screen 1
		return screenshot.GetDisplayBounds(0)
	}
	return screenshot.GetDisplayBounds(screen - 1)
}

func CaptureScreen(screen int) []byte {
	return Capture(bounds(screen))
}

func CaptureMain() []byte {
	return CaptureScreen(1)
}

func Capture(rect image.Rectangle) []byte {
	img, err := screenshot.CaptureRect(rect)
	if err != nil {
		panic(err)
	}
	buff := new(bytes.Buffer)
	_ = jpeg.Encode(buff, img, &jpeg.Options{Quality: 80})
	return buff.Bytes()
}
