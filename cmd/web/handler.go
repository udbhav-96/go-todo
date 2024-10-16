package main

import(
	"errors"
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

	err := r.ParseForm()
    if err != nil {
        app.serverError(w, err)
        return
    }

    taskName := r.FormValue("task")

    if taskName == "" {
        http.Error(w, "Task name is required", http.StatusBadRequest)
        return
    }

	_, err = app.tasks.Insert(taskName)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/todo", http.StatusSeeOther)
}

func (app *application) homeDelete(w http.ResponseWriter, r *http.Request){

	 if r.Method == http.MethodPost && r.FormValue("_method") == "DELETE" {
        // Extract the task ID from the URL
        vars := mux.Vars(r)
        idStr := vars["id"]

        id, err := strconv.Atoi(idStr)
        if err != nil {
            http.Error(w, "Invalid task ID", http.StatusBadRequest)
            return
        }

        
        err = app.tasks.Delete(id)
        if err != nil {
            app.serverError(w, err)
            return
        }

        http.Redirect(w, r, "/todo", http.StatusSeeOther)
        return
    }

    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}