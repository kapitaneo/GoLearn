package main

import (
	"net/http"

	"GoLearn/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Sayhello)
	http.ListenAndServe(":8181", nil)
}
