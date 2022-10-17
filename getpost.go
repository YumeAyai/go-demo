package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type Ret struct {
	Value string
}

var value string

func main() {
	post()
	http.HandleFunc("/data", Request)
	http.ListenAndServe("127.0.0.1:8081", nil)
}

func post() {
	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		value = r.FormValue("value")

		tmpl.Execute(w, struct{ Success bool }{true})
		fmt.Printf("%s", value)
	})
}
func Request(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ret := new(Ret)
		ret.Value = value
		ret_json, _ := json.Marshal(ret)

		io.WriteString(w, string(ret_json))
	}
}
