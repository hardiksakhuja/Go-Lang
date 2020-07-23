/*
Implemented Custom Data structure orderedMap which has a map and a slice
the slice is used to store the order of the inputs.
set , delete and show methods are implemented
*/
package main

import (
	"fmt"
	"os"
)

type orderedMap struct {
	hashmap  map[string]string
	keyArray []string
}

func main() {
	orderedMap := &orderedMap{
		hashmap:  make(map[string]string),
		keyArray: make([]string, 0),
	}
	for {
		var choice int
		var key, value string
		fmt.Println("Choose your Option\n1.Add to map\n2.Delete from Map\n3.Show Map\n4.Exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Print("Enter key(string): ")
			fmt.Scan(&key)
			fmt.Print("Enter value(string): ")
			fmt.Scan(&value)
			orderedMap.set(key, value)
		case 2:
			fmt.Print("Enter key(string): ")
			fmt.Scan(&key)
			orderedMap.delete(key)
		case 3:
			orderedMap.show()
		case 4:
			os.Exit(1)
		}
	}
}

func (mymap *orderedMap) set(k string, v string) {
	//check to see if already present
	if _, ok := mymap.hashmap[k]; ok {
		fmt.Println("Key Present. Updating Value")
		mymap.hashmap[k] = v
	} else {
		mymap.hashmap[k] = v
		mymap.keyArray = append(mymap.keyArray, k)
		fmt.Println("Key, Value added succesfully")
	}
}

func (mymap *orderedMap) show() {
	for _, key := range mymap.keyArray {
		fmt.Printf("key: %s  Value:%s\n", key, mymap.hashmap[key])
	}
}

func (mymap *orderedMap) delete(key string) {
	if k, ok := mymap.hashmap[key]; ok {
		delete(mymap.hashmap, key)
		for i, v := range mymap.keyArray {
			if v == k {
				mymap.keyArray = append(mymap.keyArray[:i], mymap.keyArray[i+1:]...)
			}
		}
	} else {
		fmt.Println("Key does not exist")
	}
}
