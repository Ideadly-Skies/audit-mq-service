@echo off
echo Setting up IBM MQ Environment for Go...

REM Set IBM MQ installation path (adjust if your installation is different)
set MQ_INSTALLATION_PATH=C:\Program Files\IBM\MQ

REM Check for GCC
where gcc >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo.
    echo ERROR: GCC compiler not found!
    echo You need to install a C compiler. Choose one option:
    echo.
    echo Option 1 - Install TDM-GCC (easiest):
    echo   Download from: https://jmeubank.github.io/tdm-gcc/download/
    echo   Install and restart this script
    echo.
    echo Option 2 - Install via Chocolatey:
    echo   choco install mingw
    echo.
    echo Option 3 - Install Git for Windows (includes gcc):
    echo   https://git-scm.com/download/win
    echo.
    pause
    exit /b 1
)

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
echo GCC found: 
gcc --version | findstr gcc
echo.
echo Now run: go mod tidy
echo Then run: go run . publisher
echo.

cmd /k
