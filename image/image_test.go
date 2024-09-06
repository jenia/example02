package image

import (
	"github.com/h2non/bimg"
	"testing"
)

func Test_Given_Image_When_Resize_Then_Dimensions_Are_Correct(t *testing.T) {
	sizeWant := 400
	buffer, err := bimg.Read("test-image.png")
	if err != nil {
		t.Fatalf("read file: %s", err.Error())
	}

	im := &ImageManipulator{}
	newBuffer, err := im.Resize(buffer, sizeWant, sizeWant)
	if err != nil {
		t.Fatalf("rezie image: %s", err.Error())
	}

	size, err := bimg.NewImage(newBuffer).Size()
	if err != nil {
		t.Fatalf("image size: %s", err.Error())
	}

	if size.Width != sizeWant || size.Height != sizeWant {
		t.Errorf("image size want %d but is: x %d and y %d", sizeWant, size.Width, size.Height)
	}
}
