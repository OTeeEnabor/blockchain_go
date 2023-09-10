package constant

// a package to get html template type in golang
import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// define struct that points to template folder location
type Template struct{
	// key templates of pointer type of base template.Template
	templates *template.Template
}
// create struct method, render, for Template type
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error{
	// write data from disk
	// name is template name
	// interface used to write data
	// context

	return t.templates.ExecuteTemplate(w, name, data)

}
// define where to find template folder
func LoadTemplate() *Template {
	// 
	template := &Template{
		// set location for where html must be searched
		templates: template.Must(template.ParseGlob("repository/templates/*.html")),
	}
	return template
}