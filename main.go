package main

import uuid "github.com/satori/go.uuid"

func main() {
	store := map[string]string{}

	for i := 0; i < 5000000; i++ {
		uuid := uuid.NewV4().String()

		if _, ok := store[uuid]; ok {
			println("duplicate:", uuid)
		} else {
			store[uuid] = ""
		}
	}

	println("total count:", len(store))
}
