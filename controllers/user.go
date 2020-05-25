package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/koosie0507/pluralsight-go-getting-started/models"
)

/*
UserController is amazing.
*/
type UserController struct {
	userIDPattern *regexp.Regexp
}

func (uc UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (uc *UserController) getAll(w http.ResponseWriter) {
	asJSON(models.GetUsers(), w)
}

func (uc *UserController) get(id int, w http.ResponseWriter) {
	u, err := models.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	asJSON(u, w)
}

func (uc *UserController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("couldn't parse user object"))
		return
	}
	u, err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("couldn't add user"))
		return
	}
	asJSON(u, w)
}

func (uc *UserController) put(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("couldn't parse user object"))
		return
	}
	u, err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("couldn't update user"))
		return
	}
	asJSON(u, w)
}

func (uc *UserController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusGone)
}

func (uc *UserController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func newUserController() *UserController {
	return &UserController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
