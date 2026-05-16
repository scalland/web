package utils

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin/render"
)

// MultiTemplateRender implements gin's HTMLRender interface
type MultiTemplateRender map[string]*template.Template

func (r MultiTemplateRender) Instance(name string, data interface{}) render.Render {
	return render.HTML{
		Template: r[name],
		Name:     "base",
		Data:     data,
	}
}

// LoadTemplates loads HTML templates from an embedded filesystem.
func (u *Utils) LoadTemplates(webFS embed.FS) render.HTMLRender {
	u.TemplatesMap = make(map[string]*template.Template)

	// Collect all files
	var partials []string
	var pages []string
	var baseFile string

	fs.WalkDir(webFS, "web/template", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		if strings.Contains(path, "partials") {
			partials = append(partials, path)
		} else if strings.Contains(path, "pages") {
			pages = append(pages, path)
		} else if strings.HasSuffix(d.Name(), "base.gohtml") {
			baseFile = path
		}
		return nil
	})

	for _, pagePath := range pages {
		fullPageName := strings.Split(pagePath, "/")[len(strings.Split(pagePath, "/"))-1]

		// Combine base, partials, and the page itself.
		// pagePath comes last to ensure its definitions win.
		files := append([]string{baseFile}, partials...)
		files = append(files, pagePath)

		// Create a new template set and parse all files.
		// We use ParseFS directly which is more efficient and handles name association correctly.
		tmpl, err := template.ParseFS(webFS, files...)
		if err != nil {
			panic(fmt.Sprintf("Error parsing templates for %s: %v", fullPageName, err))
		}

		u.TemplatesMap[fullPageName] = tmpl
	}

	return MultiTemplateRender(u.TemplatesMap)
}

// RenderTemplate renders a template to the response writer.
func (u *Utils) RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	var buf bytes.Buffer
	tmpl, ok := u.TemplatesMap[name]
	if !ok {
		u.Logger.Errorf("Template %s not found", name)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err := tmpl.ExecuteTemplate(&buf, "base", data); err != nil {
		u.Logger.Errorf("Render %s: %s", name, err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}

// RenderTemplateToString renders a template to a string (for emails).
func (u *Utils) RenderTemplateToString(name string, data interface{}) (string, error) {
	var buf bytes.Buffer
	tmpl, ok := u.TemplatesMap[name]
	if !ok {
		return "", fmt.Errorf("template %s not found", name)
	}
	if err := tmpl.ExecuteTemplate(&buf, "base", data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
