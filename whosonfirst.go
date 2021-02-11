package feature

import (
	"context"
	"github.com/paulmach/orb"
)

type WhosOnFirstFeature struct {
	Feature
	body []byte
}

func (f *WhosOnFirstFeature) Properties() *Properties {
	return nil
}

func (f *WhosOnFirstFeature) Geometry() orb.Geometry {
	ctx := context.Background()
	geom, _ := NewGeometry(ctx, f.body)
	return geom
}

func (f *WhosOnFirstFeature) Body() []byte {
	return f.body
}

func UnmarshalWhosOnFirstFeature(ctx context.Context, body []byte) (Feature, error) {

	f := &WhosOnFirstFeature{
		body: body,
	}

	return f, nil
}
