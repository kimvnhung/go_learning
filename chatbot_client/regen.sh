#!/bin/bash

# read path to proto folder from $1
# if not provided, use default path
PROTO_PATH=${1:-"src/proto"}

# then add to below command

echo "Using protoc version: $(protoc --version)"

protoc -I ${PROTO_PATH} \
    --go_out ${PROTO_PATH} --go_opt paths=source_relative \
    --go-grpc_out ${PROTO_PATH} --go-grpc_opt paths=source_relative \
    ${PROTO_PATH}/protogenerated/*.proto

echo "Generating gateway files"

protoc -I ${PROTO_PATH} --grpc-gateway_out ${PROTO_PATH} \
    --grpc-gateway_opt allow_delete_body=true \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    ${PROTO_PATH}/protogenerated/*.proto