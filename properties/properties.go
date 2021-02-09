package properties

import (
	"context"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// not sure about this yet...
type Properties interface {
	GetProperty(context.Context, string) (interface{}, error)
	SetProperty(context.Context, string, interface{}) error
}

func GetProperty(ctx context.Context, body []byte, path string) (interface{}, error) {

	rsp, err := GetPropertyAsGJSON(ctx, body, path)

	if err != nil {
		return nil, err
	}

	if !rsp.Exists() {
		return nil, &NotFoundError{path}
	}

	return rsp.Value(), nil
}

func GetPropertyAsGJSON(ctx context.Context, body []byte, path string) (gjson.Result, error) {

	rsp := gjson.GetBytes(body, path)
	return rsp, nil
}

func SetProperty(ctx context.Context, body []byte, path string, value interface{}) ([]byte, error) {

	return sjson.SetBytes(body, path, value)
}
