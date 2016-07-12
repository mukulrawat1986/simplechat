package main

import (
	"log"
	"net"
)

// ChatRoom handles the data for a chatroom
type ChatRoom struct {
}

// NewChatRoom will create an empty chatroom type
func NewChatRoom() *ChatRoom {
	return &ChatRoom{}
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
}

func NewChatUser(conn net.Conn) *ChatUser {
	return &ChatUser{}
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
}
