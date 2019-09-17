#!/usr/bin/env bash

protoc -I . ./game.proto --go_out=plugins=grpc:./
