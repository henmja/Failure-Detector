package main

import (
    "fmt"
    "net"
    "os"
)

const (
    CONN_HOST = "152.94.1.143" //Machine 3
    CONN_PORT = "16065"
    CONN_TYPE = "tcp"
)

func main() {
    // Listen for incoming connections.
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	fmt.Println("\nNodeID=0")
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	// Read the incoming connection into the buffer.
	for {
		buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
	  fmt.Println("Error reading:", err.Error())
	}
	decMsg := string(buf[:n])

	fmt.Print("Received " + decMsg + "From client" + conn.RemoteAddr().String()+"\n") //string(msg)
	if decMsg=="Heartbeat" {
	conn.Write([]byte("Heartbeat"))

	fmt.Println("Sent heartbeat to client" + conn.RemoteAddr().String()+"\n")
		}
    }
  }
