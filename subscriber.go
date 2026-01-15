package main

import (
	"fmt"
	"time"
)

func runSubscriber() {
	fmt.Println("✓ Connected to Queue Manager:", QueueManagerName)
	fmt.Printf("✓ Host: %s:%s\n", Host, Port)
	fmt.Println("✓ Subscribed to Queue:", QueueName)
	fmt.Println("\n📬 Waiting for messages (press Ctrl+C to stop)...\n")

	// Subscribe to messages
	msgChan := messageQueue.Subscribe()
	msgCount := 0

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case msg := <-msgChan:
			msgCount++
			fmt.Printf("✓ Received message #%d:\n", msgCount)
			fmt.Printf("  ID: %d\n", msg.ID)
			fmt.Printf("  Content: %s\n", msg.Content)
			fmt.Printf("  Published: %s\n", msg.Timestamp.Format(time.RFC3339))
			fmt.Printf("  Received: %s\n\n", time.Now().Format(time.RFC3339))
		case <-ticker.C:
			fmt.Println("⏱️  Still listening... (waiting for messages)")
		}
	}
}
