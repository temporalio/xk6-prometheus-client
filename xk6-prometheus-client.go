package promql

import (
	"context"
	"time"

	"go.k6.io/k6/js/modules"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

func init() {
	modules.Register("k6/x/prometheus-client", new(RootModule))
}

// RootModule is the global module object type. It is instantiated once per test
// run and will be used to create `k6/x/temporal` module instances for each VU.
type RootModule struct{}

// ModuleInstance represents an instance of the module for every VU.
type ModuleInstance struct{}

// Ensure the interfaces are implemented correctly.
var (
	_ modules.Module   = &RootModule{}
	_ modules.Instance = &ModuleInstance{}
)

// NewModuleInstance implements the modules.Module interface to return
// a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &ModuleInstance{}
}

// Exports implements the modules.Instance interface and returns the exports
// of the JS module.
func (m *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{Default: m}
}

type Client struct {
	api v1.API
}

func (c *Client) Query(query string, ts time.Time) (model.Value, v1.Warnings, error) {
	return c.api.Query(context.Background(), query, ts)
}

func (c *Client) QueryRange(query string, r v1.Range) (model.Value, v1.Warnings, error) {
	return c.api.QueryRange(context.Background(), query, r)
}

// NewClient returns a new prometheus API Client.
func (m *ModuleInstance) NewClient(address string) (*Client, error) {
	client, err := api.NewClient(api.Config{Address: address})
	if err != nil {
		return nil, err
	}

	return &Client{api: v1.NewAPI(client)}, nil
}
