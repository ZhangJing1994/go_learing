package main

import (
	"lession2/week9/goim-simulate/pkg"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	index := 0
	version := 1
	code := 1
	timeout := time.After(10 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8082")
	if err != nil {
		log.Println(err.Error())
	}
	defer conn.Close()
	for {
		select {
		case <-timeout:
			log.Println("client connect is timeout")
			return
		default:
			for i := 0; i < 5; i++ {
				data := pkg.Encoder(pkg.NewPack(version, code, index, []byte("test"+strconv.Itoa(index))))
				index++
				conn.Write(data)
			}
			time.Sleep(time.Second)
		}
	}
}
