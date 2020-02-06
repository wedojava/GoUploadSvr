package server

import (
	"context"
	"fmt"
	"github.com/wedojava/myencrypt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// Where the uploaded files save to.
const SubFolder = "files"

var uploadSrv http.Server
var downloadSrv http.Server

func hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                  // parse form data, default will not do that.
	fmt.Fprintf(w, "Hello World!") // this is an view to client as a response.
}

// upload can used by curl like: `curl -F "file=@./go_test.ipynb" http://127.0.0.1:9090/upload`
// TODO: fix cannot uploading file via `--upload-file` to curl
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // Get the type of request
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		//r.Header.Set("Content-Type", "multipart/form-data")
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println("[-] [r.FormFile(\"file\")] Error: ", err)
		}
		defer file.Close()
		_, _ = fmt.Fprintf(w, "%v", handler.Header)
		//TODO: if upload failure, delete the folder
		saveFolder := path.Join(SubFolder, myencrypt.RandStringBytesMaskImprSrc(6))

		if err = os.MkdirAll(saveFolder, os.ModePerm); err != nil {
			fmt.Println("[-] [os.MkdirAll(saveFolder, os.ModePerm)] Error: ", err)
		}
		f, err := os.OpenFile(filepath.Join(saveFolder, handler.Filename),
								os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("[-] [upload:OpenFile] Error: ", err)
		}
		defer f.Close()
		if _, err = io.Copy(f, file); err != nil {
			fmt.Println("[-] [io.Copy(f, file)e] Error: ", err)
		}
		fmt.Println("Uploading complete.")
	}
}

func byeUpload(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,shutdown the upload server!"))
	if err := uploadSrv.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		log.Printf("HTTP server Shutdown: %v", err)
	}
}

func byeDownload(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,shutdown the download server!"))
	if err := downloadSrv.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		log.Printf("HTTP server Shutdown: %v", err)
	}
}

// TODO:  rar db.json for download
func UploadSrvStart(port, action string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/bye", byeUpload)
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/"+action, upload)
	uploadSrv.Addr = ":" + port
	uploadSrv.Handler = mux
	log.Print("[+] Upload Server started on localhost:" + port + ", use `/" + action + "` for uploading files.")
	if err := uploadSrv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

func DownloadSrvStart(port, path, action string) {
	//http.Handle("/", http.FileServer(http.Dir(*directory)))
	mux := http.NewServeMux()
	mux.HandleFunc("/bye", byeDownload)
	mux.Handle("/"+action+"/", http.StripPrefix("/"+action+"", http.FileServer(http.Dir(path))))
	downloadSrv.Addr = ":" + port
	downloadSrv.Handler = mux
	log.Print("[+] Download Server started on localhost:" + port + ", use /files/{fileName} for downloading files.")
	if err := downloadSrv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

