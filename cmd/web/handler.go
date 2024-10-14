package main

import(
	"fmt"
	"net/http"
	"log"
	"html/template"

	"github.com/gorilla/mux"
)

func (app *application) allTodo(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/todo"{
		app.notFound(w)
		return
	}
	files := []string{
        "./ui/html/base.tmpl",
        "./ui/html/pages/home.tmpl",
    }

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) oneTodo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	log.Printf(params["id"]);
	w.Write([]byte("Todo List Number ID"))
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