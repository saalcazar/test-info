package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/saalcazar/ceadlbk-info/model"
)

type ServiceProjectActivity struct {
	storage Storageprojectactivity
}

func NewServiceProjectActivity(s Storageprojectactivity) *ServiceProjectActivity {
	return &ServiceProjectActivity{s}
}

func (s *ServiceProjectActivity) Create(r model.ProjectActivities) error {
	return s.storage.Create(r)
}

func (s *ServiceProjectActivity) Update(e *model.ProjectActivity) error {
	return s.storage.Update(e)
}

func (s *ServiceProjectActivity) Delete(id uint) error {
	return s.storage.Delete(id)
}

func (s *ServiceProjectActivity) GetByNameProyect(nameProyect string) (model.ProjectActivities, error) {
	return s.storage.GetByNameProyect(nameProyect)
}

func (s *ServiceProjectActivity) GetAll() (model.ProjectActivities, error) {
	return s.storage.GetAll()
}

//Handler

func (s *ServiceProjectActivity) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	var data model.ProjectActivities
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "El ingreso de datos para crear la actividad no tienen la estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = s.Create(data)
	if err != nil {
		log.Fatalf("Error: %+v", err)
		response := newResponse(Error, "Hubo un problema al crear la actividad", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "la actividad se creo correctamente", nil)
	responseJSON(w, http.StatusCreated, response)
}

func (s *ServiceProjectActivity) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.ProjectActivity{}
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

func (s *ServiceProjectActivity) delete(w http.ResponseWriter, r *http.Request) {
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

func (s *ServiceProjectActivity) getByNameProyect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	nameProyect := r.URL.Query().Get("nameProyect")

	data, err := s.GetByNameProyect(nameProyect)
	if err != nil {
		log.Fatalf("Error: %+v", err)
		response := newResponse(Error, "Hubo un problema al obtener las actividades", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "OK", data)
	responseJSON(w, http.StatusOK, response)
}

func (s *ServiceProjectActivity) getall(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := s.storage.GetAll()
	if err != nil {
		log.Fatalf("Error: %+v", err)
		response := newResponse(Error, "Hubo un problema al obtener los items de la solicitud", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "OK", data)
	responseJSON(w, http.StatusOK, response)
}
