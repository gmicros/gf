package main

import "os"
import "fmt"
import "time"
import "image"
import "image/gif"
import "image/color"
import "image/color/palette"

	type pnt struct {
		re float64
		im float64
		level int
	}

func doSomething(i int, j int, cmplPl [][]pnt, x0 float64, y0 float64, k int, img *image.Paletted)(pnt) {
	if(cmplPl[i][j].level == 0 ) {
					//x0 = float64(minX) + float64(i)*float64(maxX - minX)/float64(width-1)
					//y0 = float64(maxY) - float64(j)*float64(maxY - minY)/float64(height-1)
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


func main() {
	fmt.Println("Generating fractal gif")
	//var width, height, iter int = 600, 500, 20
	//var minX, maxX, minY, maxY = -2, 1, -1, 1
	
	var width, height, iter int = 600, 600, 20
	var minX, maxX, minY, maxY = -2, -1, 0, 1

	cmplPl := make([][]pnt, width)	
	for i := range cmplPl {
		cmplPl[i] = make([]pnt, height)	
	}
		

	levels := make([][]int, width)
	reVals := make([][]float64, width)
	imVals := make([][]float64, width)
	
	for i := range levels {
		levels[i] = make([]int, height)
		reVals[i] = make([]float64, height)
		imVals[i] = make([]float64, height)
	}


	var images []*image.Paletted
	var delays []int

	var x0, y0 float64

	start := time.Now()
/*
	var temp float64
	for k := 0; k < iter; k++ {
		img := image.NewPaletted(image.Rect(0, 0, width, height), palette.Plan9)
		images = append(images, img)
		delays = append(delays, 80) 
		for i := 0; i < width; i++ {
			x0 = float64(minX) + float64(i)*float64(maxX - minX)/float64(width-1)
			for j := 0; j < height; j++ {
				if(levels[i][j] == 0 ) {
					y0 = float64(maxY) - float64(j)*float64(maxY - minY)/float64(height-1)
					temp = reVals[i][j] * reVals[i][j] - imVals[i][j] * imVals[i][j] + x0
					imVals[i][j] = 2 * reVals[i][j] * imVals[i][j] + y0
					reVals[i][j] = temp 				
					if(reVals[i][j]*reVals[i][j] + imVals[i][j]*imVals[i][j] > 4){
						levels[i][j] = k + 1
						img.Set(i, j, color.RGBA{ uint8(7*levels[i][j]), uint8(7*levels[i][j]), uint8(55*levels[i][j]), 255})
					}
				} else {
					img.Set(i, j, color.RGBA{ uint8(7*levels[i][j]), uint8(7*levels[i][j]), uint8(55*levels[i][j]), 255})

				}
			} 
		}
	}
*/
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	next := time.Now()	
	
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

	elapsed = time.Since(next)
	fmt.Println(elapsed)
	f, _ := os.OpenFile("mandle.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images, 
		Delay: delays,
	})
}