//// 2023/10/23 // 16:20 //

//// Introduction to Pointers

// // As we have learned, a variable is a named location in memory that stores a
// // value. We can manipulate the value of a variable by assigning a new value to it
// // or by performing operations on it. When we assign a value to a variable, we
// // are storing that value in a specific location in memory.
// x := 42
// // "x" is the name of a location in memory. That location is storing the integer value of 42

// // A pointer is a variable

// // A pointer is a variable that stores the memory address of another variable. This
// // means that a pointer "points to" the locaiton of where the dat is stored NOT
// // the actual data itself.

// // The * syntax defines a pointer:
// var p *int

// // The & operator generates a pointer to its operand.
// myString := "hello"
// myStringPrt = &myString

// Pointers allow us to manipulate data in memory diretly, without making
// copies or duplicating data. This can make programs more efficient and allow
// us to do things that would be difficult or impossible without them.

//// Assignment
// // Fix the bug in the sendMessage function. It's supposed to print a nicely
// // formatted message to the console containing an SMS's recipient and message
// // body. However, it's not working as expected. Run the code and see what
// // happend, then fix the bug.

// package main

// import "fmt"

// type Message struct {
// 	Recipient string
// 	Text      string
// }

// func sendMessage(m Message) {
// 	fmt.Printf("To: %v\n", m.Recipient)
// 	fmt.Printf("Message: %v\n", m.Text)
// }

// func test(recipient string, text string) {
// 	m := Message{Recipient: recipient, Text: text}
// 	sendMessage(m)
// 	fmt.Println("=======================================")
// }

// func main() {
// 	test("Lane", "Textio is getting better everyday!")
// 	test("Allan", "This pointer stuff is weird...")
// 	test("Tiffany", "What time will you be home for dinner?")
// }

// //// 16:30

// //// Pointers Practice

// // Pointers hold the memory address of a value.

// // The * syntax defines a pointer:
// var p *int

// // A pointer's zero value is nil

// // The & operator generates a pointer to its operand.
// myString := "hello"
// myStringPtr = &myString

// // The * dereferences a pointer to gain access to the value
// fmt.Println(*myStringPtr) // read myString through the pointer
// *myStringPtr = "world" // set myString through the pointer

// // Unlike C, Go has no pointer arothmetic

// // Just because you can doesn't mean you should

// We're doing this exercise to understand that pointers can be used in this way.
// That said, pointers can be very dangerous. It's generally a better idea to have
// your functions accept non-pointers and return new values rather than
// mutating pointer inputs.

//// Assignment

// Complete the removeProfanity function.

// It should use the strings.ReplaceAll function to replace all instances of the
// following words in the input message with asterisks.
// * "dang" -> "****"
// * "shoort" -> "*****"
// * "heck" -> "****"

// // It should mutate the value in the pointer and return nothing.

// // Do not alter the function signature.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func removeProfanity(message *string) {
// 	messageVal := *message
// 	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
// 	messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
// 	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
// 	*message = messageVal
// }

// func test(messages []string) {
// 	for _, message := range messages {
// 		removeProfanity(&message)
// 		fmt.Println(message)
// 	}
// }

// func main() {
// 	messages1 := []string{
// 		"well shoot, this is awful",
// 		"dang robots",
// 		"dang them to heck",
// 	}

// 	messages2 := []string{
// 		"well shoot",
// 		"Allan is going straight to heck",
// 		"dang... that's a tough break",
// 	}

// 	test(messages1)
// 	test(messages2)
// }

//// 16:42

//// Nil Dereference

// Pointers can be very dangerous.

// If a pointer points to nothing (the zero value of the pointer type) then
// dereferencing it will cause a runtime error (a panic) that crashes the program.
// Generally speaking, whenever you're dealing with pointers you should check if
// it's nil before trying to dereference it.

//// Assignment

// // Let's make our profanity checker safe. Update the removreProfanity function.
// // If message is nil, return early to avoid a panic. After all, there are no bad
// // words to remove.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func removeProfanity(message *string) {
// 	if message == nil {
// 		return
// 	}
// 	messageVal := *message
// 	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
// 	messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
// 	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
// 	*message = messageVal
// }

// func test(messages []string) {
// 	for _, message := range messages {
// 		if message == "" {
// 			removeProfanity(nil)
// 			fmt.Println("nil message detected")
// 		} else {
// 			removeProfanity(&message)
// 			fmt.Println(message)
// 		}
// 	}
// }

// func main() {
// 	messages := []string{
// 		"well shoot, this is awful",
// 		"",
// 		"dang robots",
// 		"dang them to heck",
// 		"",
// 	}

// 	messages2 := []string{
// 		"well shoot",
// 		"",
// 		"Allan is going straight to heck",
// 		"dang... that's a tough break",
// 		"",
// 	}

// 	test(messages)
// 	test(messages2)
// }

//// 16:53
//// Pause
//// 17:25

//// Pointer Receivers

// A receiver type on a method can be a pointer.

// Methods with pointer receivers can modify the value to which the receiver
// points. Since methods often need to modify their receiver, pointer receivers
// are more common than value receivers.

// Pointer receivers and non-pointer receivers

// package main

// Pointer receiver

// type car struct {
// 	color string
// }

// func (c *car) setColor(color string) {
// 	c.color = color
// }

// func main() {
// 	c := car{
// 		color: "white",
// 	}
// 	c.setColor("blue")
// 	fmt.Println(c.color)
// 	// prints "blue"
// }

// // Non-pointer receiver
// type car struct {
// 	color string
// }

// func (c car) setColor(color string) {
// 	c.color = color
// }

// func main() {
// 	c := car{
// 		color: "white",
// 	}
// 	c.setColor("blue")
// 	fmt.Println(c.color)
// 	// prints "white"
// }

//// 17:30

//// Pointer Receiver Code

// // Methods with pointer receivers don't require that a pointer is used to call the
// // method. The pointer will automatically be derived from the value.

// package main

// import "fmt"

// type circle struct {
// 	x      int
// 	y      int
// 	radius int
// }

// func (c *circle) grow() {
// 	c.radius *= 2
// }

// func main() {
// 	c := circle{
// 		x:      1,
// 		y:      2,
// 		radius: 4,
// 	}

// 	// notice c is not a pointer in the calling function
// 	// but the method still gains access to a pointer to c
// 	c.grow()
// 	fmt.Println(c.radius)
// 	/// prints 8
// }

//// Assignment

// Fix the bug in the code so that setMessage sets the message field of the
// given email structure, and the new value persists outside the scope of the
// setMessage method.

package main

import (
	"fmt"
)

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func test(e *email, newMessage string) {
	fmt.Println("-- before -- ")
	e.print()
	fmt.Println("-- end before --")
	e.setMessage("this is my second draft")
	fmt.Println("-- after --")
	e.print()
	fmt.Println("-- end after --")
	fmt.Println("=========================")
}

func (e email) print() {
	fmt.Println("message:", e.message)
	fmt.Println("fromAddress:", e.fromAddress)
	fmt.Println("toAddress:", e.toAddress)
}

func main() {
	test(&email{
		message:     "this is my first draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my second draft")

	test(&email{
		message:     "this is my third draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my fourth draft")
}
