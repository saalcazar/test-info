package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/saalcazar/ceadlbk-info/model"
)

type ServiceActivity struct {
	storage Storageactivity
}

func NewServiceActivity(s Storageactivity) *ServiceActivity {
	return &ServiceActivity{s}
}

func (s *ServiceActivity) Create(a *model.Activity) error {
	return s.storage.Create(a)
}

func (s *ServiceActivity) Update(a *model.Activity) error {
	if a.ID == 0 {
		return errors.New("la actividad no tiene un ID")
	}
	return s.storage.Update(a)
}

func (s *ServiceActivity) Delete(id uint) error {
	return s.storage.Delete(id)
}

func (s *ServiceActivity) GetAll() (model.Activities, error) {
	return s.storage.GetAll()
}

func (s *ServiceActivity) GetByID(id uint) (*model.Activity, error) {
	return s.storage.GetByID(id)
}

//Handler

// Create
func (s *ServiceActivity) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Activity{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "El ingreso de datos para crear la actividad no tienen la estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = s.Create(&data)
	if err != nil {
		log.Fatalf("Error: %+v", err)
		response := newResponse(Error, "Hubo un problema al crear la actividad", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "La actividad se creo correctamente", nil)
	responseJSON(w, http.StatusCreated, response)
}

func (s *ServiceActivity) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Activity{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "El ingreso de datos para actualizar la actividad no tienen la estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = s.Update(&data)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al actualizar la actividad", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "La actividad se actualizo correctamente", nil)
	responseJSON(w, http.StatusOK, response)
}

func (s *ServiceActivity) delete(w http.ResponseWriter, r *http.Request) {
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
		response := newResponse(Error, "El ID de la actividad no existe", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "Hubo un problema al eliminar la actividad", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "OK", nil)
	responseJSON(w, http.StatusOK, response)
}

func (s *ServiceActivity) getall(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := s.storage.GetAll()
	if err != nil {
		log.Fatalf("Error: %+v", err)
		response := newResponse(Error, "Hubo un problema al obtener las actividades", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "OK", data)
	responseJSON(w, http.StatusOK, response)
}

func (s *ServiceActivity) getById(w http.ResponseWriter, r *http.Request) {
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
		response := newResponse(Error, "Hubo un problema al obtener la actividad", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "OK", data)
	responseJSON(w, http.StatusOK, response)
}
