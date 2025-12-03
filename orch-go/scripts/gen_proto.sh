#!/bin/bash

# Директория с proto файлами
PROTO_DIR="api/proto"
# Директория для сгенерированных файлов
GENERATED_DIR="api/generated"

# Создаем директорию, если она не существует
mkdir -p ${GENERATED_DIR}

# Генерируем Go файлы из proto
protoc --proto_path=${PROTO_DIR} \
       --go_out=${GENERATED_DIR} --go_opt=paths=source_relative \
       --go-grpc_out=${GENERATED_DIR} --go-grpc_opt=paths=source_relative \
       $(find ${PROTO_DIR} -name *.proto)

echo "Proto файлы успешно сгенерированы в ${GENERATED_DIR}"
