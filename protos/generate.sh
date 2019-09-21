#!/usr/bin/env bash

protoc -I . ./*.proto --go_out=plugins=grpc:./

protoc -I ./tracker ./tracker/*.proto --go_out=plugins=grpc:./tracker