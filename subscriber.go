package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ibm-messaging/mq-golang/v5/ibmmq"
)

func runSubscriber() {
	// Connect to queue manager
	qMgr, err := connectToQueueManager()
	if err != nil {
		log.Fatalf("Failed to connect to queue manager: %v", err)
	}
	defer qMgr.Disc()

	// Open queue for input
	qObject, err := openQueue(qMgr, ibmmq.MQOO_INPUT_AS_Q_DEF)
	if err != nil {
		log.Fatalf("Failed to open queue: %v", err)
	}
	defer qObject.Close(0)

	fmt.Println("✓ Queue opened for receiving")
	fmt.Println("\n📬 Waiting for messages (press Ctrl+C to stop)...\n")

	// Receive messages
	msgCount := 0
	for {
		getmqmd := ibmmq.NewMQMD()
		gmo := ibmmq.NewMQGMO()
		gmo.Options = ibmmq.MQGMO_NO_SYNCPOINT | ibmmq.MQGMO_WAIT
		gmo.WaitInterval = 3000 // Wait 3 seconds for a message

		buffer := make([]byte, 10240)
		datalen, err := qObject.Get(getmqmd, gmo, buffer)

		if err != nil {
			mqret := err.(*ibmmq.MQReturn)
			if mqret.MQCC == ibmmq.MQCC_FAILED && mqret.MQRC == ibmmq.MQRC_NO_MSG_AVAILABLE {
				fmt.Printf("⏱️  Still listening... (waiting for messages)\n")
				continue
			}
			log.Printf("Error getting message: %v", err)
			continue
		}

		if datalen > 0 {
			msgCount++
			fmt.Printf("✓ Received message #%d:\n", msgCount)
			fmt.Printf("  Content: %s\n", string(buffer[:datalen]))
			fmt.Printf("  Message ID: %x\n", getmqmd.MsgId)
			fmt.Printf("  Timestamp: %s\n\n", time.Now().Format(time.RFC3339))
		}
	}
}
