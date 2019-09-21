#! /bin/bash

TRACKER_PORT=50050

for i in {1..4}
do
    go run cmd/server/main.go localhost $TRACKER_PORT "${i}${i}" &
    sleep 0.9
done