package main

import (
	"github/collinewait/web-app-with-go/src/webapp/controller"
	"github/collinewait/web-app-with-go/src/webapp/middleware"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	templates := populateTemplates()
	controller.Startup(templates)
	http.ListenAndServe(":8000", &middleware.TimeoutMiddleware{new(middleware.GzipMiddleware)})
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to load template blocks directory: " + err.Error())
	}
	files, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, file := range files {
		f, err := os.Open(basePath + "/content/" + file.Name())
		if err != nil {
			panic("Failed to open template '" + file.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + file.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + file.Name() + "' as template")
		}
		result[file.Name()] = tmpl
	}
	return result
}
