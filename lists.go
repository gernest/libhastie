package libhastie

import "sort"

type PagesSlice []Page

func (p PagesSlice) Get(i int) Page         { return p[i] }
func (p PagesSlice) Len() int               { return len(p) }
func (p PagesSlice) Less(i, j int) bool     { return p[i].Date.Unix() < p[j].Date.Unix() }
func (p PagesSlice) Swap(i, j int)          { p[i], p[j] = p[j], p[i] }
func (p PagesSlice) Sort()                  { sort.Sort(p) }
func (p PagesSlice) Limit(n int) PagesSlice { return p[0:n] }

type CategoryList map[string]PagesSlice

func (c CategoryList) Get(category string) PagesSlice { return c[category] }
