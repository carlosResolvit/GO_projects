package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()
	tc := newTaskController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
	http.Handle("/tasks", *tc)
	http.Handle("/tasks/", *tc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
