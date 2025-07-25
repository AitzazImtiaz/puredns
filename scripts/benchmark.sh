#!/bin/bash
echo "Running benchmarks..."
go test -bench=. ./...
echo "Benchmark complete."
