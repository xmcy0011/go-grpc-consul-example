#!/bin/sh

protoc --go_out=. --go-grpc_out=. api/user.proto