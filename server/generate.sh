#!/bin/bash

set -e

# Install oapi-codegen if not already installed
echo "Installing oapi-codegen..."
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

# Use the full path to the binary
GOPATH=$(go env GOPATH)
OAPI_CODEGEN="${GOPATH}/bin/oapi-codegen"

# Generate server code
echo "Generating server code from OpenAPI spec..."
$OAPI_CODEGEN -package api -generate types,server,spec -o internal/api/api.gen.go api/openapi.yaml

echo "Code generation complete!"
