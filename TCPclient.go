package main

import "net"
import "time"

const (
	
)
var (
	servers []string
	alive bool = false
	alive2 bool = false
	alive3 bool = false
	conn2 net.Conn
	conn3 net.Conn
)

func main() {
	servers = append(servers,"152.94.1.141:10349")//machine 1
	servers = append(servers,"152.94.1.142:10321")//machine 2
	servers = append(servers,"152.94.1.143:16065")//machine 3
	conn, _ := net.Dial("tcp", servers[0])
	conn2, _ := net.Dial("tcp", servers[1])
	conn3, _ := net.Dial("tcp", servers[2])

	go sendHandler(conn)
	go sendHandler(conn2)
	go sendHandler(conn3)

	delay := time.NewTimer(3*time.Second)
	<- delay.C

	go replyHandler(conn)
	go replyHandler(conn2)
	go replyHandler(conn3)

	delay2 := time.NewTimer(3*time.Second)
	<- delay2.C

	go reportHandler(conn)
	go reportHandler(conn2)
	go reportHandler(conn3)
	

	for {

	}

}

func sendHandler(conn net.Conn) {
	//SEND HEARTBEAT EVERY 10TH SECOND
	ticker := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})
	for {
		select {
		 case <- ticker.C:
			switch {
			 case conn.RemoteAddr().String() == "152.94.1.141:10349":
				alive = false
			 case conn.RemoteAddr().String() == "152.94.1.142:10321":
				alive2 = false
			 case conn.RemoteAddr().String() == "152.94.1.143:10328":
				alive3 = false
			}

			 b := []byte("Heartbeat")
			 conn.Write(b)
			 println("Sent Heartbeat to server"+conn.RemoteAddr().String()+"\n")
		 case <- quit:
			 ticker.Stop()
			 return
		 }
	 }
}

func replyHandler(conn net.Conn) {
 //REPLY EVERY 10TH SECOND WITH 3 SECONDS DELAY THE FIRST TIME
 	tickerR := time.NewTicker(10 * time.Second)
 	quitR := make(chan struct{})  
				 for {
					select {
					 case <- tickerR.C:

						switch {
						case conn.RemoteAddr().String() == "152.94.1.141:10349":
							alive = true
						case conn.RemoteAddr().String() == "152.94.1.142:10321":
							alive2 = true
						case conn.RemoteAddr().String() == "152.94.1.143:10328":
							alive3 = true
						}

			 	    buf := make([]byte, 1024)		 
	 				n, _ := conn.Read(buf)
					msg := string(buf[:n])
					 println("reply from server "+conn.RemoteAddr().String()+"="+ string(msg)+"\n")
					 
					case <- quitR:
						
						tickerR.Stop()
					 }
					 }
				 
	}

	func reportHandler(conn net.Conn) {
		//REPORT TO SERVERS EVERY 10TH SECOND WITH 6 SECONDS DELAY THE FIRST TIME
	tickerRep := time.NewTicker(10 * time.Second)
	quitRep := make(chan struct{})
    		for {
       			select {
					case <- tickerRep.C:
					switch {
					case alive == false:

						conn.Write([]byte("Server: 1 disconnected!\n"))
						println("Server 1 disconnected!\n")
					case alive2 == false:
						conn.Write([]byte("Server: 2 disconnected!\n"))
						println("Server 2 disconnected!")
					case alive3 == false:
						conn.Write([]byte("Server: 3 disconnected!\n"))
						println("Server 3 disconnected!")
					}
        			case <- quitRep:
            			tickerRep.Stop()
            			return
					}
					
    			}
	}
