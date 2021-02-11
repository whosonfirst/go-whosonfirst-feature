package feature

import (
	"context"
	// "encoding/json"
	"github.com/paulmach/orb"
	// "github.com/paulmach/orb/geojson"
	"io"
	"io/ioutil"
)

type Feature interface {
	Properties() *Properties
	Geometry() orb.Geometry
	Body() []byte
}

func UnmarshalFeatureFromReader(ctx context.Context, r io.Reader) (Feature, error) {

	body, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, err
	}

	return UnmarshalFeature(ctx, body)
}

func UnmarshalFeature(ctx context.Context, body []byte) (Feature, error) {

	return UnmarshalWhosOnFirstFeature(ctx, body)
}
