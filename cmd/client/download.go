// TODO: test download

package main

import (
	"encoding/json"
	"flag"
	"github.com/wedojava/mytools"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const contentHeader, fileListHeader, contentFooter = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8">
    <title>Title</title>
</head>
<body>
`,
	`
<table>
<h2>Download list:</h2>
<tr>
	<td>Index</td><td>Files</td><td>Size</td><td>Time</td>
</tr>
`,
	`
</table>
</body>
</html>
`

type File struct {
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
	ModTime  int64  `json:"mod_time"`
}

func main() {
	// 1. download the encrypted db.json
	fileUrl := flag.String("u", "https://golangcode.com/images/avatar.jpg", "default is https://golangcode.com/images/avatar.jpg, -u and append url string can do what you need.")
	flag.Parse()
	savePath := "./db.encrypted"
	if err := GetDBFile(savePath, *fileUrl); err != nil {
		panic(err)
	}
	// 2. unrar and decrypt it
	data, err := ioutil.ReadFile(savePath)
	mytools.CheckPanic(err)
	strDB := mytools.AESDecrypt(string(data), "12345678901234567890123456789012")
	// 3. generate a html file for show file list
	var files []File
	err = json.Unmarshal([]byte(strDB), &files)
	mytools.CheckPanic(err)
	rawHtml := GenerateHtml(files)
	// 4. Write to html file
	err = ioutil.WriteFile("download.html", []byte(rawHtml), os.ModePerm)
	mytools.CheckPanic(err)
}

func GenerateHtml(files []File) (fileList string) {
	loc := time.FixedZone("UTC+8", +8*60*60)
	for i, file := range files {
		modTime := time.Unix(file.ModTime, 0)
		t := modTime.In(loc)
		fileList += "<tr><td>" + string(i) + "</td><td>" + file.Filename + "</td><td>" + string(file.Size) + "</td><td>" + t.Format(time.RFC1123Z) + "</td></tr>"
	}
	fileList = contentHeader + fileListHeader + fileList + contentFooter
	return
}

func GetDBFile(savePath, url string) error {
	resp, err := http.Get(url)
	mytools.Err(err)
	defer resp.Body.Close()

	out, err := os.Create(savePath)
	mytools.Err(err)
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
