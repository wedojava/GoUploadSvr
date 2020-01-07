package main

import (
	"encoding/hex"
	"fmt"
	"github.com/wedojava/MyErrCheck"
	"io"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  // parse form data, default will not do that.
	fmt.Println(r.Form)  // this will print at server-end
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello World!") // this is an view to client as a response.
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header  {
		for _, h:=range headers{
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// upload can used by curl like:`curl -F "file=@./go_test.ipynb" http://127.0.0.1:9090/upload`
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // Get the type of request
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		//r.Header.Set("Content-Type", "multipart/form-data")
		file, handler, err := r.FormFile("file")
		MyErrCheck.CheckErr(err, "upload:FormFile", "Println")
		defer file.Close()
		_, _ = fmt.Fprintf(w, "%v", handler.Header)
		//TODO: if upload failure, delete the folder
		saveFolder := "./files/ "+ RandStringBytesMaskImprSrc(6) + "/"
		err = os.MkdirAll(saveFolder, os.ModePerm)
		MyErrCheck.CheckErr(err, "Create folder for save file", "Println")
		f, err := os.OpenFile(saveFolder +handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		MyErrCheck.CheckErr(err, "upload:OpenFile", "Println")
        defer f.Close()
		_, err = io.Copy(f, file)
		MyErrCheck.CheckErr(err, "upload:Copy", "Println")
		fmt.Println("Uploading complete.")
	}
}



func RandStringBytesMaskImprSrc(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func randomBase16String(l int) string {
    buff := make([]byte, int(math.Round(float64(l)/2)))
    rand.Read(buff)
    str := hex.EncodeToString(buff)
    return str[:l] // strip 1 extra character we get from odd length results
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/upload", upload)

	err := http.ListenAndServe(":9090", nil)
	MyErrCheck.CheckErr(err, "ListenAndServe:", "")
}