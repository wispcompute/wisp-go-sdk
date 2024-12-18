#!/bin/bash

# TODO: Update with new endpoint when Elias is done setting up public endpoint
# Define variables
OPENAPI_URL="https://raw.githubusercontent.com/wispcompute/wisp/refs/heads/dev/web/openapi/latest.yml"
OUTPUT_DIR="./go-sdk"
DOCKER_IMAGE="openapitools/openapi-generator-cli:v5.3.0"  # OpenAPI Generator Docker image

# Download OpenAPI schema from URL
echo "Downloading OpenAPI schema..."
curl -L -o openapi-schema.yml $OPENAPI_URL

if [ $? -ne 0 ]; then
    echo "Failed to download OpenAPI schema."
    exit 1
fi

echo "OpenAPI schema downloaded successfully."

# Create output directory
mkdir -p $OUTPUT_DIR

# Generate Go SDK using OpenAPI Generator via Docker
echo "Generating Go SDK using Docker..."
docker run --rm \
  -v "$(pwd)":/local \
  -w /local \
  $DOCKER_IMAGE generate \
  -i openapi-schema.yml \
  -g go \
  -o $OUTPUT_DIR \
  --additional-properties=packageName=wisp-sdk

if [ $? -ne 0 ]; then
    echo "Failed to generate Go SDK using Docker."
    exit 1
fi

echo "Go SDK generated successfully in directory: $OUTPUT_DIR"

# Clean up
rm openapi-schema.yml

echo "Script completed."
