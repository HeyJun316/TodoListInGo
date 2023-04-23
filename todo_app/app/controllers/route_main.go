package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	/*
		t, err := template.ParseFiles("/Users/heyjun/Desktop/todo_go/todo_app/views/templates/top.html")
		if err != nil {
			log.Fatalln(err)
		}
		t.Execute(w, "Hellow")
	*/
	generateHTML(w, "hello", "layout", "top")
}
