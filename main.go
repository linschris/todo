package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Todo struct {
	Item string
	Completed bool
}
type PageData struct {
	Todos []Todo
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h);
		}
	}
}

func todoList(w http.ResponseWriter, req *http.Request) {
	temp := template.Must(template.ParseFiles("static/routes/landing/index.html"));
	todos := []Todo{
		{ Item: "Thing 1", Completed: true},
		{ Item: "Thing 3", Completed: false}, 
	}
	pageData := PageData{
		Todos: todos,
	};
	temp.Execute(w, pageData);	
}


func main() {
	http.HandleFunc("/", todoList);
	fs := http.FileServer(http.Dir("./static/routes/landing"));
	strip_fs := http.StripPrefix("/static", fs);
	http.Handle("/static", strip_fs);

	http.HandleFunc("/headers", headers);
	http.ListenAndServe(":8080", nil);
}