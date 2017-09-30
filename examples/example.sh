#!/bin/bash

#Build Service and Gateway
echo "Make sure \$GOPATH/bin is in your PATH"

echo "Building Gateway"
go install github.com/acm-uiuc/arbor/examples/arbor-example-gateway

echo "Building Service"
go install github.com/acm-uiuc/arbor/examples/product-service

trap 'kill %1; kill %2;' SIGINT

product-service \
& arbor-example-gateway "$@"

