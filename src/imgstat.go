package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

type img struct {
	file   *os.File
	config image.Config
}

func (i *img) Load(path string) *img {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	i.file = file

	var config image.Config
	switch i.Ext() {
	default:
		config, err = jpeg.DecodeConfig(file)
	case "gif":
		config, err = gif.DecodeConfig(file)
	case "png":
		config, err = png.DecodeConfig(file)
	}

	if err != nil {
		log.Fatal(err)
	}
	i.config = config

	return i
}

func (i *img) Name() string {
	stat, err := i.file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	return stat.Name()
}

func (i *img) Ext() (ext string) {
	name := i.Name()
	ext = strings.Split(name, ".")[1]
	return ext
}

func (i *img) Size() int64 {
	stat, err := i.file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	return stat.Size()
}

func (i *img) Width() int {
	return i.config.Width
}

func (i *img) Height() int {
	return i.config.Height
}

func (i *img) CloseFile() {
	i.file.Close()
}

func (i *img) Output() {
	var size int64
	var t string

	if i.Size() < 1000 {
		size = i.Size()
		t = "B"
	} else {
		size = i.Size() / 1000
		t = "KB"
	}

	fmt.Printf("width: %dpx\nheight: %dpx\nsize: %d%s\n", i.Width(), i.Height(), size, t)
}

func Run() {
	i := new(img)
	i.Load(os.Args[1])
	i.Output()
	i.CloseFile()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <path/to/image>\n", os.Args[0])
		return
	}
	Run()
}
