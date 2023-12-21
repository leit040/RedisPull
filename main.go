package RedisPull

import (
	"context"
	"fmt"
	"log"
)

func Init() {
	conn, e := NewConnections()
	if e != nil {
		log.Fatalf("Error initializing connections: %s", e)
	}
	conn1, err := conn.getConnection("domain.com")
	if err {
		log.Fatal(err)
	}
	_, err2 := conn1.GoRedis.Ping(context.Background()).Result()
	if err2 != nil {
		fmt.Println("ERROR")
	}
}
