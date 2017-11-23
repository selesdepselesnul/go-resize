package main
import (
	"image/png"
	"image/jpeg"
	"image"
	"path/filepath"
	"github.com/urfave/cli"
	"strconv"
	"io"
	"io/ioutil"
	"os"
	"archive/zip"
	"log"
)

func toUint(strInt string) uint {
	parsedUint64, _ := strconv.ParseUint(strInt, 10, 64)
	return uint(parsedUint64)
}

func Unzip(src, dest string) error {
    r, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer func() {
        if err := r.Close(); err != nil {
            panic(err)
        }
    }()

    os.MkdirAll(dest, 0755)

    // Closure to address file descriptors issue with all the deferred .Close() methods
    extractAndWriteFile := func(f *zip.File) error {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer func() {
            if err := rc.Close(); err != nil {
                panic(err)
            }
        }()

        path := filepath.Join(dest, f.Name)

        if f.FileInfo().IsDir() {
            os.MkdirAll(path, f.Mode())
        } else {
            os.MkdirAll(filepath.Dir(path), f.Mode())
            f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return err
            }
            defer func() {
                if err := f.Close(); err != nil {
                    panic(err)
                }
            }()

            _, err = io.Copy(f, rc)
            if err != nil {
                return err
            }
        }
        return nil
    }

    for _, f := range r.File {
        err := extractAndWriteFile(f)
        if err != nil {
            return err
        }
    }

    return nil
}


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
			Unzip(fileArg, outputArg)
			files, err := ioutil.ReadDir(outputArg)
			if err != nil {
				log.Fatal(err)
			}

			for _, file := range files {
				fileName := file.Name()
				extractedExtension := filepath.Ext(fileName)
				sourceFName := outputArg + "/" + fileName
				destFName := outputArg + "/" + fileName + "_resize" + extractedExtension
				if extractedExtension == ".png" {
					resizePng(sourceFName, destFName, widthUint, heightUint)
				} else if extractedExtension == ".jpg" {
					resizeJpg(sourceFName, destFName, widthUint, heightUint)
				}
			}
		}
	} 	
}









