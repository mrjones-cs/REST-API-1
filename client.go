package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ClientTask struct {
	ID       int
	Task     string
	Notes    string
	Priority int
}

var curID int

func main() {

	//Create 10 Tasks
	for i := 0; i < 10; i++ {
		post()
	}

	//Retrieve All Tasks
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

	curID++

	var newTask ClientTask
	newTask.ID = curID
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
