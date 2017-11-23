package main
import (
	"image/png"
	"image/jpeg"
	"image"
	"path/filepath"
	"github.com/urfave/cli"
	"strconv"
	"io"
	"os"
)

func toUint(strInt string) uint {
	parsedUint64, _ := strconv.ParseUint(strInt, 10, 64)
	return uint(parsedUint64)
}

func main() {

	var fileArg, outputArg, widthArg, heightArg string

	app := cli.NewApp()
	app.Name = "go-resize"
	app.Usage = "resize any image file"

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
			}.Resize()
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
			}.Resize()
		}
	}
}

