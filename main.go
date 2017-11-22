package  main
import (
	"github.com/nfnt/resize"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"github.com/urfave/cli"
	"fmt"
	"strconv"
)

func resizeImg(fileArg, outputArg string, width, height uint) {
	ext := filepath.Ext(fileArg)

	file, err := os.Open(fileArg)
	if err != nil {
		log.Fatal(err)
	}
	
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(width, height, img, resize.Lanczos3)

	out, err := os.Create(outputArg + ext)
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	png.Encode(out, m)

	fmt.Println(outputArg + ext)
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
		resizeImg(fileArg, outputArg, toUint(widthArg), toUint(heightArg))
	}
}











