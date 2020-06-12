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

func CaptureScreen(screen int, quality int) []byte {
	return Capture(bounds(screen), quality)
}

func CaptureMain(quality int) []byte {
	return CaptureScreen(1, quality)
}

func Capture(rect image.Rectangle, quality int) []byte {
	img, err := screenshot.CaptureRect(rect)
	if err != nil {
		panic(err)
	}
	buff := new(bytes.Buffer)
	_ = jpeg.Encode(buff, img, &jpeg.Options{Quality: quality})
	return buff.Bytes()
}
