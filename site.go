package libhastie

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type SiteStruct struct {
	Files         []string
	Categories    []string
	BeforeFilters []FilterFunc
	AfterFilters  []FilterFunc
	Config        *Config
	Pages         PagesSlice
	AllowedExt    []string
	Path          string
}

func (s *SiteStruct) Load(base string) error {
	return loadFiles(base, s)
}
func (s *SiteStruct) LoadPages() {
	s.loadPages()
}
func (s *SiteStruct) Before() {
	s.LoadPages()
}
func (s *SiteStruct) After() {}

func (s *SiteStruct) Build() error {
	s.Before()
	layouGlob := s.Path + "/" + s.Config.LayoutDir + "/*.html"

	lfiles, err := filepath.Glob(layouGlob)
	if err != nil {
		return err
	}
	t, err := template.ParseFiles(lfiles...)
	if err != nil {
		return err
	}
	categories := s.getCategoryList()
	recent := s.getRecentList()
	for _, page := range s.Pages {
		page.Recent = &recent
		page.Categories = &categories

		buf := new(bytes.Buffer)

		page.buildPrevNextLinks(&recent)
		if s.Config.BaseUrl != "" {
			page.Params["BaseUrl"] = s.Config.BaseUrl
		}
		templateFile := page.Layout + ".html"
		err = t.ExecuteTemplate(buf, templateFile, &page)
		if err != nil {
			return err
		}

		cdir := strings.TrimLeft(strings.ToLower(filepath.Dir(page.Path)), filepath.Join(strings.ToLower(s.Path), strings.ToLower(s.Config.SourceDir)))

		if page.Category != "" {
			cdir = page.Category
		}
		outDir := filepath.Join(s.Path, filepath.Join(s.Config.PublishDir, cdir))

		writeDir := outDir
		os.MkdirAll(writeDir, 0755)

		outfile := filepath.Join(writeDir, page.OutFile+".html")
		ioutil.WriteFile(outfile, buf.Bytes(), 0655)
	}
	s.After()
	return nil
}

func (s *SiteStruct) getCategoryList() CategoryList {
	mapList := make(CategoryList)

	reverseMap := make(map[string]string)

	for k, v := range s.Config.CategoryMash {
		cats := strings.Split(string(v), ",")
		for _, cat := range cats {
			reverseMap[cat] = string(k)
		}
	}

	for _, page := range s.Pages {
		if reverseMap[page.Category] != page.Category {
			thisCategory := reverseMap[page.Category]
			mapList[thisCategory] = append(mapList[thisCategory], page)
		}
		simpleCategory := strings.Replace(page.Category, "/", "_", -1)
		mapList[simpleCategory] = append(mapList[simpleCategory], page)
	}
	return mapList
}

func (s *SiteStruct) loadPages() {

	pgs := LoadFiles(s.Path + "/" + s.Config.SourceDir)
	for _, p := range pgs {
		for _, allow := range s.AllowedExt {

			if strings.ToLower(filepath.Ext(p)) == allow {
				page := NewPage(p)
				page.Path = p
				s.Pages = append(s.Pages, *page)
			}
		}

	}
}
func (s *SiteStruct) getRecentList() (list PagesSlice) {
	list = []Page{}
	for _, page := range s.Pages {
		list = append(list, page)
	}
	list.Sort()

	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}

	return list
}

func NewSite(base string) *SiteStruct {
	s := &SiteStruct{}
	err := s.Load(base)
	if err != nil {
		os.Exit(1)
	}
	s.Config = NewConfig(base)
	s.AllowedExt = []string{".md", ".html"}
	s.Path = base
	return s
}

func loadFiles(base string, site *SiteStruct) error {
	src, err := os.Stat(base)
	if err != nil {
		return err
	}
	if src.IsDir() {
		dir, _ := os.Open(base)
		defer dir.Close()
		objects, err := dir.Readdir(-1)
		if err != nil {
			return err
		}
		for _, obj := range objects {
			srcFile := base + "/" + obj.Name()
			if obj.IsDir() {
				loadFiles(srcFile, site)
			} else {
				site.Files = append(site.Files, srcFile)
			}
		}
	} else {
		site.Files = append(site.Files, base)
	}
	return nil
}

func getAllowedGlog(s []string) []string {
	slice := make([]string, len(s))
	for k, v := range s {
		slice[k] = "*" + v
	}
	return slice
}

type Files struct {
	slice []string
}

func LoadFiles(base string) []string {
	slice := new(Files)
	walk(base, slice)
	return slice.slice
}
func walk(base string, s *Files) {
	filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if path != "" {
			if info.IsDir() {
				return nil
			}
			s.slice = append(s.slice, path)
			return nil
		}
		return nil
	})
}
