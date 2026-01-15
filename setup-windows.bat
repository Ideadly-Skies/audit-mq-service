@echo off
echo Setting up IBM MQ Environment for Go...

REM Set IBM MQ installation path (adjust if your installation is different)
set MQ_INSTALLATION_PATH=C:\Program Files\IBM\MQ

REM Add MQ bin to PATH
set PATH=%MQ_INSTALLATION_PATH%\bin;%MQ_INSTALLATION_PATH%\bin64;%PATH%

REM Enable CGO (required for IBM MQ library)
set CGO_ENABLED=1

REM Set CGO flags to find IBM MQ headers and libraries
set CGO_CFLAGS=-I"%MQ_INSTALLATION_PATH%\tools\c\include"
set CGO_LDFLAGS=-L"%MQ_INSTALLATION_PATH%\bin64" -lmqm

echo.
echo Environment variables set:
echo MQ_INSTALLATION_PATH=%MQ_INSTALLATION_PATH%
echo CGO_ENABLED=%CGO_ENABLED%
echo.
echo Now run: go mod tidy
echo Then run: go run . publisher
echo.

cmd /k
