package server

import (
	"aytodo/handler"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

const (
	readTimeout       = 5 * time.Minute
	readHeaderTimeout = 30 * time.Second
	writeTimeout      = 5 * time.Minute
)

func check() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}
}

// SetupRoutes provides all the routes that can be used
func SetupRoutes() {
	// Setup the routes
	http.HandleFunc("/", check())
	http.HandleFunc("/tasks", handler.AllTasks())

	http.HandleFunc("/tasks/ordered", handler.OrderedTasks())
	http.HandleFunc("/tasks/ordered/due", handler.OrderedTasksDue())
	http.HandleFunc("/tasks/completed", handler.CompletedTasks())

	http.HandleFunc("/tasks/add", handler.AddTask)
	http.HandleFunc("/tasks/", handler.GetTaskById)

	http.HandleFunc("/tasks/update/", handler.UpdateTask)
}
