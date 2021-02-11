package feature

import (
	"context"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"sync"
)

type Properties struct {
	mu   *sync.RWMutex
	body []byte
}

func NewProperties(ctx context.Context, body []byte) (*Properties, error) {

	mu := new(sync.RWMutex)

	p := &Properties{
		mu:   mu,
		body: body,
	}

	return p, nil
}

func (p *Properties) Get(ctx context.Context, path string) (*gjson.Result, error) {

	p.mu.RLock()
	defer p.mu.RUnlock()

	return GetProperty(ctx, p.body, path)
}

func (p *Properties) Set(ctx context.Context, body []byte, path string, value interface{}) error {

	p.mu.Lock()
	defer p.mu.Unlock()

	new_body, err := SetProperty(ctx, p.body, path, value)

	if err != nil {
		return err
	}

	p.body = new_body
	return nil
}

func (p *Properties) Body() []byte {
	return p.body
}

func GetProperty(ctx context.Context, body []byte, path string) (*gjson.Result, error) {

	rsp := gjson.GetBytes(body, path)

	if !rsp.Exists() {
		return nil, &NotFoundError{path}
	}

	return &rsp, nil
}

func SetProperty(ctx context.Context, body []byte, path string, value interface{}) ([]byte, error) {
	return sjson.SetBytes(body, path, value)
}

func SetProperties(ctx context.Context, body []byte, properties map[string]interface{}) ([]byte, error) {

	var err error

	for path, v := range properties {

		body, err = SetProperty(ctx, body, path, v)

		if err != nil {
			return nil, err
		}
	}

	return body, nil
}
