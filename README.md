# Audit MQ Service - Demo Version

Publisher-subscriber pattern for IBM MQ queue: **ESB.RND.OTEL.QA**

## Configuration

- **Queue Manager**: SIDPOP01
- **Host**: 10.25.135.209
- **Port**: 1415
- **Queue**: ESB.RND.OTEL.QA

## Current Version: In-Memory Demo

This is a **simplified in-memory version** that works on Windows without requiring IBM MQ Client libraries. Perfect for:

- Understanding publisher-subscriber patterns
- Testing your Go setup
- Demonstrating the message flow

## Quick Start

### 1. Run Subscriber (Terminal 1)

```bash
go run . subscriber
```

### 2. Run Publisher (Terminal 2)

```bash
go run . publisher
```

The publisher sends 5 mock audit messages, subscriber receives them in real-time! 🚀

## Expected Output

**Publisher:**

```
✓ Published message 1 to ESB.RND.OTEL.QA
  Content: {"id": 1, "timestamp": "2026-01-15T...", "data": "Mock audit data 1"}
```

**Subscriber:**

```
✓ Received message #1:
  Content: {"id": 1, ...}
  Published: 2026-01-15T10:30:00Z
```

## To Connect to REAL IBM MQ

Your current error happens because the IBM MQ Go library needs native IBM MQ Client installed on Windows.

### Steps to Connect to Real MQ at 10.25.135.209:1415

#### Option 1: Install IBM MQ Client (Recommended)

1. **Download IBM MQ Redistributable Client**

   - Visit: https://www.ibm.com/support/pages/downloading-ibm-mq-clients
   - Choose "IBM MQ Client 9.x for Windows"

2. **Install & Configure**

   - Run installer
   - Add `C:\Program Files\IBM\MQ\bin` to your PATH

3. **Update Code**
   - I can provide the real IBM MQ code once client is installed
   - Update go.mod to include: `github.com/ibm-messaging/mq-golang/v5`

#### Option 2: IBM MQ REST API

If your MQ server has REST API enabled, we can use HTTP calls (no client libraries needed).

#### Option 3: STOMP Protocol

If your MQ admin enables STOMP (port 61613), we can use pure Go client.

## Files Structure

- [main.go](main.go) - Entry point
- [config.go](config.go) - MQ configuration
- [publisher.go](publisher.go) - Publishes messages
- [subscriber.go](subscriber.go) - Receives messages
- [queue.go](queue.go) - In-memory queue implementation

## Next Steps

1. ✅ Test this demo version first
2. Choose your preferred approach for real MQ connection
3. Let me know which option works for you!

---

**Need help?** Let me know which approach (Client install / REST / STOMP) you'd like to use!
