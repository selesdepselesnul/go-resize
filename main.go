package  main
import (
	"github.com/nfnt/resize"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"github.com/urfave/cli"
	"fmt"
)

func resizeImg(fileArg, outputArg string) {
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

	m := resize.Resize(1000, 0, img, resize.Lanczos3)

	out, err := os.Create(outputArg + ext)
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	png.Encode(out, m)

	fmt.Println(outputArg + ext)
}

func main() {

	var fileArg, outputArg string
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
	}

	app.Action = func(c *cli.Context) error {	
		return nil
	}

	app.Run(os.Args)

	if fileArg != "" && outputArg != "" {
		resizeImg(fileArg, outputArg)
	}
}














