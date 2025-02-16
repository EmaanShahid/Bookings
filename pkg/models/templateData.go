package models

// TemplateData holds data sent from handlers to templates

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string //required when using forms
	Flash     string //warning messages in forms
	Warning   string
	Error     string
}
