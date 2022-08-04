package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/exp/slices"
)

// vars
var helpvar int

func main() {

	activeThreads := 0
	channel := make(chan bool)
	ip_address := "192.168.0.104"
	end_port := checkFlags()
	for port := 0; port <= end_port; port++ {
		go tcp_scan(ip_address, port, channel)
		activeThreads++

	}
	for activeThreads > 0 {
		<-channel
		activeThreads--

	}
	fmt.Println("All done VÄlkommen åter")

}

func checkFlags() int {
	if slices.Contains(os.Args, "-s") {
		return 1024
	} else if slices.Contains(os.Args, "-f") {
		return 65535
	}
	return 0
}

func tcp_scan(ip_address string, port int, channel chan bool) {

	address := fmt.Sprintf("%s:%d", ip_address, port)
	conn, err := net.Dial("tcp", address)

	if err != nil {
		//fmt.Printf("Port %d is closed\n", arr[i])
		channel <- true
		return
	}
	buffer := make([]byte, 4096)
	conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	numbytesread, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		channel <- true
		return
	}
	conn.Close()
	//log.Printf("Banner from port %d\n%s\n", successful_ports[i], buffer[0:numbytesread])
	fmt.Printf("Port %d open\n", port)
	fmt.Println(string(buffer[0:numbytesread]))
	fmt.Println("------------------")

	//what_service := netdb.ServiceByPort(i, "tcp")
	//fmt.Println(fmt.Sprint(what_service.Port) + " " + what_service.Name + " OPEN")

}
