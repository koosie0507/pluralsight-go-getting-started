package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

/*RegisterControllers is not awesome.*/
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func asJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
