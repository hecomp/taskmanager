package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hecomp/taskmanager/common"
	"github.com/hecomp/taskmanager/models"
	)
var Tasks = new(tasksController)

type tasksController struct {

}

func (tc *tasksController) Create(w http.ResponseWriter, r *http.Request)  {
	var t models.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadGateway)
		return
	}

	task, err := models.Tasks.Create(t.Name, t.Desc)
	if err != nil {
		common.JsonError(w, err, http.StatusBadGateway)
		return
	}

	res, err := json.Marshal(task)
	if err != nil {
		return
	}
	common.JsonOk(w, res, http.StatusCreated)
}
func (tc *tasksController) Get(w http.ResponseWriter, r *http.Request)  {
	task, err := models.Tasks.FindAll()
	if err != nil {
		return
	}

	res, err := json.Marshal(task)
	if err != nil {
		common.JsonError(w, err, http.StatusBadGateway)
		return
	}
	common.JsonOk(w, res, http.StatusOK)
}
func (tc *tasksController) Show(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]

	task, err := models.Tasks.FindOne(id)
	if err != nil {
		return
	}

	res, err := json.Marshal(task)
	if err != nil {
		common.JsonError(w, err, http.StatusBadGateway)
		return
	}
	common.JsonOk(w, res, http.StatusOK)

}
func (tc *tasksController) Update(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]

	var t models.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadGateway)
		return
	}

	if err := models.Tasks.Update(id, t.Name, t.Desc); err != nil {
		common.JsonError(w, err, http.StatusBadGateway)
		return
	}
	common.JsonStatus(w, http.StatusNoContent)
}
func (tc *tasksController) Delete(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := models.Tasks.DeleteById(id); err != nil {
		common.JsonError(w, err, http.StatusBadGateway)
		return
	}
	common.JsonStatus(w, http.StatusNoContent)
}
