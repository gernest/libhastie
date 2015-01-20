package libhastie

const (
	BEFORE_FILTER = iota
	AFTER_FILTER
)

type PipeLine struct {
	Targets       []string
	BeforeFilters []FilterFunc
	AfterFilters  []FilterFunc
}

type FilterFunc func(string) string

func (p *PipeLine) AddFilter(filter FilterFunc, pos int) *PipeLine {
	switch pos {
	case BEFORE_FILTER:
		p.BeforeFilters = append(p.BeforeFilters, filter)
		return p
	case AFTER_FILTER:
		p.AfterFilters = append(p.AfterFilters, filter)
		return p
	}
	return p
}

func (p *PipeLine) AddTarget(s string) *PipeLine {
	p.Targets = append(p.Targets, s)
	return p
}
func (p *PipeLine) Begin() {
	for _, v := range p.Targets {
		site := NewSite(v)
		site.BeforeFilters = p.BeforeFilters
		site.AfterFilters = p.AfterFilters
		site.Build()
	}
}

func (p *PipeLine) Run() {
	p.Before()
	p.Begin()
	p.After()
}

func (p *PipeLine) Before() {}
func (p *PipeLine) After()  {}

func NewPipeline() *PipeLine {
	return &PipeLine{}
}
