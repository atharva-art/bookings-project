package models

// TemplateData holds data to send to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]string
	FloatMap map[string]float32
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Error string
}