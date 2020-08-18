package osu

import (
	"image/color"
	"strconv"
	"strings"
)

func newRGB(s string) color.RGBA {
	var rgb color.RGBA
	for i, c := range strings.Split(s, `,`) {
		f, err := strconv.ParseFloat(c, 64)
		if err != nil {
			f = 0
		}
		switch i {
		case 0:
			rgb.R = uint8(f)
		case 1:
			rgb.G = uint8(f)
		case 2:
			rgb.B = uint8(f)
		}
	}
	rgb.A = 255
	return rgb
}