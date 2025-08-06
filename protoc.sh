#!/bin/bash

protoc --go_out=. --proto_path=. proto/*.proto

rm -rf pkg/gen/proto
mv srv-eazle-advise-mock/proto pkg/gen/proto
rm -rf srv-eazle-advise-mock
