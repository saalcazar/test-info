package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/saalcazar/ceadlbk-info/model"
)

type ServiceFounder struct {
	storage Storagefounder
}

func NewServiceFounder(s Storagefounder) *ServiceFounder {
	return &ServiceFounder{s}
}

func (s *ServiceFounder) Create(f *model.Founder) error {
	return s.storage.Create(f)
}

func (s *ServiceFounder) Update(f *model.Founder) error {
	if f.ID == 0 {
		return errors.New("el financiador no tiene un ID")
	}
	return s.storage.Update(f)
}

func (s *ServiceFounder) Delete(id uint) error {
	return s.storage.Delete(id)
}

func (s *ServiceFounder) GetByID(id uint) (*model.Founder, error) {
	return s.storage.GetByID(id)
}

func (s *ServiceFounder) GetAll() (model.Founders, error) {
	return s.storage.GetAll()
}

//Handler

// Create
func (s *ServiceFounder) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Founder{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "El ingreso de datos para founder no tiene la estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = s.Create(&data)
	if err != nil {
		log.Fatalf("Error: %+v", err)
		response := newResponse(Error, "Hubo un problema al crear el financiador", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "El financiador se creo correctamente", nil)
	responseJSON(w, http.StatusCreated, response)
}

func (s *ServiceFounder) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Founder{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "El ingreso de datos para founder no tiene la estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = s.Update(&data)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al actualizar el financiador", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "El financiador se actualizo correctamente", nil)
	responseJSON(w, http.StatusOK, response)
}

func (s *ServiceFounder) delete(w http.ResponseWriter, r *http.Request) {
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
		response := newResponse(Error, "El ID del financiador no existe", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "Hubo un problema al eliminar el financiador", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "OK", nil)
	responseJSON(w, http.StatusOK, response)
}

func (s *ServiceFounder) getById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil && id < 0 {
		response := newResponse(Error, "El número debe ser positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
	}
	data, err := s.storage.GetByID(uint(id))
	if err != nil {
		log.Printf("error: %+v", err)
		response := newResponse(Error, "Hubo un problema al obtener el financiador", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "OK", data)
	responseJSON(w, http.StatusOK, response)
}

func (s *ServiceFounder) getall(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := s.storage.GetAll()
	if err != nil {
		log.Fatalf("Error: %+v", err)
		response := newResponse(Error, "Hubo un problema al obtener financiadores", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "OK", data)
	responseJSON(w, http.StatusOK, response)
}
