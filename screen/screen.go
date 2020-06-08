package screen

import (
	"bytes"
	"github.com/kbinani/screenshot"
	"github.com/nfnt/resize"
	"image/jpeg"
)

func Capture() []byte {
	n := screenshot.NumActiveDisplays()
	bounds := screenshot.GetDisplayBounds(n - 1)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	buff := new(bytes.Buffer)
	m := resize.Resize(640, 0, img, resize.Lanczos3)
	_ = jpeg.Encode(buff, m, &jpeg.Options{Quality: 70})
	return buff.Bytes()
}
