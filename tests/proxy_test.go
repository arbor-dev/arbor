package arbor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"strconv"
	"time"

	"gopkg.in/jarcoal/httpmock.v1"

	"github.com/arbor-dev/arbor/examples/gateway"
	"github.com/arbor-dev/arbor/examples/products"
	"github.com/arbor-dev/arbor/proxy"
	"github.com/arbor-dev/arbor/server"
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
	gateway.ConfigArbor()
	t := new(testingServices)
	t.testService = products.NewApp()
	t.testGateway = server.NewArborServer(gateway.RegisterRoutes().ToServiceRoutes(), "0.0.0.0", 8000)
	go t.testService.Run()
	go t.testGateway.StartServer()
	time.Sleep(250 * time.Millisecond)
	return t
}

func (t *testingServices) killTestingServices() {
	t.testGateway.KillServer()
	t.testService.Kill()
	time.Sleep(250 * time.Millisecond)
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

func TestProxyPUTRaw(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	sendingBytes := make([]byte, 8 * 1024 * 1024)
	rand.Read(sendingBytes)

	httpmock.RegisterResponder("PUT", "http://test.local/upload",
		func (req *http.Request) (*http.Response, error) {
			receivedBytes, err := ioutil.ReadAll(req.Body)

			if err != nil {
				log.Fatal(err)
			}

			if !bytes.Equal(sendingBytes, receivedBytes) {
				t.Error("Received incorrect bytes")
			}

			return httpmock.NewStringResponse(200, strconv.Itoa(len(receivedBytes))), nil
		},
	)

	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "http://test.local/upload", bytes.NewReader(sendingBytes))

	if err != nil {
		log.Fatal(err)
	}

	proxy.PUT(recorder, "http://test.local/upload", "RAW", "", req)

	resp := recorder.Result()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Received bad http code")
	}

	body, err := ioutil.ReadAll(recorder.Body)

	if err != nil {
		log.Fatal(err)
	}

	if strconv.Itoa(len(sendingBytes)) != string(body[:]) {
		t.Errorf("Expected %v\nGot %v", strconv.Itoa(len(sendingBytes)), string(body[:]))
	}
}
