package main

import (
	"os"
	"io"
	"fmt"
	"log"
	"bufio"
	"strings"
	"strconv"
	//"runtime"
	"image/gif"
	"net/http"
	"html/template"
	"io/ioutil"
	"./fractal_gen"
)

var (
	Trace 	*log.Logger
	Info	*log.Logger
	Warning	*log.Logger
	Error	*log.Logger
)

func Init(
    traceHandle io.Writer,
    infoHandle io.Writer,
    warningHandle io.Writer,
    errorHandle io.Writer) {

    Trace = log.New(traceHandle,
        "TRACE: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Info = log.New(infoHandle,
        "INFO: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Warning = log.New(warningHandle,
        "WARNING: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Error = log.New(errorHandle,
        "ERROR: ",
        log.Ldate|log.Ltime|log.Lshortfile)
}

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
	Info.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
    	Info.Println("Parsing template")
		t, _ := template.ParseFiles("gen.gtpl")
		Info.Println("Executing template")
        t.Execute(res, nil)
    } else {
        r.ParseForm()
		Info.Println("Initializing values")
		InitGen(strings.Join(r.Form["minX"], ""), strings.Join(r.Form["maxX"], ""), 
				strings.Join(r.Form["minY"], ""), strings.Join(r.Form["maxY"], ""),
				strings.Join(r.Form["iter"], ""), strings.Join(r.Form["resl"], "")) 
		Info.Println("Generating fractal")
		images, delays := fractal_gen.GenerateFractalGif()
		
		res.Header().Set("Content-Type","image/gif")
		gif.EncodeAll(res, &gif.GIF{
			Image: images, 
			Delay: delays,
		})
	}
}

func main() {	
	//fmt.Println("Generating fractal gif")
	//noProcs := runtime.GOMAXPROCS(500)
	//fmt.Println("Number of workers = ", noProcs)
    Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
    Info.Println("Starting application")
	http.HandleFunc("/", respHandler)
	  
	port := os.Getenv("PORT")
	if port == "" {
        	port = "8080"
    }
    Info.Println("Listening for requests on port = ["+port+"]")
    http.ListenAndServe(":"+port, nil)
}
