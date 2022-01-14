package config

import "text/template"

type AppConfig struct {
	UseCache       bool
	CachedTemplate map[string]*template.Template
}
