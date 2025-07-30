# graphlib
[![Go Report Card](https://goreportcard.com/badge/github.com/aio-arch/graphlib)](https://goreportcard.com/report/github.com/aio-arch/graphlib)
[![Codecov](https://img.shields.io/codecov/c/github/aio-arch/graphlib?style=flat-square&logo=codecov)](https://app.codecov.io/gh/aio-arch/graphlib)
[![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/aio-arch/graphlib/go.yml)](https://github.com/aio-arch/graphlib/actions)
![Minimum Go Version](https://img.shields.io/badge/go-%3E%3D1.20-30dff3?style=flat-square&logo=go)

A Topological sort lib.

Sorting and pruning of DAG graphs.

Ideas borrowed from [python graphlib](https://github.com/python/cpython/blob/3.14/Lib/graphlib.py)

# For Go Veriosn < 1.21
```bash
go mod edit -replace golang.org/x/exp=golang.org/x/exp@v0.0.0-20240904232852-e7e105dedf7e
```
