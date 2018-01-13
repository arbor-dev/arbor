package tests

import (
	"time"

	"github.com/acm-uiuc/arbor"
	"github.com/acm-uiuc/arbor/examples/gateway"
	"github.com/acm-uiuc/arbor/examples/products"
	"github.com/acm-uiuc/arbor/server"
)

const url string = "http://0.0.0.0:8000"

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type testingServices struct {
	testGateway *server.ArborServer
	testService *products.App
}

func newTestingServices() *testingServices {
	t := new(testingServices)
	t.testService = products.NewApp()
	t.testService.Run()
	gateway.ConfigArbor()
	t.testGateway = arbor.Boot(gateway.RegisterRoutes(), "0.0.0.0", 8000)
	time.Sleep(250 * time.Millisecond)
	return t
}

func (t *testingServices) killTestingServices() {
	t.testGateway.KillServer()
	t.testService.Kill()
	time.Sleep(250 * time.Millisecond)
}
