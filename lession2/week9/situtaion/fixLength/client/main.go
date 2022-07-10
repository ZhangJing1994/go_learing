package main

import (
	config "lession2/week9/situtaion/fixLength"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	after := time.After(5 * time.Second)
	iter := 0
	conn, err := net.Dial("tcp", "localhost:8082")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	// sending message
	for {
		select {
		case <-after:
			log.Println("client connect is time out")
			return
		default:
			for i := 0; i < 10; i++ {
				content := "test[" + strconv.Itoa(iter) + "]"
				_, err = conn.Write(patch(content))
				if err != nil {
					log.Fatal(err.Error())
				}
				iter++
			}
			time.Sleep(1 * time.Second)
		}
	}
}

// patching empty byte into origin message
func patch(message string) []byte {
	res := make([]byte, config.BufferSize)
	copy(res, []byte(message))
	return res
}
