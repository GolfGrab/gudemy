package models

// TemplateData is a struct that holds the data sent from handlers to the template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float64
	BoolMap   map[string]bool
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Worning   string
	Error     string
}
