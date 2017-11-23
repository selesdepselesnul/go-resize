package  main
import (
	"github.com/nfnt/resize"
	"image/png"
	"image/jpeg"
	"image"
	"log"
	"os"
	"path/filepath"
	"github.com/urfave/cli"
	"fmt"
	"strconv"
	"io"
)

type ImageResizer struct {
	file, output string
	width, height uint
	decoder func (io.Reader) (image.Image, error)
	encoder func (io.Writer, image.Image) error
}

func (i ImageResizer) resize() {
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

func toUint(strInt string) uint {
	parsedUint64, _ := strconv.ParseUint(strInt, 10, 64)
	return uint(parsedUint64)
}

func main() {

	var fileArg, outputArg, widthArg, heightArg string

	app := cli.NewApp()

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "file, f",
			Usage: "file to resize",
			Destination: &fileArg,
		},
		cli.StringFlag{
			Name: "output, o",
			Usage: "",
			Destination: &outputArg,
		},
		cli.StringFlag{
			Name: "width, w",
			Usage: "",
			Destination: &widthArg,
		},
		cli.StringFlag{
			Name: "height, t",
			Usage: "",
			Destination: &heightArg,
		},
	}

	app.Action = func(c *cli.Context) error {	
		return nil
	}

	app.Run(os.Args)

	if fileArg != "" && outputArg != "" && widthArg != "" && heightArg != "" {
		ext := filepath.Ext(fileArg)
		widthUint := toUint(widthArg)
		heightUint := toUint(heightArg)
		
		if ext == ".png" {
			ImageResizer {
				fileArg,
				outputArg,
				widthUint,
				heightUint,
				func (x io.Reader) (image.Image, error) {
					return png.Decode(x)
				},
				func (w io.Writer, i image.Image) error {
					return png.Encode(w, i)
				},
			}.resize()
		} else if ext == ".jpg" {
			ImageResizer {
				fileArg,
				outputArg,
				widthUint,
				heightUint,
				func (x io.Reader) (image.Image, error) {
					return jpeg.Decode(x)
				},
				func (w io.Writer, i image.Image) error {
					return jpeg.Encode(w, i, nil)
				},
			}.resize()
		}
	}
}

