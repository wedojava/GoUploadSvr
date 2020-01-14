package main

import (
	"flag"
	theSvr "github.com/wedojava/go_upload_srv/cmd/server"
)

func main() {
	uploadPort := flag.String("p1", "8100", "port to upload serve on")
	uploadAction := flag.String("a1", "upload", "default is upload, so use http://yourwebsite/upload for upload action.")
	downloadPort := flag.String("p2", "8200", "port to download serve on")
	downloadAction := flag.String("a2", "files", "default is upload, so use http://localhost:8100/files for display or download files action.")
	downloadPath := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()
	go theSvr.UploadSrvStart(*uploadPort, *uploadAction)
	go theSvr.DownloadSrvStart(*downloadPort, *downloadPath, *downloadAction)
	select {}
}
