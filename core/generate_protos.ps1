# PowerShell script to generate C# code from proto files using protoc
$PROTO_DIR = "..\proto"
$OUT_DIR = "Core\Generated"

if (-not (Test-Path $OUT_DIR)) {
    New-Item -ItemType Directory -Force -Path $OUT_DIR | Out-Null
}

Write-Host "Generating C# files from proto..."

# Find all proto files in the central directory
$protoFiles = Get-ChildItem -Path $PROTO_DIR -Filter *.proto | ForEach-Object { $_.FullName }

# Run protoc
# Note: This assumes protoc is in the PATH and Google.Protobuf/Grpc tools are available
# Usually in .NET projects, Grpc.Tools handles this automatically during build, 
# but this script is for manual generation as requested.
protoc --proto_path=$PROTO_DIR `
    --csharp_out=$OUT_DIR `
    --grpc_out=$OUT_DIR `
    --plugin=protoc-gen-grpc=C:\Users\Arsinenko\.nuget\packages\grpc.tools\2.76.0\tools\windows_x64\grpc_csharp_plugin.exe `
    $protoFiles

Write-Host "✅ C# files successfully generated in $OUT_DIR"
