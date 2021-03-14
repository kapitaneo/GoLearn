package controllers

import (
	"html/template"
	"net/http"

	"GoLearn/models"
)

type ViewData struct {
	Title string
	Users []models.Todo
}

// First method
func Sayhello(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("template.html").ParseFiles("template.html")

	todos := models.Get()

	data := ViewData{
		Title: "Users List",
		Users: todos,
	}

	tmpl.Execute(w, data)
}
