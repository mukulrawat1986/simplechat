package main

import (
	"bufio"
	"fmt"
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

	// create a new ChatUser object
	chatuser := NewChatUser(conn)

	// Login the user, by calling the ChatUser.Login method on this
	// object
	err := chatuser.Login(cr)
	if err != nil {
		log.Fatalf("Error when logging in the user: %s", err)
		return
	}

	// notifies of a new user by putting the newly created ChatUser
	// object on the joins channel
	cr.joins <- chatuser
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

	// write the banner message
	cu.WriteString("Welcome to Mukul's chat server!\n")

	// Print to the socket
	cu.WriteString("Please enter your username: ")

	// read the username from the socket
	str, err := cu.ReadLine()
	if err != nil {
		log.Fatalf("Error while reading from socket: %s", err)
		return err
	}

	// store the read username
	cu.username = str

	// log the connected user
	log.Println("User logged in: ", cu.username)

	// write back to the socket
	cu.WriteString(fmt.Sprintf("Welcome, %s\n", cu.username))

	return nil
}

func (cu *ChatUser) ReadLine() (string, error) {
	line, _, err := cu.reader.ReadLine()
	if err != nil {
		log.Fatalf("Error while reading a line: %s", err)
		return "", err
	}
	return string(line), nil
}

func (cu *ChatUser) WriteString(msg string) error {
	_, err := cu.writer.WriteString(msg)
	if err != nil {
		log.Fatalf("Error while writing the string to buffer: %s", err)
		return err
	}
	err = cu.writer.Flush()
	if err != nil {
		log.Fatalf("Error while writing to connection %s", err)
		return err
	}
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

		go chatroom.Join(conn)
	}
}
