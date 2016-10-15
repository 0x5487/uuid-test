package main

import (
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

var (
	mutex = &sync.Mutex{}
	wg    = &sync.WaitGroup{}
)

func main() {
	store := map[string]string{}
	total := 1000000
	wg.Add(total)

	startTime := time.Now()
	for i := 0; i < total; i++ {
		go createUUID(store)
	}
	wg.Wait()
	duration := int64(time.Since(startTime) / time.Millisecond)
	println("total count:", len(store))
	println("duration:", duration)
}

func createUUID(store map[string]string) {
	uuid := uuid.NewV4().String()
	mutex.Lock()
	if _, ok := store[uuid]; ok {
		println("duplicate:", uuid)
	} else {
		store[uuid] = ""
	}
	mutex.Unlock()
	wg.Done()
}
