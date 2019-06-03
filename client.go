package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Task struct {
	ID       int
	Time     time.Time
	Task     string
	Notes    string
	Priority int
}

func main() {

	post()
	get()
}

func get() {

	//HTTP GET
	resp, err := http.Get("http://localhost:8081/task")
	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func post() {

	var newTask Task
	newTask.ID = 12
	newTask.Time = time.Now()
	newTask.Task = "We need to do some work"
	newTask.Notes = "Here are some notes"
	newTask.Priority = 99

	jsonValue, err := json.Marshal(newTask) //Must convert to byte slice

	resp, err := http.Post("http://localhost:8081/task", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))

}
