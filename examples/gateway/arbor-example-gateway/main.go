package main

import (
	"github.com/acm-uiuc/arbor"
	"github.com/acm-uiuc/arbor/examples/gateway"
)

func main() {
	//Configure Arbor
	gateway.ConfigArbor()
	//Register the Routes in a Collection and Boot Arbor
	arbor.Boot(gateway.RegisterRoutes(), 8000)
}
