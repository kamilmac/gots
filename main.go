package main

import (
	"fmt"
	// "github.com/boltdb/bolt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"path"
    "flag"
    "os/exec"
	// "encoding/json"
	// "io/ioutil"
)

const (
	gDriveURL string = "https://33af57eeb9498ac053d3e355288e591ca01d5a7b.googledrive.com/host/0B-f7h8x-3DuZfktDNmtfMzlEN0hqdk1ORzJzRTNWQXlsNndmVEZTVGpjWUJkb1FGRDJGcE0/"
)

var (
    prod bool
    port int
)

func init() {
    flag.BoolVar(&prod, "prod", false, "Set to True to run in development mode.")
    flag.IntVar(&port, "port", 3000, "Specify port.")
    flag.Parse()
    if prod == false {
        fmt.Println("Running in development mode")
        gulpCSSout, _ := exec.Command("gulp", "--gulpfile", "client/gulpfile.js", "css:bundle").Output()
        fmt.Printf("%s", gulpCSSout)
        fmt.Println("Css bundling finished")        
        go func() {
            fmt.Println("Watching *.ts and *.css files for changes")        
            exec.Command("gulp", "--gulpfile", "client/gulpfile.js", "watch").Output()
        }()
    } 
    // else  {
    //     gulpJSout, _ := exec.Command("gulp", "--gulpfile", "client/gulpfile.js", "js:bundle").Output()
    //     fmt.Printf("%s", gulpJSout)
    //     gulpCSSout, _ := exec.Command("gulp", "--gulpfile", "client/gulpfile.js", "css:bundle").Output()
    //     fmt.Printf("%s", gulpCSSout)
    // }
    fmt.Println("Running on port:", port)
}

func main() {
	router := httprouter.New()
	router.ServeFiles("/vendor/*filepath", http.Dir("client/node_modules"))
	router.ServeFiles("/static/*filepath", http.Dir("client/static"))
	router.GET("/", indexHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := map[string]interface{}{
		"gdriveurl": gDriveURL,
	}
	serveTemplate(w, "index.html", data)
}

func serveTemplate(w http.ResponseWriter, file string, data map[string]interface{}) {
	data["prod"] = prod
    
    layout := path.Join("templates", "layout.html")
	body := path.Join("templates", file)

	tmpl, _ := template.ParseFiles(layout, body)
	tmpl.ExecuteTemplate(w, "layout", data)
}
