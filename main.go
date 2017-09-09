package main

import (
	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()
	var err error

	mw := imagick.NewMagickWand()

	err = mw.ReadImage("icon.png")
	if err != nil {
		panic(err)
	}

	_ = mw.ResizeImage(512, 512, imagick.FILTER_LANCZOS)
	_ = mw.WriteImage("out.png")
}
