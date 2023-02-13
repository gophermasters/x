package handlers

import (
	"os"
	"net/http"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/theGOURL/warning"
)

const API_URL = os.Getenv("API_URI")

func ServeHTML(w http.ResponseWriter, r http.Request){
	//HTML Req. Var.
	bugname     := r.PostForm.Get("bugname")
	description := r.PostForm.Get("description")
	errorcode   := r.PostForm.Get("errorcode")
	username    := r.PostForm.Get("username")
	imageurl    := r.PostForm.Get("imageurl")

	// Create a new HTML Page to send the request
	http.ServeFile(w, r, "handlers/templates/bugReports_form.html")

	// Define the JSON payload
	payload := map[string]string{"bugname: bugname, "description": description, "errorcode": errorcode, "username": username, "imageurl": imageurl }

	// Encode the payload as JSON
	jsonPayload, err := json.Marshal(payload)
		warning.PRINT_DEFAULT_ERRORS(err,"Error encoding JSON:")

	// Create a new HTTP request
	req, err := http.NewRequest("POST", API_URL , bytes.NewBuffer(jsonPayload))
		warning.PRINT_DEFAULT_ERRORS(err,"Error creating request:")
	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")


	// Print the response status code and status text
	fmt.Println("Response status:", resp.StatusCode, resp.Status)

}
