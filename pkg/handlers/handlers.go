package handlers

import (
	"book-room/pkg/render"
	"net/http"
)

// home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTmpl(w, "home.page.html")
}

// about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTmpl(w, "about.page.html")
}
