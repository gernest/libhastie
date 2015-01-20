package libhastie

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	SiteName       string
	SourceDir      string
	LayoutDir      string
	PublishDir     string
	BaseUrl        string
	CategoryMash   map[string]string
	ProcessFilters map[string][]string
	UseMarkdown    bool
}

func (c *Config) Load(filename string) error {
	name, err := resolveConfigFile(filename)
	if err != nil {
		return err
	}
	return c.loadFromFile(name)
}

func (c *Config) loadFromFile(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, c)
	if err != nil {
		os.Exit(1)
	}
	return nil
}

func NewConfig(filename string) *Config {
	c := &Config{}
	err := c.Load(filename)
	if err != nil {
		os.Exit(1)
	}
	return c
}
func resolveConfigFile(filename string) (string, error) {
	if filename == "" {
		return os.Getwd()
	}
	f, err := os.Stat(filename)
	if err != nil {
		return "", err
	}
	if f.IsDir() {
		return filepath.Join(filename, "config.json"), nil
	}
	return filename, nil
}
