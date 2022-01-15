package cyoasubpkg

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/hankehly/gophercises/pkg/cyoa"
)

// This is set equal to the value of defaultTemplate inside 'init'
// before program execution
var tpl *template.Template

var defaultTemplate string = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Choose Your Own Adventure - {{ .Title }}</title>
</head>
<body>
    <h1>{{ .Title }}</h1>
    {{ range .Story }}
    <p>{{ . }}</p>
    {{ end }}
    <ul>
        {{ range .ChapterOptions }}
        <li>
            <a href="/{{ .Arc }}">{{ .Text }}</a>
        </li>
        {{ end }}
    </ul>
</body>
</html>
`

// Using init, we can validate our HTML template before execution
// We initialize the validated tpl value here so it can be used elsewhere
func init() {
	// log.Println("init called")
	tpl = template.Must(template.New("").Parse(defaultTemplate))
}

type HandlerOption func(h *handler)

func WithTemplate(tmpl template.Template) HandlerOption {
	return func(h *handler) {
		h.t = tmpl
	}
}

// An example of "Function Options" for configuration
// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
func NewHandler(s cyoa.Story, opts ...HandlerOption) http.Handler {
	// Initialize your value with defaults
	h := handler{s, *tpl}
	// Allow the programmer to modify the handler using any number of
	// HandlerOption functions
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

// handler is an http.Handler because it implements ServeHTTP
// but we can add other properties to it
type handler struct {
	s cyoa.Story
	t template.Template
}

func (h handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// Note:
	// Argument to New must be base name of template file (if using a file)
	// https://pkg.go.dev/text/template@go1.17.6#Template.ParseFiles
	// pathTemplate := "/Users/hankehly/Projects/gophercises/cyoa/template.html"
	// baseName := filepath.Base(pathTemplate)
	// tmpl, err := template.New(baseName).ParseFiles(pathTemplate)
	//
	path := strings.TrimSpace(r.URL.Path)
	parts := strings.Split(path, "/")
	key := parts[len(parts)-1]
	if key == "" {
		key = "intro"
	}
	if chapter, ok := h.s[key]; ok {
		err := h.t.Execute(rw, chapter)
		if err != nil {
			fmt.Printf("%v", err)
			http.Error(rw, fmt.Sprintln("Internal server error"), http.StatusInternalServerError)
		}
		return
	}
	http.Error(rw, "Invalid chapter name", http.StatusNotFound)
}
