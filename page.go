package libhastie

import (
	"path/filepath"
	"time"
)

type Page struct {
	Content        string
	Title          string
	Category       string
	SimpleCategory string
	Layout         string
	OutFile        string
	Extension      string
	Url            string
	PrevUrl        string
	PrevTitle      string
	NextUrl        string
	NextTitle      string
	PrevCatUrl     string
	PrevCatTitle   string
	NextCatUrl     string
	NextCatTitle   string
	Params         map[string]string
	Recent         *PagesSlice
	Date           time.Time
	Categories     *CategoryList
}

func NewPage(filename string) *Page {
	epoch, _ := time.Parse("20060102", "19700101")
	return &Page{
		Title:          "",
		Category:       "",
		SimpleCategory: "",
		Content:        "",
		Layout:         "",
		Date:           epoch,
		OutFile:        filename,
		Extension:      filepath.Ext(filename),
		Url:            "",
		PrevUrl:        "",
		PrevTitle:      "",
		NextUrl:        "",
		NextTitle:      "",
		PrevCatUrl:     "",
		PrevCatTitle:   "",
		NextCatUrl:     "",
		NextCatTitle:   "",
		Params:         make(map[string]string),
	}
}
