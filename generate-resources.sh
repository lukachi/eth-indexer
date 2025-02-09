#!/bin/bash
set -e

# Check if oapi-codegen is installed.
if ! command -v oapi-codegen >/dev/null 2>&1; then
    echo "Error: oapi-codegen is not installed."
    echo "Please run:"
    echo "  go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest"
    exit 1
fi

# Check if npx is installed (for Redocly bundling).
if ! command -v npx >/dev/null 2>&1; then
    echo "Error: npx is not installed. Please install Node.js and npm."
    exit 1
fi

# Define paths.
#DOCS_MASTER="http://localhost:3000/api"
OUTPUT_YAML="./docs/tsp-output/@typespec/openapi3/openapi.yaml"
OUTPUT_GO="./resources/generated.go"
OAPI_CONFIG="./oapi-codegen-config.yml"

# Ensure the resources folder exists.
mkdir -p ./resources

## Bundle the OpenAPI spec using Redocly's bundle command.
#npx @redocly/cli bundle "$DOCS_MASTER" -o "$OUTPUT_YAML"
#echo "Bundled OpenAPI spec to $OUTPUT_YAML"

# Generate Go resources using oapi-codegen.
oapi-codegen --config="$OAPI_CONFIG" "$OUTPUT_YAML" > "$OUTPUT_GO"
echo "Generated Go code to $OUTPUT_GO"
