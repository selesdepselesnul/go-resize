package main
import (
	"strconv"
	"os"
)

func toUint(strInt string) uint {
	parsedUint64, _ := strconv.ParseUint(strInt, 10, 64)
	return uint(parsedUint64)
}

func isDirectory(path string) (bool, error) {
    fileInfo, err := os.Stat(path)
    return fileInfo.IsDir(), err
}

