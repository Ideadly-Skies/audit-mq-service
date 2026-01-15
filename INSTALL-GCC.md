# Quick Fix: Install GCC on Windows

You're getting the error because Go needs a C compiler (gcc) to compile the IBM MQ library.

## Easiest Solution: Install TDM-GCC

### 1. Download TDM-GCC

- Visit: https://jmeubank.github.io/tdm-gcc/download/
- Download: **tdm64-gcc-10.3.0-2.exe** (or latest version)

### 2. Install

- Run the installer
- Choose **"Create"** (new installation)
- Select **"MinGW-w64/TDM64 (32-bit and 64-bit)"**
- Click **Install**
- Default installation path is fine: `C:\TDM-GCC-64`

### 3. Verify Installation

Open a **NEW** Command Prompt and run:

```cmd
gcc --version
```

Should show something like: `gcc (tdm64-1) 10.3.0`

### 4. Run the setup script again

```cmd
cd C:\Users\OA01762X\Documents\New Project - Archie\audit-mq-service
setup-windows.bat
```

### 5. Build and run

```cmd
go mod tidy
go run . subscriber
```

---

## Alternative: If you have Chocolatey installed

```powershell
choco install mingw
```

## Alternative: If you have Git for Windows

Git for Windows includes gcc. Just make sure `C:\Program Files\Git\mingw64\bin` is in your PATH.

---

After installing gcc, close and reopen your terminal, then try the setup script again!
