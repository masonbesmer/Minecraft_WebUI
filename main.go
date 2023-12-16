package main

import (
	"encoding/json"
	"html/template"
	"log"
	"math/big"
	"net/http"
)

var items []Item
var tmpl *template.Template

type Item struct {
	Damage      int    `json:"damage"`
	HasTag      bool   `json:"hasTag"`
	IsCraftable bool   `json:"isCraftable"`
	Label       string `json:"label"`
	MaxDamage   int    `json:"maxDamage"`
	MaxSize     int    `json:"maxSize"`
	Name        string `json:"name"`
	Size        int    `json:"size"`
}

type Fluid struct {
	Amount big.Int `json:"amount"`
	HasTag bool    `json:"hasTag"`
	Label  string  `json:"label"`
	Name   string  `json:"name"`
}

type Todo struct {
	Title string
	Done  bool
}

type IdxPageData struct {
	PageTitle string
	Items     []Item
}

// minecraft posts to this endpoint every time the Opencomputers Program checks AE stock levels.
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&items)
	if err != nil {
		log.Fatal(err.Error())
	}

	print(items)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1>Hellod</h1>"))
	data := IdxPageData{
		PageTitle: "My TOsDO list",
		Items:     items,
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	tmpl = template.Must(template.ParseFiles("index.html"))
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/update", UpdateHandler)
	log.Fatal(http.ListenAndServe(":3000", mux))
}

// package main

// import (
// 	"html/template"
// 	"net/http"
// )

// type Todo struct {
// 	Title string
// 	Done  bool
// }

// type TodoPageData struct {
// 	PageTitle string
// 	Todos     []Todo
// }

// func main() {
// 	tmpl := template.Must(template.ParseFiles("index.html"))
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		data := TodoPageData{
// 			PageTitle: "My TODO list",
// 			Todos: []Todo{
// 				{Title: "Task 1", Done: false},
// 				{Title: "Task 2", Done: true},
// 				{Title: "Task 3", Done: true},
// 			},
// 		}
// 		tmpl.Execute(w, data)
// 	})
// 	err := http.ListenAndServe(":3000", nil)
// 	if err != nil {
// 		panic(err)
// 	}
// }
