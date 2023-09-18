package uuidx

import "github.com/google/uuid"

type uuidx interface {
	New() string
	Reset()
}

var Default = &Uuid{}
var UUIDX uuidx = Default

func New() string {
	return UUIDX.New()
}

func Reset() {
	UUIDX.Reset()
}

type Uuid struct{}

func (g *Uuid) New() string {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	return id.String()
}

func (g *Uuid) Reset() {}
