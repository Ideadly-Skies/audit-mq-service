# Setup Instructions for Windows Machine

## Prerequisites

- IBM MQ Explorer installed (which you already have)
- Go installed on Windows

## Steps:

### 1. Copy this entire folder to Windows

Copy to: `C:\Users\OA01762X\Documents\audit-mq-service`

### 2. Set Environment Variables (PowerShell)

```powershell
$env:MQ_INSTALLATION_PATH="C:\Program Files\IBM\MQ"
$env:PATH="$env:MQ_INSTALLATION_PATH\bin;$env:PATH"
$env:CGO_CFLAGS="-I$env:MQ_INSTALLATION_PATH\tools\c\include"
$env:CGO_LDFLAGS="-L$env:MQ_INSTALLATION_PATH\bin64 -lmqm"
```

### 3. Install Dependencies

```bash
cd C:\Users\OA01762X\Documents\audit-mq-service
go mod download
go mod tidy
```

### 4. Run Subscriber (Terminal 1)

```bash
go run . subscriber
```

### 5. Run Publisher (Terminal 2)

```bash
go run . publisher
```

## Troubleshooting

### If channel name is wrong:

Open [config.go](config.go) and update the `Channel` constant to match your MQ server's channel name (check in IBM MQ Explorer).

### If queue doesn't exist:

Verify `ESB.RND.OTEL.QA` exists in IBM MQ Explorer under your queue manager.

### Connection timeout:

Check firewall allows connection to 10.25.135.209:1415
