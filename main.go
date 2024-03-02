package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

var calcTmp = template.Must(template.ParseFiles("static/calc.html"))

type Data struct {
	Result float64
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello you've requested: %s/n", r.URL.Path)
	// w.Write([]byte("<a action=\"/pic\" method=\"GET\"> here! </a>"))
	w.Write([]byte("<h1>Hello World!</h1>\n<a href=\"/pic/Screenshot_20240112_034109.png\">cac pic</a>"))
	w.Write([]byte("<br><a href=\"/pic/test.html\">html test</a>"))
	w.Write([]byte("<br><a href=\"/page/1\">page 1</a>"))
	w.Write([]byte("<br><a href=\"/calc/0\">calc</a>"))

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
func do(a float64, b float64, op string) float64{
	switch op{
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	case "":
		return 0

	}
	return float64(0)
}
func calcHandler(w http.ResponseWriter, r *http.Request) {
	data := &Data{
		Result: 0,
	}
	var val1 float64 = 0
	var val2 float64 = 0
	op := ""
	Cval1, err := r.Cookie("Val1")
    if err != nil {
        switch {
        case errors.Is(err, http.ErrNoCookie):
            // http.Error(w, "cookie not found", http.StatusBadRequest)
        default:
            log.Println(err)
            http.Error(w, "server error", http.StatusInternalServerError)
        }
        // return
    } else {
		val1,_ = strconv.ParseFloat(Cval1.Value,64)
	}
	Cval2, err := r.Cookie("Val2")
    if err != nil {
        switch {
        case errors.Is(err, http.ErrNoCookie):
            // http.Error(w, "cookie not found", http.StatusBadRequest)
        default:
            log.Println(err)
            http.Error(w, "server error", http.StatusInternalServerError)
        }
        // return
    } else {
		val2,_ = strconv.ParseFloat(Cval2.Value,64)
	}
	Cres, err := r.Cookie("Res")
    if err != nil {
        switch {
        case errors.Is(err, http.ErrNoCookie):
            // http.Error(w, "cookie not found", http.StatusBadRequest)
        default:
            log.Println(err)
            http.Error(w, "server error", http.StatusInternalServerError)
        }
        // return
    } else {
		data.Result,_ = strconv.ParseFloat(Cres.Value,64)
	}
	Cop, err := r.Cookie("Op")
    if err != nil {
        switch {
        case errors.Is(err, http.ErrNoCookie):
            // http.Error(w, "cookie not found", http.StatusBadRequest)
        default:
            log.Println(err)
            http.Error(w, "server error", http.StatusInternalServerError)
        }
        // return
    } else {
		op = Cop.Value
	}




	vars := mux.Vars(r)

	a, err := strconv.Atoi(vars["action"])
	if err != nil {
		fmt.Printf("fuck me")
	}
	switch a {
	case 10:
		if val1 == 0 {
			data.Result = math.Sqrt(data.Result)
		} else { 
			val2 = data.Result
			data.Result = do(val1,val2,op)
		}
		op = "%"
	case 11:
		data.Result = math.Sqrt(data.Result)
		op = "r"
	case 12:
		data.Result = math.Floor(data.Result/10)
		op = "ce"
	case 13:
		data.Result = 0
		val1 = 0
		val2 = 0
		op = ""
	case 14:
		if val1 == 0 {
			val1 = data.Result
			data.Result = 0
		} else { 
			val2 = data.Result
			data.Result = do(val1,val2,op)
			val1 = data.Result
			val2 = 0
		}
		op = "-"
	case 15:
		if val1 == 0 {
			val1 = data.Result
			data.Result = 0
		} else { 
			val2 = data.Result
			data.Result = do(val1,val2,op)
			val1 = data.Result
			val2 = 0
		}
		op = "/"
	case 16:
		if val1 == 0 {
			val1 = data.Result
			data.Result = 0
		} else { 
			val2 = data.Result
			data.Result = do(val1,val2,op)
			val1 = data.Result
			val2 = 0
			data.Result = 0
		}
		op = "*"
//todo-make the . right
	case 17:
		
		op = "d"
	case 18:
		if val1 == 0 {
			
		} else {
			val2 = data.Result
			data.Result = do(val1,val2,op)
			val1 = data.Result
			val2 = 0
		}
		// op = ""
	case 19:
		if val1 == 0 {
			val1 = data.Result
			data.Result = 0
		} else { 
			val2 = data.Result
			data.Result = do(val1,val2,op)
			val1 = data.Result
			val2 = 0
		}
		op = "+"
	default:
		// data.Result,err = strconv.ParseFloat(strconv.FormatFloat(data.Result*10, 'f', -1, 64)  + strconv.Itoa(a),64)
		data.Result = data.Result * 10 + float64(a)
		
	if err != nil {
		fmt.Printf("fuck me")
	}
	}


	Csval1 := http.Cookie{
		Name:     "Val1",
		Value:     strconv.FormatFloat(val1,'g', -1, 64),
	}
	Csval2 := http.Cookie{
		Name:    "Val2",
		Value:   strconv.FormatFloat(val2, 'g', -1, 64),
	}
	Csres := http.Cookie{
		Name:     "Res",
		Value:	  strconv.FormatFloat(data.Result,'g',-1,64),
	}
	Csop := http.Cookie{
		Name:     "Op",
		Value:	  op,
	}
	http.SetCookie(w, &Csval1)
	http.SetCookie(w, &Csval2)
	http.SetCookie(w, &Csres)
	http.SetCookie(w, &Csop)

	buf := &bytes.Buffer{}
	err = calcTmp.Execute(buf, data)
	if err != nil {
		// fmt.Print("hey")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}


func main() {
	fmt.Print("start!")

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("static/"))

	// r.Handle("/pic/{rest}", http.StripPrefix("/pic", fs))
	r.PathPrefix("/pic/").Handler(http.StripPrefix("/pic", fs))

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/page/{page}", pageHandler)
	r.HandleFunc("/calc/{action}", calcHandler)
	http.ListenAndServe(":8080", r)
}
