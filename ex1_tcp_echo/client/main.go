package main

import (
	"bufio"   // buffers for reading
	"context" // pkg that allows us to set timeout
	"fmt"     // output into terminal
	"log"     // logging pkg
	"net"     // main pkg for networking
	"time"    // pkg for time constant
)

func main() {
	// create a dialer
	var d net.Dialer

	// our msg
	message := "Hello there!\n"
	// server port number
	const port = 8081

	// message[:len(message) - 1] - removes '\n' for logging
	fmt.Printf("Sending message: %s; to port: %d\n", message[:len(message)-1], port)

	// create call context that should close when timeout reached
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	// call cancel function to context when we end with our tasks
	defer cancel()

	// connect to server with context
	conn, err := d.DialContext(ctx, "tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	// call close to connection when we end with our tasks
	defer conn.Close()

	// send some data to server
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}

	// create buffer and read message from server
	getMessage, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Message recieved: %s\n", getMessage[:len(getMessage)-1])
}
