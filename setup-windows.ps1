# Setup Environment for IBM MQ on Windows
Write-Host "Setting up IBM MQ Environment for Go..." -ForegroundColor Green

# Set IBM MQ installation path (adjust if your installation is different)
$env:MQ_INSTALLATION_PATH = "C:\Program Files\IBM\MQ"

# Check if MQ is installed
if (-not (Test-Path $env:MQ_INSTALLATION_PATH)) {
    Write-Host "ERROR: IBM MQ not found at $env:MQ_INSTALLATION_PATH" -ForegroundColor Red
    Write-Host "Please update the path in this script to match your IBM MQ installation" -ForegroundColor Yellow
    exit 1
}

# Add MQ bin to PATH
$env:PATH = "$env:MQ_INSTALLATION_PATH\bin;$env:MQ_INSTALLATION_PATH\bin64;$env:PATH"

# Enable CGO (required for IBM MQ library)
$env:CGO_ENABLED = "1"

# Set CGO flags to find IBM MQ headers and libraries
$env:CGO_CFLAGS = "-I`"$env:MQ_INSTALLATION_PATH\tools\c\include`""
$env:CGO_LDFLAGS = "-L`"$env:MQ_INSTALLATION_PATH\bin64`" -lmqm"

Write-Host ""
Write-Host "Environment variables set:" -ForegroundColor Green
Write-Host "  MQ_INSTALLATION_PATH = $env:MQ_INSTALLATION_PATH"
Write-Host "  CGO_ENABLED = $env:CGO_ENABLED"
Write-Host ""
Write-Host "Now run the following commands:" -ForegroundColor Yellow
Write-Host "  1. go mod tidy"
Write-Host "  2. go run . subscriber   (in one terminal)"
Write-Host "  3. go run . publisher    (in another terminal)"
Write-Host ""
