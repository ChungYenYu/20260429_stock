#!/bin/bash
set -e

PROJECT_ROOT="/home/cary_yu/20260429_stock"

echo "Running backend tests..."
cd "$PROJECT_ROOT/backend"
go test ./...

echo "Running frontend unit tests..."
cd "$PROJECT_ROOT/frontend"
npm run test:unit

echo "Starting the app via docker-compose..."
cd "$PROJECT_ROOT"
docker-compose up --build -d

echo "Waiting for frontend to be ready..."
# Use wait-on from the root package.json
./node_modules/.bin/wait-on http://localhost:3000 -t 60000

echo "Running E2E tests..."
npx playwright test

EXIT_CODE=$?

echo "Stopping the app..."
docker-compose down

if [ $EXIT_CODE -eq 0 ]; then
  echo "All tests passed!"
  exit 0
else
  echo "E2E tests failed!"
  exit 1
fi
