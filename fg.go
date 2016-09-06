package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"runtime"
	"image/gif"
	"./fractal_gen"
)

func userInputAndInitGen() {
	fmt.Print("Enter minX, maxX, minY, maxY, noIterations, resolution : ") 

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)	
	result := strings.Split(text, ",")
	
	if(len(result) != 6) {
		fmt.Println("Usage")
		return
	}
	minX, _ := strconv.Atoi(strings.TrimSpace(result[0]))	
	maxX, _ := strconv.Atoi(strings.TrimSpace(result[1]))
	minY, _ := strconv.Atoi(strings.TrimSpace(result[2]))
	maxY, _ := strconv.Atoi(strings.TrimSpace(result[3]))
	iter, _ := strconv.Atoi(strings.TrimSpace(result[4]))
	resl, _ := strconv.Atoi(strings.TrimSpace(result[5]))
	
	fractal_gen.InitVariables(minX, maxX, minY, maxY, iter, resl)
}

func main() {	
	fmt.Println("Generating fractal gif")
	noProcs := runtime.GOMAXPROCS(500)
	fmt.Println("Number of workers = [%d]", noProcs)

	userInputAndInitGen()	
	images, delays := fractal_gen.GenerateFractalGif()
	
	f, _ := os.OpenFile("mandle.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images, 
		Delay: delays,
	})
}
