package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/pluralsight/webservice/models"
)

type taskController struct {
	taskIDPattern *regexp.Regexp
}

func (tc taskController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/tasks" {
		switch r.Method {
		case http.MethodGet:
			tc.getAll(w, r)
		case http.MethodPost:
			tc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := tc.taskIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		switch r.Method {
		case http.MethodGet:
			tc.get(id, w)
		case http.MethodPut:
			tc.put(id, w, r)
		case http.MethodDelete:
			tc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

//GET -> http://localhost:3000/tasks
//body -> none
func (tc *taskController) getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	encodeResponseAsJSON(models.GetTasks(), w) //call reflect.TypeOf(w).String()
}

//GET -> http://localhost:3000/tasks/1
//body -> none
func (tc *taskController) get(id int, w http.ResponseWriter) {
	t, err := models.GetTaskByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	encodeResponseAsJSON(t, w)
}

//POST -> http://localhost:3000/tasks
//body -> raw
//{"Description":"Ford", "Progress":1}
func (tc *taskController) post(w http.ResponseWriter, r *http.Request) {
	t, err := tc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse Task object"))
		return
	}
	t, err = models.AddTask(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	encodeResponseAsJSON(t, w)
}

//GET -> http://localhost:3000/tasks/1
//body -> raw
//{"ID":1,"Description":"Ford", "Progress":1}
func (tc *taskController) put(id int, w http.ResponseWriter, r *http.Request) {
	t, err := tc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse Task object"))
		return
	}
	if id != t.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted Task must match ID in URL"))
		return
	}
	t, err = models.UpdateTask(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	encodeResponseAsJSON(t, w)
}

//GET -> http://localhost:3000/tasks/1
func (tc *taskController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveTaskById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(204)
}
func (tc *taskController) parseRequest(r *http.Request) (models.Task, error) {
	dec := json.NewDecoder(r.Body)
	var t models.Task
	err := dec.Decode(&t)
	if err != nil {
		return models.Task{}, err
	}
	return t, nil
}
func newTaskController() *taskController {
	return &taskController{
		taskIDPattern: regexp.MustCompile(`^/tasks/(\d+)/?`),
	}
}
