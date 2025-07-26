package main

import (
	"fmt"
	"iter"
	"time"
)

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func dump() {
	fmt.Println("Hello, World!!")
	//var i int
	for i := range 3 {
		fmt.Println(i)
	}
	type Technologies string
	const (
		Python Technologies = "python"
		Node   Technologies = "nodejs"
		Golang Technologies = "golang"
	)

	PrintPackageManager := func(tech Technologies) {
		switch tech {
		case Python:
			fmt.Println("PyPi")
		case Node:
			fmt.Println("NPM")
		case Golang:
			fmt.Println("Github")
		default:
			fmt.Println("Unknown")
		}
	}
	//fmt.Println(Python)
	PrintPackageManager(Node)
	b := [...]int{1, 4: 40, 3, 4}
	fmt.Println(b)

	var threeD [2][3]int
	for i := range 2 {
		for j := range 3 {
			threeD[i][j] = i + j
		}
	}
	fmt.Println(threeD)
	slc := make([]string, 3)
	slc[0] = "a"
	slc[1] = "b"
	slc[2] = "c"
	c := make([]string, len(slc))
	slc = append(slc, "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z")
	copy(c, slc[10:])
	fmt.Println(c)
	fmt.Println(len(slc), cap(slc))
	next_value := counter()
	for range 20 {
		fmt.Println(next_value())
	}
}

func SliceIndex[S ~[]E, E comparable](s S, v E) int {
	for i, vs := range s {
		if v == vs {
			return i
		}
	}
	return -1

}

type List[T any] struct {
	head, tail *element[T]
}
type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				break
			}
		}
	}
}
func implementChannels() {
	newChannel := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		newChannel <- "Hello World"
	}()
	go func() {
		value := <-newChannel
		fmt.Println("From inside goroutine" + value)
	}()
	value := <-newChannel
	time.Sleep(2 * time.Second)
	fmt.Println("This comes after the channel")
	fmt.Print(value)

}
func channelDirections() {
	// Use anonymous functions instead of named functions inside a function
	pingVal := func(c chan<- string, msg string) {
		c <- msg
	}
	pongVal := func(ping <-chan string) {
		msg := <-ping
		fmt.Println(msg)
	}
	ping := make(chan string, 1)
	pingVal(ping, "passed message")
	pongVal(ping)
}

func timeouts() {
	channel := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello World"
	}()
	select {
	case msg := <-channel:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}
}
func main() {
	//dump()
	// implementChannels()
	// channelDirections()
	timeouts()
}
