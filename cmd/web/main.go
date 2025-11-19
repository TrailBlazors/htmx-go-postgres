package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

type Application struct {
	DB        *sql.DB
	Templates *template.Template
}

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

func main() {
	// Get port from environment (Railway sets this)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Get database URL from environment
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable required")
	}

	// Connect to database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Test connection
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// Create table if not exists
	createTable(db)

	// Parse templates
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	app := &Application{
		DB:        db,
		Templates: tmpl,
	}

	// Setup router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Serve static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes
	r.Get("/", app.homeHandler)
	r.Get("/todos", app.getTodos)
	r.Post("/todos", app.createTodo)
	r.Delete("/todos/{id}", app.deleteTodo)
	r.Put("/todos/{id}/toggle", app.toggleTodo)
	r.Get("/health", healthHandler)

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func createTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS todos (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			completed BOOLEAN DEFAULT FALSE
		);
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}

func (app *Application) homeHandler(w http.ResponseWriter, r *http.Request) {
	app.Templates.ExecuteTemplate(w, "index.html", nil)
}

func (app *Application) getTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := app.DB.Query("SELECT id, title, completed FROM todos ORDER BY id DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	app.Templates.ExecuteTemplate(w, "todo-list.html", todos)
}

func (app *Application) createTodo(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	if title == "" {
		http.Error(w, "Title required", http.StatusBadRequest)
		return
	}

	var id int
	err := app.DB.QueryRow(
		"INSERT INTO todos (title) VALUES ($1) RETURNING id",
		title,
	).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the todo list
	app.getTodos(w, r)
}

func (app *Application) deleteTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, err := app.DB.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return updated list
	app.getTodos(w, r)
}

func (app *Application) toggleTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, err := app.DB.Exec(
		"UPDATE todos SET completed = NOT completed WHERE id = $1",
		id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return updated list
	app.getTodos(w, r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}