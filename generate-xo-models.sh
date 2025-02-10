#!/bin/bash
set -e

# Check if .env exists
if [ ! -f .env ]; then
  echo "Error: .env file not found."
  exit 1
fi

# Load environment variables from .env (ignoring comment lines)
export $(grep -v '^#' .env | xargs)

# Build the PostgreSQL connection string from env variables
CONN="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"
echo "Using connection string: $CONN"

# Ensure the output directory exists
mkdir -p internal/db/models

# Run xo to generate models
# - The first argument is the connection string
# - -o internal/db/models tells xo to output the generated files in internal/db/models folder
# - --pkg models sets the package name in the generated Go code to "models"
xo schema "$CONN" -o internal/db/models

echo "Models generated in internal/db/models folder."
