package arbor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/acm-uiuc/arbor"
	"github.com/acm-uiuc/arbor/examples/gateway"
	"github.com/acm-uiuc/arbor/examples/products"
	"github.com/acm-uiuc/arbor/server"
)

const url string = "http://127.0.0.1:8000"

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
	t.testGateway = arbor.Boot(gateway.RegisterRoutes(), "127.0.0.1", 8000)
	return t
}

func (t *testingServices) killTestingServices() {
	t.testGateway.KillServer()
	t.testService.Kill()
}

func TestProxyGETEmpty(t *testing.T) {
	os.Args = []string{"tests", "-u"}

	testSrvs := newTestingServices()

	c := http.Client{Timeout: time.Second * 2}

	req, err := http.NewRequest(http.MethodGet, url+"/products", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := c.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.StatusCode != http.StatusOK {
		t.Error(
			"For", res,
			"expected", http.StatusOK,
			"got", res.StatusCode,
		)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	p := make([]product, 0)
	jsonErr := json.Unmarshal(body, &p)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	if len(p) != 0 {
		t.Error(
			"For", p,
			"expected", 0,
			"got", len(p),
		)
	}

	testSrvs.killTestingServices()
}
func TestProxyPOST(t *testing.T) {
	os.Args = []string{"tests", "-u"}
	testSrvs := newTestingServices()

	p := new(product)
	p.ID = 0
	p.Name = "Test Product"
	p.Price = 9.99

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(p)

	res, err := http.Post(url+"/product", "application/json; charset=utf-8", b)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error(
			"For", res,
			"expected", http.StatusOK,
			"got", res.StatusCode,
		)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		testSrvs.killTestingServices()
		t.Error(readErr)
		return
	}

	serverP := new(product)
	jsonErr := json.Unmarshal(body, &serverP)
	if jsonErr != nil {
		testSrvs.killTestingServices()
		t.Error(readErr)
		return
	}

	if !reflect.DeepEqual(p, serverP) {
		t.Error(
			"For", serverP,
			"expected", p,
			"got", serverP,
		)
	}

	testSrvs.killTestingServices()
}

func TestProxyGET(t *testing.T) {
	os.Args = []string{"tests", "-u"}
	testSrvs := newTestingServices()

	p := new(product)
	p.ID = 0
	p.Name = "Test Product"
	p.Price = 9.99

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(p)

	res, err := http.Post(url+"/product", "application/json; charset=utf-8", b)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error(
			"For", res,
			"expected", http.StatusOK,
			"got", res.StatusCode,
		)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		testSrvs.killTestingServices()
		t.Error(readErr)
		return
	}

	serverP := new(product)
	jsonErr := json.Unmarshal(body, &serverP)
	if jsonErr != nil {
		testSrvs.killTestingServices()
		t.Error(readErr)
		return
	}

	if !reflect.DeepEqual(p, serverP) {
		t.Error(
			"For", serverP,
			"expected", p,
			"got", serverP,
		)
	}

	c := http.Client{Timeout: time.Second * 2}

	req, err := http.NewRequest(http.MethodGet, url+"/products/0", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := c.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatal("Gateway returned " + http.StatusText(res.StatusCode))
	}

	body, readErr = ioutil.ReadAll(res.Body)
	if readErr != nil {
		testSrvs.killTestingServices()
		t.Error(readErr)
		return
	}

	serverP = new(product)
	jsonErr = json.Unmarshal(body, &serverP)
	if jsonErr != nil {
		testSrvs.killTestingServices()
		t.Error(readErr)
		return
	}

	if !reflect.DeepEqual(p, serverP) {
		t.Error(
			"For", serverP,
			"expected", p,
			"got", serverP,
		)
	}

	testSrvs.killTestingServices()
}

func TestProxyPUT(t *testing.T) {
	os.Args = []string{"tests", "-u"}
	testSrvs := newTestingServices()

	p := new(product)
	p.ID = 0
	p.Name = "Test Product"
	p.Price = 9.99

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(p)

	res, err := http.Post(url+"/product", "application/json; charset=utf-8", b)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error(
			"For", res,
			"expected", http.StatusOK,
			"got", res.StatusCode,
		)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		testSrvs.killTestingServices()
		t.Error(readErr)
		return
	}

	serverP := new(product)
	jsonErr := json.Unmarshal(body, &serverP)
	if jsonErr != nil {
		testSrvs.killTestingServices()
		t.Error(readErr)
		return
	}

	if !reflect.DeepEqual(p, serverP) {
		t.Error(
			"For", serverP,
			"expected", p,
			"got", serverP,
		)
	}

	c := http.Client{Timeout: time.Second * 2}

	p = new(product)
	p.ID = 0
	p.Name = "Test Product"
	p.Price = 10.99

	b = new(bytes.Buffer)
	json.NewEncoder(b).Encode(p)

	req, err := http.NewRequest(http.MethodPut, url+"/products/0", b)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := c.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatal("Gateway returned " + http.StatusText(res.StatusCode))
	}

	body, readErr = ioutil.ReadAll(res.Body)
	if readErr != nil {
		testSrvs.killTestingServices()
		t.Error(readErr)
		return
	}

	serverP = new(product)
	jsonErr = json.Unmarshal(body, &serverP)
	if jsonErr != nil {
		testSrvs.killTestingServices()
		t.Error(readErr)
		return
	}

	if !reflect.DeepEqual(p, serverP) {
		t.Error(
			"For", serverP,
			"expected", p,
			"got", serverP,
		)
	}

	testSrvs.killTestingServices()

}

func TestProxyDELETE(t *testing.T) {
	os.Args = []string{"tests", "-u"}
	testSrvs := newTestingServices()

	p := new(product)
	p.ID = 0
	p.Name = "Test Product"
	p.Price = 9.99

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(p)

	res, err := http.Post(url+"/product", "application/json; charset=utf-8", b)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error(
			"For", res,
			"expected", http.StatusOK,
			"got", res.StatusCode,
		)
	}

	c := http.Client{Timeout: time.Second * 2}

	req, err := http.NewRequest(http.MethodDelete, url+"/products/0", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := c.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatal("Gateway returned " + http.StatusText(res.StatusCode))
	}

	c = http.Client{Timeout: time.Second * 2}

	req, err = http.NewRequest(http.MethodGet, url+"/products", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr = c.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.StatusCode != http.StatusOK {
		t.Error(
			"For", res,
			"expected", http.StatusOK,
			"got", res.StatusCode,
		)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	srvp := make([]product, 0)
	jsonErr := json.Unmarshal(body, &srvp)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	if len(srvp) != 0 {
		t.Error(
			"For", srvp,
			"expected", 0,
			"got", len(srvp),
		)
	}

	testSrvs.killTestingServices()
}
