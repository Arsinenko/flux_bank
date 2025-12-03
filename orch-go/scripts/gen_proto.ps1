# Директория с proto файлами
$PROTO_DIR = "api/proto"
# Директория для сгенерированных файлов
$GENERATED_DIR = "api/generated"

# Создаем директорию, если она не существует
if (-not (Test-Path $GENERATED_DIR)) {
    New-Item -ItemType Directory -Force -Path $GENERATED_DIR | Out-Null
}

# Получаем относительные пути proto файлов
$protoFiles = Get-ChildItem -Path $PROTO_DIR -Filter *.proto -Recurse |
        ForEach-Object {
            $full = $_.FullName
            $root = (Resolve-Path $PROTO_DIR).ToString()
            $rel = $full.Substring($root.Length).TrimStart("\", "/")
            $rel
        }

if ($protoFiles.Count -eq 0) {
    Write-Host "Нет .proto файлов в каталоге $PROTO_DIR"
    exit 1
}

# Запускаем protoc
protoc @(
    "--proto_path=$PROTO_DIR"
    "--go_out=$GENERATED_DIR"
    "--go_opt=paths=source_relative"
    "--go-grpc_out=$GENERATED_DIR"
    "--go-grpc_opt=paths=source_relative"
    $protoFiles
)

# Проверяем результат
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Ошибка генерации proto файлов!"
    exit $LASTEXITCODE
}

Write-Host "✅ Proto файлы успешно сгенерированы в $GENERATED_DIR"
