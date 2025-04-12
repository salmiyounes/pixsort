package pixsort

import (
	"image"
	"image/png"
	"log"
	"os"
)

func LoadImg(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()
	return png.Decode(file)
}

func GetBounds(img image.Image) (int, int) {

	bound := img.Bounds()
	return bound.Dx(), bound.Dy()
}
