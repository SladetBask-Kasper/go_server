package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"strings"
)


const SITE_DOES_NOT_EXIST string = "std/404.html"
const PORT = ":8080" // Should be 8080 on test
const STD_SITE string = "templates/index.html"
const STOCK_PAGE string = "templates/"
const IMG_PATH string = "pics" // NOTE : url should already include slash

func onLoad() {
	fmt.Println("Server is now running on "+string(PORT))
}

func main() {
	onLoad()
	http.HandleFunc("/", serve)
	http.ListenAndServe(PORT, nil)
}

func serve(w http.ResponseWriter, r *http.Request) {
	path := SITE_DOES_NOT_EXIST
	url := r.URL.Path[1:]
	if url == "" {
		path = STD_SITE
	} else if strings.Contains(url, ".html"){
		path = STOCK_PAGE + url
	} else if strings.Contains(url, "<img>") { 
		path = strings.ReplaceAll(url, "<img>", IMG_PATH)
	} else {
		path = STOCK_PAGE + url + ".html"
	}

	

	if _, err := os.Stat(path); err == nil {
		// path exists
		fmt.Println("File Exists. Sending file...")
		body, _ := ioutil.ReadFile(path)
		//fmt.Println(string(body))
		//fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		fmt.Fprintf(w, string(body))
	} else if os.IsNotExist(err) {
		// path does *not* exist
		fmt.Println("File doesn't exist.")
		body, _ := ioutil.ReadFile(SITE_DOES_NOT_EXIST)
		fmt.Fprintf(w, string(body))
	}
}

// DoH DNS over HTTPS
