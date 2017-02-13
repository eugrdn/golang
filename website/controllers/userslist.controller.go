package controllers

import (
	"io/ioutil"
	"net/http"
	filepath "path/filepath"

	"github.com/julienschmidt/httprouter"

	"github.com/golang/website/models"
	"github.com/golang/website/views"
)

// UserController is controller for /user/.. routing
type UserController struct{}

// UserC used for import UC structure
var UserC UserController

// UserListRender used for user list render on the web page
func (c UserController) UserListRender(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := &models.UsersArray

	userslistPath, err := filepath.Abs("templates/users.template.html")
	if err != nil {
		panic(err)
	}

	file, err := ioutil.ReadFile(userslistPath)
	if err != nil {
		panic(err)
	}

	views.CreateUsersListView(w, file, data)
}

//Redirector used for redirect cunsumer from request url to '/users' url
func (c UserController) Redirector(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.Redirect(w, r, "/users", http.StatusPermanentRedirect)
}
