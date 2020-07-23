/*
Implementing using a doubly linked list for fast deletion and insertion  and a hash map for fast access.
1. newnode
2.add
3.get
4.show
5.len
6.move to front
7. remove
*/

package main

import (
	"errors"
	"fmt"
	"os"
)

type node struct {
	Key, Value int
	prev, next *node
}

type queue struct {
	head       *node
	tail       *node
	dictionary map[int]*node
	capacity   int
}

func newNode(key, value int) *node {
	temp := &node{
		Key:   key,
		Value: value,
		prev:  nil,
		next:  nil,
	}
	return temp
}

func initqueue(cap int) queue {
	newqueue := queue{
		capacity:   cap,
		dictionary: make(map[int]*node),
	}
	return newqueue
}

func (LRU *queue) get(key int) (int, error) {
	if val, ok := LRU.dictionary[key]; ok {
		LRU.removeNode(val)
		LRU.addToFront(val)
		return val.Value, nil
	}
	return -1, errors.New("key not found")
}

func (LRU *queue) add(key, value int) {
	//Check if exits , if yes then overwrite and push to front
	if item, ok := LRU.dictionary[key]; ok {
		item.Value = value

		LRU.removeNode(item)
		LRU.addToFront(item)
	} else {
		//check for capacity
		if len(LRU.dictionary) >= LRU.capacity {
			delete(LRU.dictionary, LRU.head.Key)
			LRU.removeNode(LRU.head)
		}

		new := newNode(key, value)
		LRU.addToFront(new)
		LRU.dictionary[key] = new
	}
}

func (LRU *queue) removeNode(n *node) {
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		LRU.head = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	} else {
		LRU.tail = n.prev
	}
}

func (LRU *queue) addToFront(n *node) {

	if LRU.tail != nil {
		LRU.tail.next = n
	}
	n.prev = LRU.tail
	n.next = nil
	LRU.tail = n

	// one item
	if LRU.head == nil {
		LRU.head = LRU.tail
	}
}

func (LRU *queue) length() int {
	temp := LRU.head
	count := 0
	for temp != nil {
		count++
		temp = temp.next
	}
	return count
}

func (LRU *queue) show() {
	temp := LRU.head
	for temp != nil {
		fmt.Printf("(%d : %d) \n", temp.Key, temp.Value)

		temp = temp.next
	}
}

func main() {
	var lru queue
	var n, key, val int
	fmt.Println("Enter the capacity: ")
	fmt.Scan(&n)
	lru = initqueue(n)
	for {
		var choice int
		fmt.Println("\nChoose your command.\n 1.Add\n2.Get\n3.Show\n4.Length\n5.Exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Print("enter key: ")
			fmt.Scan(&key)
			fmt.Print("enter Value: ")
			fmt.Scan(&val)
			lru.add(key, val)
			fmt.Println("\nSuccessfully entered\n")
		case 2:
			fmt.Print("Enter the key :")
			fmt.Scan(&key)
			v, err := lru.get(key)
			if err != nil {
				fmt.Println("\n", err)
			} else {
				fmt.Printf("\nkey:%d : Value: %d\n", key, v)
			}
		case 3:
			lru.show()
		case 4:
			fmt.Println("lenght of LRU: ", lru.length())
		case 5:
			os.Exit(1)
		}
	}
}
