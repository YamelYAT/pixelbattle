package controller

import (
	"log"
	"net/http"
	"pixelbattle/canvas"
)

type CanvasController struct {
	Canvas	*canvas.Canvas
}

func NewCanvas() *CanvasController{
	return &CanvasController{Canvas: canvas.NewCanvas()}
}

func (c *CanvasController)GetCanvas(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	w.Header().Set("content-type", "text/html; charset=utf-8")
	w.Header().Set("content-length", "636401")
	if _, err := w.Write(c.Canvas.Points); err != nil{
		log.Println(err)
	}
}
