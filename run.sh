#!/usr/bin/env bash

go clean && go clean -modcache && go mod download && go run httpServer.go