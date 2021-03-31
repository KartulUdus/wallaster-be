package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kartuludus/wallaster-be/api/models"
	"io/ioutil"
	"net/http"

	"github.com/kartuludus/wallaster-be/api/responses"
	"github.com/kartuludus/wallaster-be/api/utils/formaterror"
)

func (server *Server) CreateCustomer(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.Customer{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveCustomer(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)
}

func (server *Server) GetCustomers(w http.ResponseWriter, r *http.Request) {

	user := models.Customer{}
	users, err := user.FindAllCustomers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func (server *Server) FindCustomers(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	user := models.Customer{}

	users, err := user.FindCustomers(server.DB, vars["search"])
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func (server *Server) GetCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	user := models.Customer{}
	userGotten, err := user.FindCustomerByID(server.DB, vars["id"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
}

func (server *Server) UpdateCustomer(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.Customer{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	fmt.Printf("%#v \n\n  vars \n", user)

	err = user.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedCustomer, err := user.UpdateACustomer(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedCustomer)
}

func (server *Server) DeleteCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	user := models.Customer{}

	_, err := user.DeleteACustomer(server.DB, vars["id"])
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", vars["id"]))
	responses.JSON(w, http.StatusNoContent, "")
}