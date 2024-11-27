package main

import (
	"fmt"
	"sync"
	"time"
)

func write(users map[string]int, mutex *sync.RWMutex) {
	var user string
	var age int
	mutex.Lock()
	users[user] = age
	fmt.Println(users)
	mutex.Unlock()
}
func read(users map[string]int, mutex *sync.RWMutex) {
	mutex.RLock()
	fmt.Println(users, "--")
	mutex.RUnlock()
}

func main() {
	users := make(map[string]int)
	rw := &sync.RWMutex{}

	users["Garay"] = 20
	users["Ali"] = 18
	users["Madina"] = 28
	go write(users, rw)
	go read(users, rw)
	go read(users, rw)
	go read(users, rw)
	go read(users, rw)
	time.Sleep(5 * time.Second)
}
