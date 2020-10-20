package main

import (
	"log"
	"net/http"
	"pixelbattle/controller"
)

func main(){
	canvasController := controller.NewCanvas()
	http.HandleFunc("/", canvasController.GetCanvas)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
