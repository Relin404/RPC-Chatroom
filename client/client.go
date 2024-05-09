package main

import (
	"bufio"
	"chatrooms/commons"
	"fmt"
	"log"
	"net"
	rpc "net/rpc"
	"os"
	"strconv"
)

var clientPort int

func GetAvailablePort() (int, error) {
	newAddress, err := net.ResolveTCPAddr("tcp", "0.0.0.0:0")

	if err != nil {
		return 0, err
	}

	inbound, err := net.ListenTCP("tcp", newAddress)
	if err != nil {
		return 0, err
	}

	defer inbound.Close()

	return inbound.Addr().(*net.TCPAddr).Port, nil

}

func main() {
	client, err := rpc.Dial("tcp", commons.GetServerAddress())

	if err != nil {
		log.Fatal(err)
	}

	in := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your nickname to join the chat:")

	nicknameBuffer, _, err := in.ReadLine()
	var nickname = string(nicknameBuffer)

	clientPort, err = GetAvailablePort()

	go clientListen()

	var newClientAck bool

	err = client.Call("Listener.RegisterClient", clientPort, &newClientAck)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Your message: ")

	for {
		msg, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		args := &commons.Args{Message: string(msg), Name: nickname}
		var messageSentAck bool
		err = client.Call("Listener.SendMessage", args, &messageSentAck)
		if err != nil {
			log.Fatal(err)
		}
	}

}

type ClientListener int

func (l *ClientListener) ReceiveMessage(args *commons.Args, ack *bool) error {
	message := "[" + args.Name + "]: " + "\"" + args.Message + "\""
	fmt.Println(message)
	fmt.Println("Enter your message:")

	*ack = true

	return nil

}

func clientListen() {
	addrResolved, err := net.ResolveTCPAddr("tcp", "0.0.0.0:"+strconv.Itoa(clientPort))
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addrResolved)

	if err != nil {
		log.Fatal(err)
	}

	clientListener := new(ClientListener)
	rpc.Register(clientListener)
	rpc.Accept(inbound)
}
