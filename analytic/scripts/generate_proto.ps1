$PSScriptRoot = Split-Path -Parent $MyInvocation.MyCommand.Path
$ProtoDir = Join-Path $PSScriptRoot "../../proto"
$ProtoDir = [System.IO.Path]::GetFullPath($ProtoDir)

$OutDir = Join-Path (Get-Location) "api/generated"
$OutDir = [System.IO.Path]::GetFullPath($OutDir)

if (-not (Test-Path $OutDir)) {
    New-Item -ItemType Directory -Force -Path $OutDir | Out-Null
}

Write-Host "Generating gRPC Python files from $ProtoDir..."

# Get just filenames
$protoFiles = Get-ChildItem -Path $ProtoDir -Filter *.proto | ForEach-Object { $_.Name }

uv run -m grpc_tools.protoc `
    -I "$ProtoDir" `
    --python_out="$OutDir" `
    --grpc_python_out="$OutDir" `
    --mypy_out="$OutDir" `
    $protoFiles

Write-Host "Done."
