package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  go run . publisher   - Run publisher to send messages")
		fmt.Println("  go run . subscriber  - Run subscriber to receive messages")
		os.Exit(1)
	}

	mode := os.Args[1]

	switch mode {
	case "publisher", "pub":
		fmt.Println("🚀 Starting Publisher...")
		fmt.Printf("Target: %s:%s (QM: %s, Queue: %s)\n\n", Host, Port, QueueManagerName, QueueName)
		runPublisher()
	case "subscriber", "sub":
		fmt.Println("📡 Starting Subscriber...")
		fmt.Printf("Target: %s:%s (QM: %s, Queue: %s)\n", Host, Port, QueueManagerName, QueueName)
		runSubscriber()
	default:
		fmt.Printf("Unknown mode: %s\n", mode)
		fmt.Println("Use 'publisher' or 'subscriber'")
		os.Exit(1)
	}
}
