package uuidx

type Fixed struct {
	Value string
}

func (g *Fixed) New() string {
	return g.Value
}

func (g *Fixed) Reset() {}
