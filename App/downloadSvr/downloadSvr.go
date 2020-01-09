/*
Serve is a very simple static file server in go
Usage:
	-p="8100": 	port to serve on
	-d=".":    	the directory of static files to host
	-a="files":	so use http://localhost:8100/files for display or download files action.
Navigating to http://localhost:8100/files will display the index.html or directory
listing file.
*/

package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	action := flag.String("a", "files", "default is upload, so use http://localhost:8100/files for display or download files action.")
	flag.Parse()

	//http.Handle("/", http.FileServer(http.Dir(*directory)))
	fs := http.FileServer(http.Dir(*directory))
	http.Handle("/"+*action+"/", http.StripPrefix("/"+*action+"", fs))
	log.Print("Server started on localhost:" + *port + ", use /upload for uploading files and /files/{fileName} for downloading files.")
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
