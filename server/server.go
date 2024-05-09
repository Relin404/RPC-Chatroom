package main

import (
	"chatrooms/commons"
	"fmt"
	"log"
	"net"
	rpc "net/rpc"
	"strconv"
)

var registeredPorts []int

type Listener int

func (l *Listener) SendMessage(args *commons.Args, ack *bool) error {
	for _, port := range registeredPorts {
		client, err := rpc.Dial("tcp", "0.0.0.0:"+strconv.Itoa(port))
		if err != nil {
			log.Fatalf("Error in dialing: %v", err)
		}

		var messageReceivedAck bool

		err = client.Call("ClientListener.ReceiveMessage", args, &messageReceivedAck)
		if err != nil {
			log.Fatalf("Error in remote call: %v", err)
		}

	}

	message := "[" + args.Name + "]: " + "\"" + args.Message + "\""
	fmt.Println(message)

	*ack = true

	return nil
}

func (l *Listener) RegisterClient(clientPort int, ack *bool) error {
	registeredPorts = append(registeredPorts, clientPort)

	fmt.Println("A new client has been registered with port: " + strconv.Itoa(clientPort))

	*ack = true

	return nil
}

func main() {
	address, err := net.ResolveTCPAddr("tcp", commons.GetServerAddress())
	if err != nil {
		log.Fatalf("Error in resolving address: %v", err)
	}

	inbound, err := net.ListenTCP("tcp", address)
	if err != nil {
		log.Fatalf("Error in listening: %v", err)
	}

	fmt.Println("Server is listening on port: " + strconv.Itoa(inbound.Addr().(*net.TCPAddr).Port))
	fmt.Println("Abandon hope all, ye who enter here")

	listener := new(Listener)
	rpc.Register(listener)
	rpc.Accept(inbound)
}
