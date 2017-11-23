package main
import (
	"image/png"
	"image/jpeg"
	"image"
	"path/filepath"
	"github.com/urfave/cli"
	"io"
	"io/ioutil"
	"os"
	"log"
	"strings"
)


func resizePng(file, output string, width, height uint) {
	ImageResizer {
		file,
		output,
		width,
		height,
		func (x io.Reader) (image.Image, error) {
			return png.Decode(x)
		},
		func (w io.Writer, i image.Image) error {
			return png.Encode(w, i)
		},
	}.Resize()
}

func resizeJpg(file, output string, width, height uint) {
	ImageResizer {
		file,
		output,
		width,
		height,
		func (x io.Reader) (image.Image, error) {
			return jpeg.Decode(x)
		},
		func (w io.Writer, i image.Image) error {
			return jpeg.Encode(w, i, nil)
		},
	}.Resize()
}

func resizeImagesInDir(sourceArg, destArg string, width, height uint) {
	files, err := ioutil.ReadDir(sourceArg)

	os.Mkdir(destArg, os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()
		extractedExtension := filepath.Ext(fileName)
		fileNameNoExt := strings.Replace(fileName, extractedExtension, "", -1)
				
		sourceFName := sourceArg + "/" + fileName
		destFName := destArg + "/" + fileNameNoExt + "_resize"

		if extractedExtension == ".png" {
			resizePng(sourceFName, destFName, width, height)
		} else if extractedExtension == ".jpg" {
			resizeJpg(sourceFName, destFName, width, height)
		}
	}
}

func resizeImagesInZip(fileArg, outputArg string, width, height uint) {
	unzip(fileArg, outputArg)
	resizeImagesInDir(outputArg, outputArg, width, height)
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
			resizePng(fileArg, outputArg, widthUint, heightUint)
		} else if ext == ".jpg" {
			resizeJpg(fileArg, outputArg, widthUint, heightUint)
		} else if ext == ".zip" {
			resizeImagesInZip(fileArg, outputArg, widthUint, heightUint)
		} else {
			if isDir, _ := isDirectory(fileArg); isDir {
				resizeImagesInDir(fileArg, outputArg, widthUint, heightUint)
			} 
		}
	} 	
}

