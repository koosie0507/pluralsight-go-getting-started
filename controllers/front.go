package controllers

import "net/http"

/*RegisterControllers is not awesome.*/
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
