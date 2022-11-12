package renders

import (
	"bytes"
	"github.com/justinas/nosurf"
	config2 "github.com/prashanth-gajula/bookings/internal/config"
	models2 "github.com/prashanth-gajula/bookings/internal/models"
	//config2 "github/prashanth-gajula/go-course/pkg/config"
	//"github/prashanth-gajula/go-course/pkg/models"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var app *config2.AppConfig

func NewTemplates(a *config2.AppConfig) {
	app = a
}
func AddDefaultData(td *models2.TemplateData, r *http.Request) *models2.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}
func RenderTemplate(w http.ResponseWriter, r *http.Request, html string, td *models2.TemplateData) {
	//the function will take the response variable and the name of the template
	//parse it and write the information to the browser window
	// create a template cache
	//get the template cache from the app config
	var tc map[string]*template.Template
	//fmt.Println("welcome")
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//Get requested template from cache
	t, ok := tc[html]
	if !ok {
		log.Fatal("couldn't get the template from template cache not working")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println("Error while executing:", err)
	}
	//rendering the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("Error while writing to buffer:", err)
	}
}
func CreateTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	//the above and below code will work similarly
	myCache := map[string]*template.Template{}
	//get all the files starting with .html from templates
	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return myCache, err
	}
	//range through all the files in pages
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		//fmt.Println(len(matches))
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts //log.Println(myCache)
	} //fmt.Println(myCache)
	return myCache, nil
}
