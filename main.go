package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello you've requested: %s/n", r.URL.Path)
	// w.Write([]byte("<a action=\"/pic\" method=\"GET\"> here! </a>"))
	w.Write([]byte("<h1>Hello World!</h1>\n<a href=\"/pic/Screenshot_20240112_034109.png\">cac pic</a>"))
	w.Write([]byte("<br><a href=\"/pic/test.html\">html test</a>"))
	w.Write([]byte("<br><a href=\"/page/1\">page 1</a>"))

	// fmt.Fprint(w, "get : ", r.URL.Query().Get("A"))
	// fmt.Fprint(w, "\npost : ", r.FormValue("B"))
}
func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a := vars["page"]
	w.Write([]byte("<h1>page number is" + a + "</h1>"))
	i, err := strconv.Atoi(a)
	if err != nil {
		fmt.Printf("fuck me")
	}
	w.Write([]byte("<br><a href=\"/page/" + strconv.Itoa(i+1) + "\">next</a>"))
	w.Write([]byte("<br><a href=\"/page/" + strconv.Itoa(i-1) + "\">previous</a>"))
	w.Write([]byte("<br><a href=\"/\">main page</a>"))
}

func main() {
	fmt.Print("start!")

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("static"))

	r.Handle("/pic/", http.StripPrefix("/pic/", fs))
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/page/{page}", pageHandler)
	http.ListenAndServe(":8080", r)
}
