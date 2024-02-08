package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/saalcazar/ceadlbk-info/model"
)

type ServiceUser struct {
	storage Storageuser
}

func NewServiceUser(s Storageuser) *ServiceUser {
	return &ServiceUser{s}
}

func (s *ServiceUser) Create(u *model.User) error {
	return s.storage.Create(u)
}

func (s *ServiceUser) Update(u *model.User) error {
	if u.ID == 0 {
		return errors.New("el usuario no tiene un ID")
	}
	return s.storage.Update(u)
}

func (s *ServiceUser) Delete(id uint) error {
	return s.storage.Delete(id)
}

func (s *ServiceUser) GetAll() (model.Users, error) {
	return s.storage.GetAll()
}

func (s *ServiceUser) GetByID(id uint) (*model.User, error) {
	return s.storage.GetByID(id)
}

//Handler

// Create
func (s *ServiceUser) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.User{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "El ingreso de datos para crear el usuario no tienen la estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = s.Create(&data)
	if err != nil {
		log.Fatalf("Error: %+v", err)
		response := newResponse(Error, "Hubo un problema al crear el usuario", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "El usuario se creo correctamente", nil)
	responseJSON(w, http.StatusCreated, response)
}

func (s *ServiceUser) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.User{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "El ingreso de datos para actualizar el usuario no tienen la estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = s.Update(&data)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al actualizar el usuario", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "El usuario se actualizo correctamente", nil)
	responseJSON(w, http.StatusOK, response)
}

func (s *ServiceUser) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil && id < 0 {
		response := newResponse(Error, "El ID debe ser un número entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
	}
	err = s.Delete(uint(id))
	if err != nil {
		response := newResponse(Error, "El ID del usuario no existe", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "Hubo un problema al eliminar el usuario", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "OK", nil)
	responseJSON(w, http.StatusOK, response)
}

func (s *ServiceUser) getall(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := s.storage.GetAll()
	if err != nil {
		log.Fatalf("Error: %+v", err)
		response := newResponse(Error, "Hubo un problema al obtener los usuarios", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "OK", data)
	responseJSON(w, http.StatusOK, response)
}

// GetByID
func (s *ServiceUser) getById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil && id < 0 {
		response := newResponse(Error, "El número debe ser positivos", nil)
		responseJSON(w, http.StatusBadRequest, response)
	}
	data, err := s.storage.GetByID(uint(id))
	if err != nil {
		log.Printf("error: %+v", err)
		response := newResponse(Error, "Hubo un problema al obtener el usuario", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "OK", data)
	responseJSON(w, http.StatusOK, response)
}
