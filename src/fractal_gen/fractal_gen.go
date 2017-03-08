package fractal_gen

import (
    "fmt"
    "image"
    "image/color"
    "image/color/palette"
)

type pnt struct {
    re float64
	im float64
	level int
}

var minX, maxX, minY, maxY, iter, resl, width, height int 
var cmplPl [][]pnt

func InitVariables(xMin int, xMax int, yMin int, yMax int, noIter int, rsl int) {
    minX = xMin
    maxX = xMax
    minY = yMin
    maxY = yMax
    iter = noIter
    resl = rsl
    width = (maxX - minX) * resl
    height = (maxY - minY) * resl
    cmplPl = make([][]pnt, width)	
	for i := range cmplPl {
		cmplPl[i] = make([]pnt, height)	
	}
    fmt.Println("Initializing values...")
}

func doSomething(i int, j int, cmplPl [][]pnt, x0 float64, y0 float64, k int, img *image.Paletted)(pnt) {
	if(cmplPl[i][j].level == 0 ) {
        temp := cmplPl[i][j].re * cmplPl[i][j].re - cmplPl[i][j].im * cmplPl[i][j].im + x0
		cmplPl[i][j].im = 2 * cmplPl[i][j].re * cmplPl[i][j].im + y0
		cmplPl[i][j].re = temp 				
		if(cmplPl[i][j].re*cmplPl[i][j].re + cmplPl[i][j].im*cmplPl[i][j].im > 4){
			cmplPl[i][j].level = k + 1
			img.Set(i, j, color.RGBA{ uint8(7*cmplPl[i][j].level), uint8(7*cmplPl[i][j].level), uint8(55*cmplPl[i][j].level), 255})
		}
	} else {
		img.Set(i, j, color.RGBA{ uint8(7*cmplPl[i][j].level), uint8(7*cmplPl[i][j].level), uint8(55*cmplPl[i][j].level), 255})
	}
	return cmplPl[i][j]
}



func GenerateFractalGif() ([]*image.Paletted, []int) {
    var images []*image.Paletted
	var delays []int
    var x0, y0 float64
    
    for k := 0; k < iter; k++ {
		img := image.NewPaletted(image.Rect(0, 0, width, height), palette.Plan9)
		images = append(images, img)
		delays = append(delays, 80) 
		for i := 0; i < width; i++ {
			x0 = float64(minX) + float64(i)*float64(maxX - minX)/float64(width-1)
			for j := 0; j < height; j++ {
				y0 = float64(maxY) - float64(j)*float64(maxY - minY)/float64(height-1)
				go func(i int, j int, cmplPl [][]pnt, x0 float64, y0 float64, k int, img *image.Paletted) {
					cmplPl[i][j] = doSomething(i, j, cmplPl, x0, y0, k, img);
				}(i, j, cmplPl, x0, y0, k, img);
				
			} 
		}
	}

    return images, delays
}

