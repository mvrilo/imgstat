package main

import (
	"reflect"
	"testing"
)

func TestOpen(t *testing.T) {
	tests := []struct {
		path    string
		want    *img
		wantErr bool
	}{
		{
			"testdata/chromium.jpg",
			&img{
				size:   int64(16478),
				width:  440,
				height: 300,
			},
			false,
		},
		{
			"testdata/debian.png",
			&img{
				size:   int64(25281),
				width:  331,
				height: 350,
			},
			false,
		},
		{
			"missing.png",
			nil,
			true,
		},
	}
	for _, tt := range tests {
		got, err := open(tt.path)
		if (err != nil) != tt.wantErr {
			t.Errorf("open() error = %v, wantErr %v", err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("open() = %v, want %v", got, tt.want)
		}
	}
}

func TestImgString(t *testing.T) {
	tests := []struct {
		rheight  int
		rwidth   int
		rsize    int64
		rverbose bool
		want     string
	}{
		{
			rverbose: false,
			rheight:  147,
			rwidth:   98,
			rsize:    44,
			want: "44B	98x147",
		},
		{
			rverbose: false,
			rheight:  120,
			rwidth:   120,
			rsize:    15392,
			want: "15KB	120x120",
		},
		{
			rverbose: true,
			rheight:  920,
			rwidth:   3120,
			rsize:    1,
			want: `size: 1B
width: 3120px
height: 920px`,
		},
		{
			rverbose: true,
			rheight:  9,
			rwidth:   30,
			rsize:    10000,
			want: `size: 10KB
width: 30px
height: 9px`,
		},
	}
	for _, tt := range tests {
		i := &img{
			height:  tt.rheight,
			width:   tt.rwidth,
			size:    tt.rsize,
			verbose: tt.rverbose,
		}
		if got := i.String(); got != tt.want {
			t.Errorf("img.String() = %v, want %v", got, tt.want)
		}
	}
}
