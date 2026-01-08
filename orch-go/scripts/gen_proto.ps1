# Добавляем путь к Go бинарникам в PATH
$env:PATH += ";C:\Users\Arsinenko\go\bin"

# Директория с proto файлами (относительно скрипта)
$PSScriptRoot = Split-Path -Parent $MyInvocation.MyCommand.Path
$PROTO_DIR = Join-Path $PSScriptRoot "../../proto"
$PROTO_DIR = [System.IO.Path]::GetFullPath($PROTO_DIR)

# Директория для сгенерированных файлов (относительно текущей рабочей директории, обычно корень проекта)
$GENERATED_DIR = Join-Path (Get-Location) "api/generated"
$GENERATED_DIR = [System.IO.Path]::GetFullPath($GENERATED_DIR)

# Создаем директорию, если она не существует
if (-not (Test-Path $GENERATED_DIR)) {
    New-Item -ItemType Directory -Force -Path $GENERATED_DIR | Out-Null
}

# Получаем относительные пути proto файлов
$protoFiles = Get-ChildItem -Path $PROTO_DIR -Filter *.proto | ForEach-Object { $_.Name }

if ($protoFiles.Count -eq 0) {
    Write-Host "Нет .proto файлов в каталоге $PROTO_DIR"
    exit 1
}

Write-Host "Генерируем Go файлы из $PROTO_DIR в $GENERATED_DIR..."

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
