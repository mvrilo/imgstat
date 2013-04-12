package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

type img struct {
	file          *os.File
	width, height int
	size          int64
}

func (i *img) Load(path string) (file *os.File, err error) {
	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	i.file = file

	stat, err := file.Stat()
	if err != nil {
		return file, err
	}

	i.size = stat.Size()

	var ext string = strings.Split(stat.Name(), ".")[1]
	var config image.Config

	switch ext {
	default:
		config, err = jpeg.DecodeConfig(file)
	case "gif":
		config, err = gif.DecodeConfig(file)
	case "png":
		config, err = png.DecodeConfig(file)
	}

	if err != nil {
		return file, err
	}

	i.width = config.Width
	i.height = config.Height
	return file, nil
}

func (i *img) Print() {
	var size int64
	var t string

	if i.size < 1000 {
		size = i.size
		t = "B"
	} else {
		size = i.size / 1000
		t = "KB"
	}

	fmt.Printf("width: %dpx\nheight: %dpx\nsize: %d%s\n", i.width, i.height, size, t)
}

func Run() {
	i := new(img)
	_, err := i.Load(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer i.file.Close()
	i.Print()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <path/to/image>\n", os.Args[0])
		return
	}
	Run()
}
