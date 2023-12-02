package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	// "strconv"
)

// Define an application struct to hold the application-wide dependencies for the
// web application. For now we'll only include fields for the two custom loggers, but
// we'll add more to it as the build progresses.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":9000", "http Network Address")

	flag.Parse()
	// Use log.New() to create a logger for writing information message . this take
	// three parameter: destinaton to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flag to indicate what addtional information to include (local date and time).
	//  Note that flaged are joined usimg the bitwise OR operator .

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// Initialize a new instance of our application struct, containing the
	// dependencies.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//set the ErrorLog field so that the server now uses the custom errorLog logger in
	// the event of any problems.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// Modified below line for custom Information Logger
	infoLog.Printf("Starting  server at port:%s", *addr)

	err := srv.ListenAndServe()
	// Modified below line for custom Error/Debuggeing Logger
	errorLog.Fatal(err)
}
