#!/bin/bash

#Build Service and Gateway
echo "Make sure \$GOPATH/bin is in your PATH"

echo "Building Gateway"
go install github.com/arbor-dev/arbor/examples/gateway/arbor-example-gateway

echo "Building Service"
go install github.com/arbor-dev/arbor/examples/products/products-service

trap 'kill %1; kill %2;' SIGINT

products-service \
& arbor-example-gateway "$@"

