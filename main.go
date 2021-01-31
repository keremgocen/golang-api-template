package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(Person{Name: "Kerem", Age: 35})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("KFG-734", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get("KFG-734").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}
