package libhastie

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Page struct {
	Content        string
	Title          string
	Category       string
	SimpleCategory string
	Layout         string
	OutFile        string
	Path           string
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

func (page *Page) buildPrevNextLinks(recentList *PagesSlice) {
	foundPage := false

	nextPage := Page{}
	prevPage := Page{}
	nextPageCat := Page{}
	prevPageCat := Page{}
	lastPageCat := Page{}

	for i, rp := range *recentList {
		if rp.Category == page.Category {
			if foundPage {
				prevPageCat = rp
				break
			}
		}

		if rp.Title == page.Title {
			foundPage = true
			nextPageCat = lastPageCat
			if i != 0 {
				nextPage = recentList.Get(i - 1)
			}
			if i+1 < recentList.Len() {
				prevPage = recentList.Get(i + 1)
			}
		}

		if rp.Category == page.Category {
			lastPageCat = rp // previous page
		}
	}

	page.NextUrl = nextPage.Url
	page.NextTitle = nextPage.Title
	page.PrevUrl = prevPage.Url
	page.PrevTitle = prevPage.Title

	page.NextCatUrl = nextPageCat.Url
	page.NextCatTitle = nextPageCat.Title
	page.PrevCatUrl = prevPageCat.Url
	page.PrevCatTitle = prevPageCat.Title
}
func NewPage(filename string) *Page {
	front, body, err := FrontMatter(filename)
	if err != nil {
		os.Exit(1)
	}
	epoch, _ := time.Parse("20060102", "19700101")
	outfile := strings.ToLower(filepath.Base(filename))
	outfile = strings.TrimSuffix(outfile, filepath.Ext(outfile))
	page := &Page{
		Title:          front["title"],
		Category:       front["category"],
		SimpleCategory: "",
		Content:        body,
		Layout:         front["layout"],
		Date:           epoch,
		OutFile:        outfile,
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
		Params:         front,
	}
	return page
}
