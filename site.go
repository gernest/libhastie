package libhastie

import "os"

type SiteStruct struct {
	Files       []string
	Directories []string
	Categories  []string
}

func (s *SiteStruct) Load(base string) error {
	return loadSite(base, s)
}

// Fill a given site with files and directories
func loadSite(base string, site *SiteStruct) error {
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
				site.Directories = append(site.Directories, srcFile)
				err = loadSite(srcFile, site)
			} else {
				site.Files = append(site.Files, srcFile)
			}
		}
	} else {
		site.Files = append(site.Files, base)
	}
	return nil
}
