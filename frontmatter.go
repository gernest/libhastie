package libhastie

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"strings"
)

func FrontMatter(filename string) (map[string]string, string, error) {
	data, err := os.Open(filename)
	defer data.Close()
	if err != nil {
		return nil, "", err
	}
	scan := bufio.NewScanner(data)
	return readFrontMatter(scan)
}
func readFrontMatter(s *bufio.Scanner) (map[string]string, string, error) {
	m := make(map[string]string)
	b := new(bytes.Buffer)
	infm := false
	n := 0
	for s.Scan() {
		l := s.Text()
		n++
		if l == "---" {
			if infm {
				infm = false
			} else {
				infm = true
			}
		} else if infm {
			sections := strings.SplitN(l, ":", 2)
			if len(sections) != 2 {
				return nil, "", errors.New("Invarid front matter")
			}
			m[sections[0]] = strings.Trim(sections[1], " ")

		} else {
			b.WriteString(s.Text())
		}
	}
	if err := s.Err(); err != nil {
		return nil, "", err
	}
	return m, b.String(), nil
}
