#!/usr/bin/env bash

set -e

PROTO_DIR="../api/proto"
OUT_DIR="../api/generated"

echo "Generating gRPC Python files..."
uv run -m grpc_tools.protoc \
  -I ${PROTO_DIR} \
  --python_out=${OUT_DIR} \
  --grpc_python_out=${OUT_DIR} \
  --mypy_out=${OUT_DIR} \
  ${PROTO_DIR}/*.proto

echo "Done."
