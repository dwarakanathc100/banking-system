package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

type Client struct {
	conn net.Conn
	name string
}

var (
	clients    = make(map[net.Conn]string)
	mutex      sync.Mutex
	messagesCh = make(chan string) // channel to broadcast messages
)

func StartServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started on port 8080")

	go broadcaster() // Start broadcaster goroutine

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go handleClient(conn)
	}
}

func broadcaster() {
	for msg := range messagesCh {
		mutex.Lock()
		for conn := range clients {
			fmt.Fprintln(conn, msg)
		}
		mutex.Unlock()
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	fmt.Fprint(conn, "Enter your name: ")
	nameInput := bufio.NewReader(conn)
	name, _ := nameInput.ReadString('\n')
	name = strings.TrimSpace(name)

	mutex.Lock()
	clients[conn] = name
	mutex.Unlock()

	fmt.Println(name, "joined the chat.")
	messagesCh <- fmt.Sprintf("%s joined the chat!", name)

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(name, "disconnected.")
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			messagesCh <- fmt.Sprintf("%s left the chat.", name)
			return
		}

		message = strings.TrimSpace(message)
		messagesCh <- fmt.Sprintf("[%s]: %s", name, message)
	}
}

// Reuse same binary for client mode
func StartClient() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer conn.Close()

	go readMessages(conn)

	fmt.Println("Connected to chat. Type your name and start chatting:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Fprintln(conn, msg)
	}
}

func readMessages(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Disconnected from server.")
			os.Exit(0)
		}
		fmt.Print(msg)
	}
}

func main() {
	fmt.Println("=== Go Channel Chat App ===")
	fmt.Println("1. Start Chat Server")
	fmt.Println("2. Start Chat Client")
	fmt.Print("Enter choice: ")

	reader := bufio.NewReader(os.Stdin)
	choice, _, _ := reader.ReadRune()

	switch choice {
	case '1':
		StartServer() // Starts the TCP chat server
	case '2':
		StartClient() // Starts a client that connects to it
	default:
		fmt.Println("Invalid option")
	}
}
