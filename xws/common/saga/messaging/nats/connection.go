package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func getConnection(host, port, user, password string) (*nats.Conn, error) {
	url := fmt.Sprintf("nats://%s:%s@%s:%s", user, password, host, port)
	connection, err := nats.Connect(url)
	//connection, err := nats.Connect(fmt.Sprintf("nats://localhost:%s", port))
	if err != nil {
		return nil, err
	}
	return connection, nil
}
