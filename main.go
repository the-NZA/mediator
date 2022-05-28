package main

import "fmt"

// Userer defines interface for user.
type Userer interface {
	Name() string
	Send(Userer, string)
	Receive(Userer, string)
}

// Chatter defines interface for chat.
type Chatter interface {
	Register(Userer) error
	Send(from Userer, to Userer, message string)
}

// User defines user structure.
type User struct {
	chat Chatter
	name string
}

// Name returns name of user.
func (u User) Name() string {
	return u.name
}

// Send tries to send message to user with chat interface.
func (u User) Send(to Userer, s string) {
	u.chat.Send(u, to, s)
}

// Receive tries to receive message from user with chat interface.
func (u User) Receive(from Userer, s string) {
	fmt.Printf("%s received from %s: %s\n", u.name, from.Name(), s)
}

// ChatRoom defines chat structure.
type ChatRoom struct {
	users map[string]struct{}
}

// Register tries to register user in current chat room.
func (c ChatRoom) Register(userer Userer) error {
	if _, ok := c.users[userer.Name()]; ok {
		return fmt.Errorf("user %s already registered", userer.Name())
	}

	c.users[userer.Name()] = struct{}{}
	return nil
}

// Send tries to send message to user.
func (c ChatRoom) Send(from, to Userer, message string) {
	to.Receive(from, message)
}

func main() {
	chat := ChatRoom{users: make(map[string]struct{})}

	alice := User{chat: chat, name: "Alice"}
	bob := User{chat: chat, name: "Bob"}

	err := chat.Register(alice)
	if err != nil {
		fmt.Println(err)
	}
	err = chat.Register(bob)
	if err != nil {
		fmt.Println(err)
	}

	alice.Send(bob, "Hi Bob")
	bob.Send(alice, "What's up Alice")
	alice.Send(bob, "Not much, how about you?")
	bob.Send(alice, "Not much either, just chilling")
	alice.Send(bob, "That's fine, see you later")
	alice.Send(bob, "Bye")
	bob.Send(alice, "Yeah, see you soon")
}
