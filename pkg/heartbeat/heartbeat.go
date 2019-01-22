package heartbeat

import (
	"objectStorage/pkg/heartbeat/mq"
	"os"
	"time"
)

func StartHeartbeat() {
	q := mq.Nwe(os.Getenv("MQ_SERVER"))
	defer q.Close()
	for {
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5 * time.Second)
	}
}
