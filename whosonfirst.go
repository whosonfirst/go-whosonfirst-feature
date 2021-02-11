package feature

import (
	"context"
	"github.com/paulmach/orb"
)

type WhosOnFirstFeature struct {
	Feature
	properties *Properties
	geometry   orb.Geometry
}

func (f *WhosOnFirstFeature) Properties() *Properties {
	return f.properties
}

func (f *WhosOnFirstFeature) Geometry() orb.Geometry {
	return f.geometry
}

func UnmarshalWhosOnFirstFeature(ctx context.Context, geom orb.Geometry, props *Properties) (Feature, error) {

	f := &WhosOnFirstFeature{
		properties: props,
		geometry:   geom,
	}

	return f, nil
}
