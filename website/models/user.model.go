package models

// User represents model of user
type User struct {
	FirstName string
	Email     string
	Username  string
}

// UserList represents array of Users
type UserList []User

var UsersArray UserList = UserList{
	{"jhonny aviony", "info@johnyaviony.com", "jhonny"},
	{"mark hammock", "markhammock@gmail.com", "HAMM"},
	{"richicl cross", "crossstation@gmail.com", "cute_richicl"},
	{"flinn's shack", "help@flinns.com", "flinn1989"},
	{"filip justic", "filip.justic@gmail.com", "justicjuscic"},
	{"summer evans", "summer@gmail.com", "SUM42"},
	{"marco s. diaz", "s.diaz@gmail.com", "dIaZ"},
	{"ricky stwerat", "r.stwerat@gmail.com", "r1cky_st3w4rd"},
}
