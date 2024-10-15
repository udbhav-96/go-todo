package main

import(
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/udbhav-96/go-todo/internal/models"
	"github.com/gorilla/mux"
)

func (app *application) allTodo(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/todo"{
		app.notFound(w)
		return
	}

	tasks, err := app.tasks.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Tasks = tasks

	app.render(w, http.StatusOK, "home.tmpl", data)
}

func (app *application) oneTodo(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	params := mux.Vars(r) // Get parameters
	idStr := params["id"] // Extract the "id" parameter

	// Convert the "id" string to an integer
	myId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	// // Log the ID
	// log.Printf("Todo ID: %d", myId)

	// Fetch the task based on the ID
	task, err := app.tasks.Get(myId)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Task = task

	app.render(w, http.StatusOK, "view.tmpl", data)
}

func (app *application) homeCreate(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return 
	}

	title := "React Project"

	id, err := app.tasks.Insert(title)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/todo/%d",id), http.StatusSeeOther)
}

func (app *application) homeDelete(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Todo List"))
}