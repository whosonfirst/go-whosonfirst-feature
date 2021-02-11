package feature

import (
	"context"
	"encoding/json"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/tidwall/gjson"
)

func NewGeometry(ctx context.Context, body []byte) (orb.Geometry, error) {

	geom_rsp := gjson.GetBytes(body, "geometry")

	if !geom_rsp.Exists() {
		return nil, &NotFoundError{"geometry"}
	}

	var geom *geojson.Geometry

	err := json.Unmarshal([]byte(geom_rsp.String()), &geom)

	if err != nil {
		return nil, err
	}

	return geom.Geometry(), nil
}
