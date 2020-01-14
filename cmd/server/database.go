package server

import (
	"encoding/json"
	"github.com/wedojava/mytools"
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
	mytools.Check(err)
	return Files
}

func visit(p string, info os.FileInfo, err error) error {
	mytools.Check(err)
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
		//TODO: save it to csv?
		b, err := json.Marshal(files)
		mytools.Check(err)
		b = []byte(mytools.AESEncrypt(string(b), "12345678901234567890123456789012"))
		err = ioutil.WriteFile(dbFilename, b, os.ModePerm)
		mytools.CheckPanic(err)
	}
}
