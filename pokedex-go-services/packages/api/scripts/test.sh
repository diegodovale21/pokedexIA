#!/bin/bash

# Test script for PokedexIA Backend

echo "🧪 Running tests for PokedexIA Backend..."
echo "=========================================="

# Run tests with coverage
echo "📊 Running tests with coverage..."
go test -v -coverprofile=coverage.out ./...

# Show coverage report
echo ""
echo "📈 Coverage Report:"
go tool cover -func=coverage.out

# Show coverage in HTML (optional)
echo ""
echo "🌐 Generating HTML coverage report..."
go tool cover -html=coverage.out -o coverage.html

echo ""
echo "✅ Tests completed!"
echo "📄 HTML coverage report: coverage.html" 