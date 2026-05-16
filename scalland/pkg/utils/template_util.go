package utils

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"bytes"
)

// LoadTemplates loads HTML templates from an embedded filesystem.
func (u *Utils) LoadTemplates(webFS embed.FS) {
	tmpl := template.New("")
	fs.WalkDir(webFS, "web/template", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		data, _ := fs.ReadFile(webFS, path)
		tmpl, _ = tmpl.New(d.Name()).Parse(string(data))
		return nil
	})
	u.Templates = tmpl
}

// RenderTemplate renders a template to the response writer.
func (u *Utils) RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	var buf bytes.Buffer
	if err := u.Templates.ExecuteTemplate(&buf, name, data); err != nil {
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
	if err := u.Templates.ExecuteTemplate(&buf, name, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
