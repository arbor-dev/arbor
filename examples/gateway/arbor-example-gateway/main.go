package main

import (
	"github.com/arbor-dev/arbor"
	"github.com/arbor-dev/arbor/examples/gateway"
)

func main() {
	//Configure Arbor
	gateway.ConfigArbor()
	//Register the Routes in a Collection and Boot Arbor
	arbor.Boot(gateway.RegisterRoutes(), "0.0.0.0", 8000)
}
