package main

import (
	"flag"
	"fmt"
	"github.com/wedojava/MyTools"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// Where the uploaded files save to.
const SubFolder = "files"

func hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                  // parse form data, default will not do that.
	fmt.Fprintf(w, "Hello World!") // this is an view to client as a response.
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// upload can used by curl like:`curl -F "file=@./go_test.ipynb" http://127.0.0.1:9090/upload`
// TODO: fix cannot uploading file via `--upload-file` to curl
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // Get the type of request
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		//r.Header.Set("Content-Type", "multipart/form-data")
		file, handler, err := r.FormFile("file")
		MyTools.CheckErr(err, "upload:FormFile", "Println")
		defer file.Close()
		_, _ = fmt.Fprintf(w, "%v", handler.Header)
		//TODO: if upload failure, delete the folder
		saveFolder := path.Join(SubFolder, MyTools.RandStringBytesMaskImprSrc(6))
		err = os.MkdirAll(saveFolder, os.ModePerm)
		MyTools.CheckErr(err, "Create folder for save file", "Println")
		f, err := os.OpenFile(filepath.Join(saveFolder, handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
		MyTools.CheckErr(err, "upload:OpenFile", "Println")
		defer f.Close()
		_, err = io.Copy(f, file)
		MyTools.CheckErr(err, "upload:Copy", "Println")
		fmt.Println("Uploading complete.")
	}
}

// TODO:  rar db.json for download
func main() {
	port := flag.String("p", "8100", "port to serve on")
	action := flag.String("a", "upload", "default is upload, so use http://yourwebsite/upload for upload action.")
	flag.Parse()

	http.HandleFunc("/", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/"+*action, upload)

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
