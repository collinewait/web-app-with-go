package controller

import (
	"github/collinewait/web-app-with-go/src/webapp/viewmodel"
	"html/template"
	"net/http"
)

type shop struct {
	shopTemplate *template.Template
}

func (s shop) registerRoutes() {
	http.HandleFunc("/shop", s.handleShop)
}

func (s shop) handleShop(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewShop()
	s.shopTemplate.Execute(w, vm)
}
