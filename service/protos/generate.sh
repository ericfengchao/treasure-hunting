#!/usr/bin/env bash

protoc -I . ./player.proto --go_out=plugins=grpc:./
