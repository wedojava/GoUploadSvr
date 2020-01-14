/*
Serve is a very simple static file server in go
Usage:
	If you want to run quickly, without any params, just run `go start.go`, it will work as default!
	By default, the params is below:
	-p1="8100": 	port of upload server
	-p2="8200": 	port of download server
	-a1="upload":	upload action
	-a2="files":	download action
	-d=".":    	the directory of static files to host

The default params will work as these:
`-a1="upload"`: you can upload file via `curl -F "file=@./go_test.ipynb" http://127.0.0.1:8100/upload`
`-a2="files"`: use http://localhost:8200/files for display or download files action.
Navigating to http://localhost:8100/files will display the index.html or directory
listing file.
*/
package main
