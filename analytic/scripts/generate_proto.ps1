$PSScriptRoot = Split-Path -Parent $MyInvocation.MyCommand.Path
$ProtoDir = Join-Path $PSScriptRoot "../../proto"
$ProtoDir = [System.IO.Path]::GetFullPath($ProtoDir)

$OutDir = Join-Path $PSScriptRoot "../api/generated"
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



Start-Sleep -Seconds 2
Write-Host "Done."

# Create __init__.py to make it a package
$InitFile = Join-Path $OutDir "__init__.py"
if (-not (Test-Path $InitFile)) {
    New-Item -ItemType File -Force -Path $InitFile | Out-Null
    Write-Host "Created __init__.py"
}

# Fix imports in generated files (Python 3 absolute/relative import issue)
Write-Host "Fixing imports in generated files..."
Get-ChildItem -Path $OutDir -Filter "*_pb2*.py" | ForEach-Object {
    $content = Get-Content -Path $_.FullName -Raw
    
    # Replace 'import x_pb2 as' with 'from . import x_pb2 as'
    # This matches imports of sibling generated files
    $content = $content -replace '(?m)^import (.*_pb2) as', 'from . import $1 as'
    
    # Replace 'import x_pb2' with 'from . import x_pb2' (if any without alias)
    $content = $content -replace '(?m)^import (.*_pb2)$', 'from . import $1'

    Set-Content -Path $_.FullName -Value $content -Encoding UTF8
}

Write-Host "Imports fixed."
