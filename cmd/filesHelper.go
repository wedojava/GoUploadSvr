package main

import (
	"encoding/json"
	"github.com/wedojava/MyErrCheck"
	"io/ioutil"
	"os"
	"path/filepath"
)

type File struct {
	Filename string `json:"filename"`
	ModTime  int64  `json:"mod_time"`
}

var Files []File

func GetFileList() []File {
	err := filepath.Walk("../"+SubFolder, visit)
	MyErrCheck.Check(err)
	return Files
}

func visit(p string, info os.FileInfo, err error) error {
	MyErrCheck.Check(err)
	if !info.IsDir() {
		//loc := time.FixedZone("UTC+8", +8*60*60)
		//t := info.ModTime().In(loc)
		//t.Format(time.RFC1123Z)
		f := File{
			Filename: p,
			ModTime:  info.ModTime().Unix(),
		}
		Files = append(Files, f)
	}
	return nil
}

// SaveFileLstInfo is the function to save uploaded file list information.
// It's write action truncates the file before writing.
func SaveFileLstInfo(files []File, dbFile string) {
	if len(files) > 0 {
		b, err := json.Marshal(files)
		MyErrCheck.Check(err)
		err = ioutil.WriteFile(dbFile, b, os.ModePerm)
		MyErrCheck.CheckPanic(err)
		//if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		//	f, err := os.Create(dbFile)
		//	MyErrCheck.CheckPanic(err)
		//	defer f.Close()
		//	_, err = f.Write(b)
		//	MyErrCheck.CheckPanic(err)
		//}else{
		//	f, err := os.Open(dbFile)
		//	MyErrCheck.CheckPanic(err)
		//	defer f.Close()
		//	_, err = f.Write(b)
		//	MyErrCheck.CheckPanic(err)
		//}
	}
}
