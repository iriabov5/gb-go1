package main

import "fmt"

type node struct {
	element string
	next    *node
	prev    *node
}

type linkedList struct {
	size  int
	first *node
	last  *node
}

func initList() *linkedList {
	return &linkedList{}
}

func (l *linkedList) AddFirst(name string) {
	first := l.first
	newNode := &node{
		element: name,
		next:    first,
	}
	l.first = newNode
	if first == nil {
		l.last = newNode
	} else {
		first.prev = newNode
	}
	l.size++
	return
}

func (l *linkedList) AddLast(name string) {
	last := l.last
	newNode := &node{
		prev:    last,
		element: name,
	}

	l.last = newNode

	if last == nil {
		l.first = newNode
	} else {
		last.next = newNode
	}

	l.size++
	return
}

func (l *linkedList) PollFirst() (string, error) {
	if l.first == nil {
		return "", fmt.Errorf("list is empty")
	}
	first := l.first.element
	l.first = l.first.next
	l.first.prev = nil
	l.size--
	return first, nil
}

func (l *linkedList) PollLast() (string, error) {
	if l.first == nil {
		return "", fmt.Errorf("list is empty")
	}
	last := l.last.element
	l.last = l.last.prev
	l.last.next = nil
	l.size--
	return last, nil
}

func (l *linkedList) PeekFirst() (string, error) {
	if l.first == nil {
		return "", fmt.Errorf("list is empty")
	}
	return l.first.element, nil
}

func (l *linkedList) PeekLast() (string, error) {
	if l.first == nil {
		return "", fmt.Errorf("list is empty")
	}
	return l.first.element, nil
}

func (l *linkedList) Size() int {
	return l.size
}

func main() {
	list := initList()
	fmt.Printf("Add First: A\n")
	list.AddFirst("A")
	fmt.Printf("Add First: B\n")
	list.AddFirst("B")
	fmt.Printf("Add Last: C\n")
	list.AddLast("C")
	list.AddLast("D")

	fmt.Printf("Size: %d\n", list.Size())

	s, err := list.PollFirst()
	if err != nil {
		fmt.Printf("Poll First Error: %s\n", err.Error())
	} else {
		fmt.Println("Poll First: ", s)
	}

	s, err = list.PollLast()
	if err != nil {
		fmt.Printf("Poll Last Error: %s\n", err.Error())
	} else {
		fmt.Println("Poll Last: ", s)
	}

	fmt.Printf("Add First: E\n")
	list.AddFirst("E")

	s, err = list.PeekFirst()
	if err != nil {
		fmt.Printf("Peek First Error: %s\n", err.Error())
	} else {
		fmt.Println("Peek First: ", s)
	}

	fmt.Printf("Size: %d\n", list.Size())
}
