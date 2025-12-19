$ProtoDir = "api\proto"
$OutDir = "api\generated"

Write-Host "Generating gRPC Python files..."

uv run -m grpc_tools.protoc `
    -I $ProtoDir `
    --python_out=$OutDir `
    --grpc_python_out=$OutDir `
    --mypy_out=$OutDir `
    "$ProtoDir\*.proto"

Write-Host "Done."
