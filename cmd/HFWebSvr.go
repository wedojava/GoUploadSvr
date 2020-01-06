package hfwebserver

import (
	"encoding/json"
	"net/http"
	"os"
	"path"
	"strconv"
)

// ErrJustRt will just return while error occur
func ErrJustRt(e error) {
	if e != nil {
		return
	}
}

func enterHurdle(keyword string) {
	if len(os.Args) > 1 {
		if os.Args[1] == keyword {
			println("Here we go.")
		}
	}
	os.Exit(3)
}

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func httpSvr(Port string) {
	server := http.Server{
		Addr: "127.0.0.1:" + Port,
	}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method{
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path)) // ?
	ErrJustRt(err)
	post, err := retrieve(id)
	ErrJustRt(err)
	output, err := json.MarshalIndent(&post, "", "\t\t")
	ErrJustRt(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var post Post
	json.Unmarshal(body, &post)
	err = post.create()
	ErrJustRt(err)
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	ErrJustRt(err)
	post, err := retrieve(id)
	ErrJustRt(err)
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &post)
	err = post.update()
	ErrJustRt(err)
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	ErrJustRt(err)
	post, err := retrieve(id)
	ErrJustRt(err)
	err = post.delete()
	ErrJustRt(err)
	w.WriteHeader(200)
	return
}


func hfwebserver() {
	enterHurdle("yeezy700")
}
