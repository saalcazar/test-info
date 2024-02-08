package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/saalcazar/ceadlbk-info/model"
)

type ServiceAccountability struct {
	storage Storageaccountability
}

func NewServiceAccountability(s Storageaccountability) *ServiceAccountability {
	return &ServiceAccountability{s}
}

func (s *ServiceAccountability) Create(a *model.Accountability) error {
	return s.storage.Create(a)
}

func (s *ServiceAccountability) Update(a *model.Accountability) error {
	if a.ID == 0 {
		return errors.New("la rendición no tiene un ID")
	}
	return s.storage.Update(a)
}

func (s *ServiceAccountability) Approved(ua *model.Accountability) error {
	return s.storage.Approved(ua)
}

func (s *ServiceAccountability) Delete(id uint) error {
	return s.storage.Delete(id)
}

func (s *ServiceAccountability) GetAll() (model.Accountabilities, error) {
	return s.storage.GetAll()
}

func (s *ServiceAccountability) GetByID(id uint) (*model.Accountability, error) {
	return s.storage.GetByID(id)
}

//Handler

// Create
func (s *ServiceAccountability) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Accountability{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "El ingreso de datos para crear la rendición no tienen la estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = s.Create(&data)
	if err != nil {
		log.Fatalf("Error: %+v", err)
		response := newResponse(Error, "Hubo un problema al crear la rendición", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "La rendición se creo correctamente", nil)
	responseJSON(w, http.StatusCreated, response)
}

func (s *ServiceAccountability) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Accountability{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "El ingreso de datos para actualizar la rendición no tienen la estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = s.Update(&data)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al actualizar la rendición", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "La rendición se actualizo correctamente", nil)
	responseJSON(w, http.StatusOK, response)
}

func (s *ServiceAccountability) approved(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Accountability{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "El ingreso de datos para actualizar la solicitud no tienen la estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = s.Approved(&data)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al actualizar la solicitud", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "La solicitud se actualizo correctamente", nil)
	responseJSON(w, http.StatusOK, response)
}

func (s *ServiceAccountability) delete(w http.ResponseWriter, r *http.Request) {
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
		response := newResponse(Error, "El ID de la rendición no existe", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "Hubo un problema al eliminar la rendición", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "OK", nil)
	responseJSON(w, http.StatusOK, response)
}

func (s *ServiceAccountability) getall(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := s.storage.GetAll()
	if err != nil {
		log.Fatalf("Error: %+v", err)
		response := newResponse(Error, "Hubo un problema al obtener las rendiciones", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "OK", data)
	responseJSON(w, http.StatusOK, response)
}

func (s *ServiceAccountability) getById(w http.ResponseWriter, r *http.Request) {
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
		response := newResponse(Error, "Hubo un problema al obtener la rendición de cuentas", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "OK", data)
	responseJSON(w, http.StatusOK, response)
}
