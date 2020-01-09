package main

import (
	"encoding/json"
	"github.com/wedojava/MyTools"
	"io/ioutil"
	"os"
	"path/filepath"
)

type File struct {
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
	ModTime  int64  `json:"mod_time"`
}

var Files []File

func GetFileList() []File {
	err := filepath.Walk(filepath.Join("..", "downloadSvr", SubFolder), visit)
	MyTools.Check(err)
	return Files
}

func visit(p string, info os.FileInfo, err error) error {
	MyTools.Check(err)
	if !info.IsDir() {
		//loc := time.FixedZone("UTC+8", +8*60*60)
		//t := info.ModTime().In(loc)
		//t.Format(time.RFC1123Z)
		f := File{
			Filename: p,
			Size:     info.Size(),
			ModTime:  info.ModTime().Unix(),
		}
		Files = append(Files, f)
	}
	return nil
}

// SaveFileLstInfo is the function to save uploaded file list information.
// It's write action truncates the file before writing.
func SaveFileLstInfo(files []File, dbFilename string) {
	if len(files) > 0 {
		b, err := json.Marshal(files)
		MyTools.Check(err)
		b = []byte(MyTools.AESEncrypt(string(b), "12345678901234567890123456789012"))
		err = ioutil.WriteFile(dbFilename, b, os.ModePerm)
		MyTools.CheckPanic(err)
		//if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		//	f, err := os.Create(dbFile)
		//	MyTools.CheckPanic(err)
		//	defer f.Close()
		//	_, err = f.Write(b)
		//	MyTools.CheckPanic(err)
		//}else{
		//	f, err := os.Open(dbFile)
		//	MyTools.CheckPanic(err)
		//	defer f.Close()
		//	_, err = f.Write(b)
		//	MyTools.CheckPanic(err)
		//}
	}
}
