package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type img struct {
	file          *os.File
	width, height int
	size          int64
}

func (i *img) Load(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	i.file = file

	stat, err := file.Stat()
	if err != nil {
		return file, err
	}

	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return file, err
	}

	i.size = stat.Size()
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
