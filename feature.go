package feature

import (
	_ "context"
	"github.com/paulmach/orb"
)

type Feature interface {
	Bytes() []byte
	Geometry() orb.Geometry
}


