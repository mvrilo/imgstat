package main

import (
	"flag"
	"fmt"
	"image"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

const help = `imgstat - utility to know the size, width and height of an image.
Usage:
	%s [-hv] <path/to/image>
`

type img struct {
	height  int
	width   int
	size    int64
	verbose bool
}

func open(path string) (*img, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return nil, err
	}

	return &img{
		size:   stat.Size(),
		width:  config.Width,
		height: config.Height,
	}, nil
}

// String formats the output the image data
func (i *img) String() string {
	size := i.size
	unit := "B"

	if i.size >= 1000 {
		size = i.size / 1000
		unit = "KB"
	}

	output := "%d%s\t%dx%d"
	if i.verbose {
		output = "size: %d%s\nwidth: %dpx\nheight: %dpx"
	}
	return fmt.Sprintf(output, size, unit, i.width, i.height)
}

func main() {
	h := flag.Bool("h", false, "Usage information")
	v := flag.Bool("v", false, "Verbose output")
	flag.Parse()

	if *h || len(flag.Args()) == 0 {
		os.Stderr.WriteString(fmt.Sprintf(help, os.Args[0]))
		os.Exit(1)
		return
	}

	i, err := open(flag.Args()[0])
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
		return
	}

	i.verbose = *v
	fmt.Println(i)
}
