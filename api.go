package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Task struct {
	ID int
	Time time.Time
	Task string
	Notes string
	Priority int
}

type Tasks []Task
var AllTasks Tasks

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/task", taskGetAllHandler).Methods("GET") //GET
	r.HandleFunc("/task/{id:[0-9]+}", taskGetHandler).Methods("GET") //GET
	r.HandleFunc("/task", taskPostHandler).Methods("POST") //POST
	r.HandleFunc("/task/{id:[0-9]+}", taskPutHandler).Methods("PUT") //PUT
	r.HandleFunc("/task/{id:[0-9]+}", taskDeleteHandler).Methods("DELETE") //DELETE

	log.Fatal(http.ListenAndServe(":8081", r))
}

func getID(r *http.Request) int {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalln(err)
	}

	return id
}

func taskGetAllHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Method: taskGetAllHandler: GET")

	if AllTasks != nil {
		json.NewEncoder(w).Encode(AllTasks)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func taskGetHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Method: taskGetHandler: GET")

	id := getID(r)

	//Find the single element
	for _, v := range AllTasks {
		if v.ID == id {
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	//Didn't find anything
	w.WriteHeader(http.StatusNotFound)

}

func taskPostHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Method: taskPostHandler: POST")

	var newTask Task
	json.NewDecoder(r.Body).Decode(&newTask)

	newTask.Time = time.Now()
	AllTasks = append(AllTasks, newTask)
}

func taskPutHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Method: taskPutHandler: PUT")

	id := getID(r)

	var newTask Task
	json.NewDecoder(r.Body).Decode(&newTask)

	//Update the element
	for i, v := range AllTasks {
		if v.ID == id {
			v.ID = newTask.ID
			v.Time = time.Now()
			v.Task = newTask.Task
			v.Notes = newTask.Notes
			v.Priority = newTask.Priority

			AllTasks[i] = v
			return
		}
	}

	//Didn't find anything
	w.WriteHeader(http.StatusNotFound)
}

func taskDeleteHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("DELETE Task")

	id := getID(r)

	//Delete the element
	for i, v := range AllTasks {
		if v.ID == id {

			AllTasks = append(AllTasks[:i], AllTasks[i+1:]...)
			return
		}
	}

	//Didn't find anything
	w.WriteHeader(http.StatusNotFound)
}





