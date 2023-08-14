package main

import (
	"context"
	"feyin/go-restapi/handlers"
	"feyin/go-restapi/initialize"
	"log"
	"net/http"
	"os"
	"os/signal"
	//"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main(){
// initialize database and renderer
	db,rnd := initialize.Init()
	//constructing for creating an instance of the handler struct
	//as you can see from the initialize.init() we are getting db,rnd 
	// so we are then passing them both as a parameter to New
	// which then creates an instance of handler
	h := handlers.New(db,rnd)

	const port string = ":9000"

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

/*
	   Creates a new chi router and attaches a logger middleware to it to log all requests.
	   Then it registers a handler for the root URL path ("/") using the GET method, which is the homeHandler function.
	*/

	r := chi.NewRouter()
	// log all requests
	r.Use(middleware.Logger)


// Mounts the subrouter returned by the todoHandlers() function under the "/todo" URL path.
r.Mount("/code-snippets", snippetsHandlers(*h))

/*
		Creates an instance of http.Server with various settings,
		including the address to listen on (Addr), the router to handle requests (Handler), and timeout settings
	*/
	srv := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

/*
		This starts a new goroutine (using go func() { ... }()) to listen and serve incoming HTTP requests.
		 It logs the start of the server and handles any errors that might occur during the server's execution.
	*/
	go func() {
		log.Println("Listening on port ", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()



/*
	   (<-stopChan) waits for a signal to be received on stopChan, which happens when the interrupt signal is triggered (e.g., Ctrl+C).
	    When the signal is received, it triggers a graceful shutdown process. It creates a context with a timeout of 5 seconds,
	   attempts to gracefully shut down the server using srv.Shutdown(ctx), and logs the successful server shutdown.

	*/

	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()
	log.Println("Server gracefully stopped!")
	}
	/*
The snippetsHandlers() function returns an http.Handler (which is a router) for managing routes related to todos.

	It creates a subrouter using chi.NewRouter(), groups the routes using rg.Group(...),

and maps each HTTP method to its corresponding handler function.
*/
func snippetsHandlers(h handlers.Handler) http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/", h.GetAllSnippet)
		r.Get("/{snippetName}", h.GetSnippet)
		r.Post("/", h.AddSnippet)
		r.Put("/{codeid}", h.UpdateSnippet)
		r.Delete("/{id}", h.DeleteSnippet)
	})
	return rg
}




func checkErr(err error) {
	if err != nil {
		log.Fatal(err) //respond with error page or message
	}

}