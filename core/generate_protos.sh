#!/bin/bash
PROTO_DIR="../proto"
OUT_DIR="Core/Generated"

mkdir -p ${OUT_DIR}

echo "Generating C# files from proto..."

protoc --proto_path=${PROTO_DIR} \
       --csharp_out=${OUT_DIR} \
       --grpc_out=${OUT_DIR} \
       ${PROTO_DIR}/*.proto

echo "Done."
