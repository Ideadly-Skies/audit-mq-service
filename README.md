# Audit MQ Service

Simple Go publisher-subscriber for IBM MQ.

## Configuration

- **Queue Manager**: SIDPOP01
- **Host**: 10.25.135.209
- **Port**: 1415
- **Queue**: ESB.RND.OTEL.QA

## Setup

1. Install dependencies:

```bash
go mod download
```

## Usage

### Run Publisher

Sends 5 mock audit messages to the queue:

```bash
go run . publisher
```

### Run Subscriber

Listens for messages from the queue:

```bash
go run . subscriber
```

## Testing

1. Open two terminal windows
2. In terminal 1: `go run . subscriber`
3. In terminal 2: `go run . publisher`
4. Watch messages flow from publisher to subscriber

## Notes

- The subscriber waits indefinitely for messages (Ctrl+C to stop)
- Messages are JSON formatted with id, timestamp, and mock data
- No Docker required - runs directly against the MQ server
