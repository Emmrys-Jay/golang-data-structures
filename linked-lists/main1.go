package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	next *node
	data int
}

type linkedlist struct {
	head   *node
	length int
}

// Create an instance of a node from the passed in data
func (l *linkedlist) prepend(value int) {
	newNode := node{data: value}
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.length++
	} else {
		l.head = &newNode
		l.length++
	}
	return
}

func (l *linkedlist) printList() {
	if l.head == nil {
		return
	}
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}

func (l *linkedlist) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

// The head element is assigned index 0
func (l *linkedlist) insertIntoList(value, index int) {
	newNode := &node{data: value}

	if l.length == 0 {
		l.head = newNode
		l.length++
		return
	}

	if l.length < index {
		return
	}

	if index == 0 {
		newNode.next = l.head
		l.head = newNode
		l.length++
		return
	}

	previousFromIndexNode := l.head
	for i := 0; i < index-1; i++ {
		previousFromIndexNode = previousFromIndexNode.next
	}
	newNode.next = previousFromIndexNode.next
	previousFromIndexNode.next = newNode
	return
}

func main() {
	myList := linkedlist{}
	//myList.prepend(21)
	//myList.prepend(78)
	//myList.prepend(64)
	//myList.prepend(57)
	//myList.printList()
	//fmt.Println()
	//myList.deleteWithValue(21)
	//myList.printList()
	//myList.insertIntoList(55, 1)
	//myList.printList()

	fmt.Print("Input your command: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		line := strings.Split(text, " ")
		var lineInt []int
		for _, val := range line[1:] {
			x, _ := strconv.Atoi(val)
			lineInt = append(lineInt, x)
		}
		if line[0] == "insert" {
			myList.insertIntoList(lineInt[1], lineInt[0])
		} else if (line[0]) == "k" {
			break
		}
		myList.printList()
	}
	myList.printList()
}
