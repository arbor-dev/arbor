// main.go

package main

import (
	"github.com/arbor-dev/arbor/examples/products"
)

func main() {
	a := products.NewApp()
	a.Run()
}
