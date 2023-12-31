package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Change the signature of the home handler so it is defined as a method against
// *application.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.

	// ts, err := template.ParseFiles("./ui/html/pages/home.tmpl")
	// if err!= nil{
	// 	log.Println(err.Error())
	// 	http.Error(w, "Internal server error",500)
	// 	return
	// }

	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body.

	// We then use the Execute() method on the template set to write the
	// template content as the response body. The last parameter to Execute()
	// represents any
	// dynamic data that we want to pass in, which for now we'll
	// leave as nil.
	// err = ts.Execute(w,nil)
	// if err!= nil {
	// 	log.Println(err.Error())
	// 	http.Error(w,"Internal Server error", 500)
	// }
	// w.Write([]byte("Hello from Snippetbox"))

	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.
	files := []string{
		"./ui/html/pages/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}

	// Use the template.ParseFiles() function to read the files and store the
	// templates in a template set. Notice that we can pass the slice of file
	// paths as a variadic parameter?
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// Because the home handler function is now a method against application
		// it can access its fields, including the error logger. We'll write the log
		// message to this instead of the standard logger.
		app.errorLog.Println(err.Error())

		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Use the ExecuteTemplate() method to write the content
	// of the "base" template as the response body.
	err = ts.ExecuteTemplate(w, "baseTemplate", nil)
	if err != nil {
		// Also update the code here to use the error logger from the application
		// struct.
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}

// Change the signature of the snippetView handler so it is defined as a method
// against *application.
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a Specific snippet with ID %d...", id)
}

// Change the signature of the snippetCreate handler so it is defined as a method
// against *application.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create an New snippet"))
}
