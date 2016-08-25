package main

import "os"
import "fmt"
import "image"
import "image/gif"
import "image/color"
import "image/color/palette"

func main() {
	fmt.Println("Generating fractal gif")
	//var width, height, iter int = 600, 500, 20
	//var minX, maxX, minY, maxY = -2, 1, -1, 1
	
	var width, height, iter int = 600, 600, 20
	var minX, maxX, minY, maxY = -2, -1, 0, 1

	type pnt struct {
		re float64
		im float64
		level int
	}

	cmplPl := make([][]pnt, width)	
	for _, xi := range cmplPl {
		xi = make([]pnt, height)	
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
	f, _ := os.OpenFile("mandle.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images, 
		Delay: delays,
	})
}
