package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"sbstn.net/snippetbox/pkg/forms"
	"sbstn.net/snippetbox/pkg/models"
	"sbstn.net/snippetbox/ui"
)

type templateData struct {
	CSRFToken       string
	CurrentYear     int
	Flash           string
	Form            *forms.Form
	IsAuthenticated bool
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/*.page.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, page)
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseFS(ui.Files, "html/*.layout.tmpl")
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseFS(ui.Files, "html/*.partial.tmpl")
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil

}
