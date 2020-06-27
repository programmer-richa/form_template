package models
// struct Page contains data to be presented on the template, url of the page and template file name.
type Page struct {
	Title             string
	FileName          string
	Forms map[string]*Form
}

// NewPage accepts Title, File name , URL path and slice of forms
// It returns a variable of type Page
func NewPage(title string, fileName string, forms map[string]*Form) *Page {
	page := Page{
		Title:    title,
		FileName: fileName,
		Forms:    forms,
	}
	return &page
}

