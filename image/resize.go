package image

import (
	"fmt"

	"github.com/h2non/bimg"
)

type ImageManipulator struct {
	Quality int
}

func (i *ImageManipulator) Resize(img []byte, x, y int) ([]byte, error) {
	newImage, err := bimg.NewImage(img).Resize(x, y)
	if err != nil {
		return nil, fmt.Errorf("resize image: %w", err)
	}
	return newImage, nil
}

func (i *ImageManipulator) Convert(img []byte) ([]byte, error) {
	newImage, err := bimg.NewImage(img).Convert(bimg.JPEG)
	if err != nil {
		return nil, fmt.Errorf("convert image image: %w", err)
	}

	return newImage, nil
}

func (i *ImageManipulator) Compress(img []byte) ([]byte, error) {
	newImage, err := bimg.NewImage(img).Process(
		bimg.Options{Quality: i.Quality},
	)
	if err != nil {
		return nil, fmt.Errorf("process image: %w", err)
	}
	return newImage, nil
}
