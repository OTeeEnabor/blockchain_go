package constant

// a package to get html template type in golang
import (
	"fmt"
	"html/template"
	"io"

	"os"
	"path/filepath"

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
	//  get app path
	path, _ := os.Executable()
	// get file path
	filePath := filepath.Dir(path)
	// store  template folder path into variable (CLOUD server)
	templateFolder := fmt.Sprintf("%v/repository/templates/*",filePath)

	template := &Template{
		// set location for where html must be searched
		templates: template.Must(template.ParseGlob(templateFolder)),
	}
	return template
}