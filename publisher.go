package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ibm-messaging/mq-golang/v5/ibmmq"
)

func runPublisher() {
	// Connect to queue manager
	qMgr, err := connectToQueueManager()
	if err != nil {
		log.Fatalf("Failed to connect to queue manager: %v", err)
	}
	defer qMgr.Disc()

	// Open queue for output
	qObject, err := openQueue(qMgr, ibmmq.MQOO_OUTPUT)
	if err != nil {
		log.Fatalf("Failed to open queue: %v", err)
	}
	defer qObject.Close(0)

	// Send mock messages
	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf(`{"id": %d, "timestamp": "%s", "data": "Mock audit data %d"}`,
			i, time.Now().Format(time.RFC3339), i)

		putmqmd := ibmmq.NewMQMD()
		pmo := ibmmq.NewMQPMO()
		pmo.Options = ibmmq.MQPMO_NO_SYNCPOINT

		err = qObject.Put(putmqmd, pmo, []byte(message))
		if err != nil {
			log.Printf("Failed to put message: %v", err)
		} else {
			fmt.Printf("✓ Published message %d: %s\n", i, message)
		}

		time.Sleep(1 * time.Second)
	}

	fmt.Println("\nPublisher completed successfully!")
}

func connectToQueueManager() (ibmmq.MQQueueManager, error) {
	cno := ibmmq.NewMQCNO()
	cd := ibmmq.NewMQCD()

	cd.ChannelName = Channel
	cd.ConnectionName = fmt.Sprintf("%s(%s)", Host, Port)
	cd.ChannelType = ibmmq.MQCHT_CLNTCONN

	cno.ClientConn = cd
	cno.Options = ibmmq.MQCNO_CLIENT_BINDING

	qMgr, err := ibmmq.Connx(QueueManagerName, cno)
	if err != nil {
		return qMgr, err
	}

	fmt.Println("✓ Connected to queue manager:", QueueManagerName)
	return qMgr, nil
}

func openQueue(qMgr ibmmq.MQQueueManager, openOptions int32) (ibmmq.MQObject, error) {
	mqod := ibmmq.NewMQOD()
	mqod.ObjectType = ibmmq.MQOT_Q
	mqod.ObjectName = QueueName

	qObject, err := qMgr.Open(mqod, openOptions)
	if err != nil {
		return qObject, err
	}

	fmt.Println("✓ Opened queue:", QueueName)
	return qObject, nil
}
