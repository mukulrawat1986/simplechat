package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

// ChatRoom handles the data for a chatroom
type ChatRoom struct {
	users       map[string]*ChatUser
	incoming    chan string
	joins       chan *ChatUser
	disconnects chan string
}

// NewChatRoom will create a new chatroom type
func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		users:       make(map[string]*ChatUser),
		incoming:    make(chan string),
		joins:       make(chan *ChatUser),
		disconnects: make(chan string),
	}
}

func (cr *ChatRoom) ListenForMessages() {
}

func (cr *ChatRoom) Logout(username string) {
}

func (cr *ChatRoom) Join(conn net.Conn) {
}

func (cr *ChatRoom) BroadCast(msg string) {
}

type ChatUser struct {
	conn       net.Conn
	disconnect bool
	username   string
	outgoing   chan string
	reader     *bufio.Reader
	writer     *bufio.Writer
}

func NewChatUser(conn net.Conn) *ChatUser {
	return &ChatUser{
		conn:       conn,
		disconnect: false,
		outgoing:   make(chan string),
		reader:     bufio.NewReader(conn),
		writer:     bufio.NewWriter(conn),
	}
}

func (cu *ChatUser) ReadIncomingMessages(chatroom *ChatRoom) {
}

func (cu *ChatUser) WriteOutgoingMessages(chatroom *ChatRoom) {
}

func (cu *ChatUser) Login(chatroom *ChatRoom) error {
	return nil
}

func (cu *ChatUser) ReadLine() (string, error) {
	return "", nil
}

func (cu *ChatUser) WriteString(msg string) error {
	return nil
}

func (cu *ChatUser) Send(msg string) {
}

func (cu *ChatUser) Close() {
}

// main will create a socket, bind to port 6677,
// and loop while waiting for connections.
//
// When it receives a connection it will pass it to
// `chatroom.Join()`.

func main() {
	log.Println("Chat server starting!")

	// create a TCP listener on port 6677
	listener, err := net.Listen("tcp", ":6677")

	if err != nil {
		log.Fatalf("Error when listening on port 6677 : %s", err)
		os.Exit(1)
	}

	// create a new instance of the chatroom using NewChatRoom()
	chatroom := NewChatRoom()

	// start listening for messages
	chatroom.ListenForMessages()

	// Loop and listen for accepted connections on port 6677
	for {
		// wait for the next call and return a connection
		conn, err := listener.Accept()

		if err != nil {
			log.Fatalf("Error when accepting connection: %s", err)
			os.Exit(1)
		}

		log.Printf("Remote address is: %s\n", conn.RemoteAddr())
	}
}
