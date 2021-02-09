package feature

import (
	"context"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"io"
	"io/ioutil"
)

type Feature interface {
	Properties() []byte
	Geometry() orb.Geometry
}

func UnmarshalFeatureFromReader(ctx context.Context, r io.Reader) (Feature, error) {

	body, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, err
	}

	return UnmarshalFeature(ctx, body)
}

func UnmarshalFeature(ctx context.Context, body []byte) (Feature, error) {

	f, err := geojson.UnmarshalFeature(body)

	if err != nil {
		return nil, err
	}

	// FIX ME
	props := body

	return UnmarshalWhosOnFirstFeature(ctx, f.Geometry, props)
}
