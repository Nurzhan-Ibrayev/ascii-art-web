package main

import (
	"html/template"
	"log"
	"net/http"
	"template/asciiart"
)

type AsciiText struct {
	Text   string
	Style  string
	Result string
	Err    string
}

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("template/*.html"))
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorsPage(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	temp.ExecuteTemplate(w, "template.html", nil)
}

func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorsPage(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	var data AsciiText
	data.Text = r.FormValue("text")
	data.Style = r.Form["choice"][0]

	for k := range r.PostForm {
		if !(k == "text" || k == "choice") {
			ErrorsPage(w, http.StatusBadRequest, "Bad Request")
			return
		}
	}

	_, err0 := asciiart.CheckHash(data.Style)
	if !err0 {
		ErrorsPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	var err error
	data.Result, err = asciiart.AsciiArt(data.Text, data.Style)
	if check(err) {
		ErrorsPage(w, http.StatusBadRequest, "Bad Request")
		return
	}
	err = temp.ExecuteTemplate(w, "template.html", data)
	if check(err) {
		log.Fatalln("error in head Handler")
	}
}

func check(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		{
			GetHandler(w, r)
		}
	case "/ascii-art":
		{
			HtmlHandler(w, r)
		}
	default:
		{
			ErrorsPage(w, http.StatusNotFound, "Page not found")
		}
	}
}

func ErrorsPage(w http.ResponseWriter, code int, str string) {
	w.WriteHeader(code)
	var data AsciiText
	data.Err = str
	var err error
	err = temp.ExecuteTemplate(w, "error.html", data.Err)
	if check(err) {
		log.Fatalln("error in head Handler")
	}
}

func main() {
	log.Print("http://localhost:8000")
	http.HandleFunc("/", pathHandler)
	http.HandleFunc("/ascii-art", pathHandler)
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	err := http.ListenAndServe(":8000", nil)
	log.Fatal(err)
}
