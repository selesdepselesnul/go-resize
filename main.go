package  main
import (
	"github.com/nfnt/resize"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fileArg := os.Args[1]
	ext := filepath.Ext(fileArg)

	fileNameNoExt := strings.Replace(fileArg, ext, "", -1)
		
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

	out, err := os.Create(fileNameNoExt + "_resized.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	png.Encode(out, m)
}
