package converter

import (
	"gopkg.in/gographics/imagick.v3/imagick"
)

type Image struct {
	FileName    string
	ContentType string
}

func (i *Image) Resize() []byte {
	imagick.Initialize()
	defer imagick.Terminate()
	var err error

	mw := imagick.NewMagickWand()

	err = mw.ReadImage(i.FileName)
	if err != nil {
		panic(err)
	}

	_ = mw.ResizeImage(512, 512, imagick.FILTER_LANCZOS)

	return mw.GetImageBlob()
}
