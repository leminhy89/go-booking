package models

//TemplateData holds data of a template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} //interface{} mean any kind of data in Go
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
