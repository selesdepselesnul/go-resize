package main
import (
	"github.com/nfnt/resize"
	"log"
	"fmt"
	"path/filepath"
	"os"
	"io"
	"image"
)

type ImageResizer struct {
	file, output string
	width, height uint
	decoder func (io.Reader) (image.Image, error)
	encoder func (io.Writer, image.Image) error
}

func (i ImageResizer) Resize() {
	ext := filepath.Ext(i.file)

	file, err := os.Open(i.file)
	if err != nil {
		log.Fatal(err)
	}
	
	img, err := i.decoder(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(i.width, i.height, img, resize.Lanczos3)

	out, err := os.Create(i.output + ext)
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	i.encoder(out, m)

	fmt.Println(i.output + ext)
}





