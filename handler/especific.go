package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/saalcazar/ceadlbk-info/model"
)

type ServiceEspecific struct {
	storage Storageespecific
}

func NewServiceEspecific(s Storageespecific) *ServiceEspecific {
	return &ServiceEspecific{s}
}

func (s *ServiceEspecific) Create(e model.Especifics) error {
	return s.storage.Create(e)
}

func (s *ServiceEspecific) Update(e *model.Especific) error {
	return s.storage.Update(e)
}

func (s *ServiceEspecific) Delete(id uint) error {
	return s.storage.Delete(id)
}

func (s *ServiceEspecific) GetByNameProyect(nameProyect string) (model.Especifics, error) {
	return s.storage.GetByNameProyect(nameProyect)
}

func (s *ServiceEspecific) GetAll() (model.Especifics, error) {
	return s.storage.GetAll()
}

//Handler

// Create
func (s *ServiceEspecific) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	var data model.Especifics
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "El ingreso de datos para crear el objetivo especifico no tienen la estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = s.Create(data)
	if err != nil {
		log.Fatalf("Error: %+v", err)
		response := newResponse(Error, "Hubo un problema al crear el objetivo especifico", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "El pobjetivo especifico se creo correctamente", nil)
	responseJSON(w, http.StatusCreated, response)
}

func (s *ServiceEspecific) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Especific{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "El ingreso de datos para actualizar el objetivo especifico no tienen la estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = s.Update(&data)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al actualizar el objetivo especifico", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "El objetivo especifico se actualizo correctamente", nil)
	responseJSON(w, http.StatusOK, response)
}

func (s *ServiceEspecific) delete(w http.ResponseWriter, r *http.Request) {
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
		response := newResponse(Error, "El ID del objetivo especifico no existe", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "Hubo un problema al eliminar el objetivo especifico", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "OK", nil)
	responseJSON(w, http.StatusOK, response)
}

func (s *ServiceEspecific) getByNameProyect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	nameProyect := r.URL.Query().Get("nameProyect")

	data, err := s.GetByNameProyect(nameProyect)
	if err != nil {
		log.Fatalf("Error: %+v", err)
		response := newResponse(Error, "Hubo un problema al obtener los objetivos específicos", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "OK", data)
	responseJSON(w, http.StatusOK, response)

}

func (s *ServiceEspecific) getall(w http.ResponseWriter, r *http.Request) {
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
