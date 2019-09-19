#!/usr/bin/env bash

protoc -I . ./tracker.proto --go_out=plugins=grpc:./