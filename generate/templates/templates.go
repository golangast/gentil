package generate

var Servertemp = `

package main

	import (
		"fmt"
		"html/template"
		"io"
		"net/http"
		_ "net/http/pprof"
		"os"
		"path/filepath"
		_"time"
		"context"
		"strings"
		//#import
		"github.com/labstack/echo/v4"
		"github.com/labstack/echo/v4/middleware"
		"github.com/labstack/gommon/log"

	)
	
	type TemplateRenderer struct {
		templates *template.Template
	}
	
	// Render renders a template document
	func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	
		// Add global methods if data is a map
		if viewContext, isMap := data.(map[string]interface{}); isMap {
			viewContext["reverse"] = c.Echo().Reverse
		}
	
		return t.templates.ExecuteTemplate(w, name, data)
	}
	
	var err error
	
	func main() {
		//#createdb
		e := echo.New()
		t, err := ParseDirectory("templates/")
		if err != nil {
			fmt.Println(err)
		}
		renderer := &TemplateRenderer{
			templates: template.Must(t, err),
		}
	
		e.Renderer = renderer
	
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		}))
	
		Routes(e)
	
		// Route
		e.Logger.SetLevel(log.ERROR)
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
			XSSProtection:         "1; mode=block",
			ContentTypeNosniff:    "nosniff",
			XFrameOptions:         "SAMEORIGIN",
			HSTSMaxAge:            3600,
			ContentSecurityPolicy: "",
		}))
		e.Use(middleware.BodyLimit("3M"))
		e.IPExtractor = echo.ExtractIPDirect()
		e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}))
		e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(30)))
		e.Static("/static", "static")
		// Start server
		e.Logger.Fatal(e.Start(":3000"))
		
	
		
		
	}
	func GetAllFilePathsInDirectory(dirpath string) ([]string, error) {
		var paths []string
		err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				paths = append(paths, path)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	
		return paths, nil
	}
	
	func ParseDirectory(dirpath string) (*template.Template, error) {
		paths, err := GetAllFilePathsInDirectory(dirpath)
		if err != nil {
			return nil, err
		}
		return template.ParseFiles(paths...)
	}
	


func Routes(e *echo.Echo) {
	e.GET("/", Home)
	e.GET("/route/:routes", List)
	//#routes
}

func Home(c echo.Context) error {

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{})

}
//#handler
func List(c echo.Context) error {
	var data []Data
	routes := c.Param("routes")
	nospaceroutes := strings.ReplaceAll(routes, " ", "")
	nospaceroutesnoslash := strings.ReplaceAll(nospaceroutes, "/", "")
	//#databaseconn
	return c.Render(http.StatusOK, nospaceroutesnoslash+".html", map[string]interface{}{
		"data":data,
	})

}


`

var Bodytemp = `
{{ .header }}
<!-- ### -->
{{ .footer }}
`
var Footertemp = `
{{.footer}}
</body>
<!-- ### -->
</html>
{{.end}}
`
var Headertemp = `
{{.header}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Welcome</title>
    <!-- ### -->
</head>
{{.end}}
`
