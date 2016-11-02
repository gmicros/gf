package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"runtime"
	"image/gif"
	"net/http"
	// "io"
	"html/template"
	// "io/ioutil"
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

func InitGen(xMin string, xMax string, yMin string, yMax string,
						itert string, rsl string) {
	minX, _ := strconv.Atoi(strings.TrimSpace(xMin))	
	maxX, _ := strconv.Atoi(strings.TrimSpace(xMax))
	minY, _ := strconv.Atoi(strings.TrimSpace(yMin))
	maxY, _ := strconv.Atoi(strings.TrimSpace(yMax))
	iter, _ := strconv.Atoi(strings.TrimSpace(itert))
	resl, _ := strconv.Atoi(strings.TrimSpace(rsl))
	
	fractal_gen.InitVariables(minX, maxX, minY, maxY, iter, resl)
}

func respHandler(res http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("gen.gtpl")
        t.Execute(res, nil)
    } else {
        r.ParseForm()
		
		InitGen(strings.Join(r.Form["minX"], ""), strings.Join(r.Form["maxX"], ""), 
				strings.Join(r.Form["minY"], ""), strings.Join(r.Form["maxY"], ""),
				strings.Join(r.Form["iter"], ""), strings.Join(r.Form["resl"], "")) 

		images, delays := fractal_gen.GenerateFractalGif()
		
		res.Header().Set("Content-Type","image/gif")
		gif.EncodeAll(res, &gif.GIF{
			Image: images, 
			Delay: delays,
		})
	}
}

func main() {	
	fmt.Println("Generating fractal gif")
	noProcs := runtime.GOMAXPROCS(500)
	fmt.Println("Number of workers = ", noProcs)

	//userInputAndInitGen()	
	//images, delays := fractal_gen.GenerateFractalGif()
	
	//f, _ := os.OpenFile("mandle.gif", os.O_WRONLY|os.O_CREATE, 0600)
	
	//gif.EncodeAll(f, &gif.GIF{
	//	Image: images, 
	// 	Delay: delays,
	// })
	// f.Close()
	http.HandleFunc("/", respHandler)

        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }
        http.ListenAndServe(":"+port, nil)


//    http.ListenAndServe(":8086", nil)
}
