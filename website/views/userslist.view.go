package views

import (
	"net/http"
	"html/template"
	"github.com/golang/website/models"
)

// CreateUsersListView creates web page with UsersList data  
func CreateUsersListView (w http.ResponseWriter, file []byte, data *models.UserList)  {
	
	t := template.Must(template.New("usersTmpl").Parse(string (file)))
	t.Execute(w, data)
}
