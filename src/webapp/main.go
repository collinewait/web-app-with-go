package main

import (
	"database/sql"
	"fmt"
	"github/collinewait/web-app-with-go/src/webapp/controller"
	"github/collinewait/web-app-with-go/src/webapp/middleware"
	"github/collinewait/web-app-with-go/src/webapp/model"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	templates := populateTemplates()
	db := connectToDatabase()
	defer db.Close()
	controller.Startup(templates)
	http.ListenAndServeTLS(":8000", "cert.pem", "key.pem", &middleware.TimeoutMiddleware{new(middleware.GzipMiddleware)})
}

func connectToDatabase() *sql.DB {
	db, err := sql.Open("postgres", "postgres://collinewaitire:wait@localhost/wait?sslmode=disable") // in a real app use ssl
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database: %v", err))
	}
	model.SetDatabase(db)
	return db
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
