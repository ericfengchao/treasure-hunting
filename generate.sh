#!/usr/bin/env bash

#protoc -I ./tracker ./tracker/*.proto --go_out=plugins=grpc:./tracker

protoc -I . ./*.proto --go_out=plugins:./
