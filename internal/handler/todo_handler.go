package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NemishGorasiya/Go-Todo/internal/app"
	"github.com/NemishGorasiya/Go-Todo/internal/model"
	"github.com/gorilla/mux"
)

// RegisterRoutes registers all todo-related endpoints
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/todos", listTodos).Methods("GET")
	r.HandleFunc("/todos", createTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")
}

// listTodos godoc
// @Summary Get all todos
// @Tags todos
// @Produce json
// @Success 200 {array} model.Todo
// @Router /todos [get]
func listTodos(w http.ResponseWriter, r *http.Request) {
	todos := app.ListTodos()
	json.NewEncoder(w).Encode(todos)
}

// createTodo godoc
// @Summary Create a new todo
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body model.Todo true "Todo to create"
// @Success 201 {object} model.Todo
// @Failure 400 {string} string "Invalid JSON"
// @Router /todos [post]
func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	newTodo := app.CreateTodo(todo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

// updateTodo godoc
// @Summary Update an existing todo
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body model.Todo true "Updated todo"
// @Success 200 {string} string "Updated successfully"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 404 {string} string "Todo not found"
// @Router /todos/{id} [put]
func updateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var todo model.Todo
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	ok := app.UpdateTodo(uint(id), todo)
	if ok {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Updated successfully"))
	} else {
		http.Error(w, "Todo not found", http.StatusNotFound)
	}
}

// deleteTodo godoc
// @Summary Delete a todo
// @Tags todos
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {string} string "Deleted successfully"
// @Failure 404 {string} string "Todo not found"
// @Router /todos/{id} [delete]
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := app.DeleteTodo(uint(id))

	if ok {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Deleted successfully"))
	} else {
		http.Error(w, "Todo not found", http.StatusNotFound)
	}
}
