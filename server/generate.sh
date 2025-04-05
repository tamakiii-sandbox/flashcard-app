#!/bin/bash

set -e

# Install OpenAPI Generator if not already installed
echo "Installing OpenAPI Generator..."
# Check if Java is installed, as OpenAPI Generator requires Java
if ! command -v java &> /dev/null; then
    echo "Error: Java is required for OpenAPI Generator but it's not installed."
    echo "Please install Java and try again."
    exit 1
fi

# Using openapi-generator-cli via Maven
if ! command -v openapi-generator &> /dev/null; then
    echo "OpenAPI Generator not found, installing..."
    # Install with Homebrew, or provide instructions for manual installation
    if command -v brew &> /dev/null; then
        brew install openapi-generator
    else
        echo "Please install OpenAPI Generator manually:"
        echo "Visit: https://github.com/OpenAPITools/openapi-generator#1---installation"
        echo "For example, you can download the JAR file:"
        echo "curl -O https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/6.0.1/openapi-generator-cli-6.0.1.jar"
        echo "mv openapi-generator-cli-6.0.1.jar openapi-generator-cli.jar"
        echo "chmod +x openapi-generator-cli.jar"
        exit 1
    fi
fi

# Path to the OpenAPI spec
SPEC_FILE="internal/api/openapi.yaml"
OUTPUT_DIR="internal/api/openapi"

# Generate server code
echo "Generating server code from OpenAPI spec..."
openapi-generator generate \
    -i ${SPEC_FILE} \
    -g go-server \
    -o ${OUTPUT_DIR} \
    --additional-properties=packageName=api,serverPort=8080,sourceFolder=

# Remove unnecessary files (optional)
rm -f ${OUTPUT_DIR}/go.mod ${OUTPUT_DIR}/go.sum ${OUTPUT_DIR}/main.go 2>/dev/null || true

echo "Code generation complete!"
echo "Note: You may need to update your imports and implementation to match the newly generated code."
