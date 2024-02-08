package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/saalcazar/ceadlbk-info/model"
)

type ServiceReport struct {
	storage Storagereport
}

func NewServiceReport(s Storagereport) *ServiceReport {
	return &ServiceReport{s}
}

func (s *ServiceReport) Create(p *model.Report) error {
	return s.storage.Create(p)
}

func (s *ServiceReport) Update(p *model.Report) error {
	if p.ID == 0 {
		return errors.New("el informe no contiene un ID")
	}
	return s.storage.Update(p)
}

func (s *ServiceReport) Delete(id uint) error {
	return s.storage.Delete(id)
}

func (s *ServiceReport) GetByID(id uint) (*model.Report, error) {
	return s.storage.GetByID(id)
}

// func (s *ServiceReport) GetByTitle(title string) (*model.Report, error) {
// 	return s.storage.GetByTitle(title)
// }

func (s *ServiceReport) GetAll() (model.Reports, error) {
	return s.storage.GetAll()
}

// Create
func (s *ServiceReport) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Report{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "El informe no tiene una estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = s.Create(&data)
	if err != nil {
		log.Fatalf("Error: %+v", err)
		response := newResponse(Error, "Hubo un problema al crear el inform", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Si se pudo crear el informe", nil)
	responseJSON(w, http.StatusCreated, response)
}

// Update
func (s *ServiceReport) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Report{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "El informe no tiene una estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = s.Update(&data)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al crear el informe", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Si se pudo actualizar el informe", nil)
	responseJSON(w, http.StatusOK, response)
}

// DELETE
func (s *ServiceReport) delete(w http.ResponseWriter, r *http.Request) {
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
		response := newResponse(Error, "El ID del informe no existe", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "Hubo un problema al eliminar el informe", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "OK", nil)
	responseJSON(w, http.StatusOK, response)
}

// GetAll
// func (s *ServiceReport) getAll(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		//Devolvemos un JSON con el mensaje de error
// 		response := newResponse(Error, "Método no permitido", nil)
// 		responseJSON(w, http.StatusBadRequest, response)
// 		return
// 	}

// 	l, err := strconv.Atoi(r.URL.Query().Get("l"))
// 	if err != nil {
// 		response := newResponse(Error, "El número debe ser positivo", nil)
// 		responseJSON(w, http.StatusBadRequest, response)
// 	}

// 	pg, err := strconv.Atoi(r.URL.Query().Get("pg"))
// 	if err != nil {
// 		response := newResponse(Error, "El número debe ser positivo", nil)
// 		responseJSON(w, http.StatusBadRequest, response)
// 	}

// 	data, err := s.GetAll(uint(l), uint(pg))
// 	if err != nil {
// 		response := newResponse(Error, "Hubo un problema al obtener los posts", nil)
// 		responseJSON(w, http.StatusInternalServerError, response)
// 		return
// 	}

// 	response := newResponse(Message, "OK", data)
// 	responseJSON(w, http.StatusOK, response)

// }

// GetByID
func (s *ServiceReport) getById(w http.ResponseWriter, r *http.Request) {
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
	data, err := s.GetByID(uint(id))
	if err != nil {
		log.Printf("error: %+v", err)
		response := newResponse(Error, "Hubo un problema al obtener el informe", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "OK", data)
	responseJSON(w, http.StatusOK, response)

}

// // GetByTitle
// func (s *ServiceReport) getByTitle(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		response := newResponse(Error, "Método no permitido", nil)
// 		responseJSON(w, http.StatusBadRequest, response)
// 		return
// 	}
// 	title := r.URL.Query().Get("title")

// 	data, err := s.GetByTitle(title)
// 	if err != nil {
// 		log.Printf("Contenido de: %+v", title)
// 		log.Printf("error: %+v", err)
// 		response := newResponse(Error, "Hubo un problema al obtener el post", nil)
// 		responseJSON(w, http.StatusInternalServerError, response)
// 		return
// 	}

// 	response := newResponse(Message, "OK", data)
// 	responseJSON(w, http.StatusOK, response)

// }

func (s *ServiceReport) getall(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := s.storage.GetAll()
	if err != nil {
		log.Fatalf("Error: %+v", err)
		response := newResponse(Error, "Hubo un problema al obtener los informes", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "OK", data)
	responseJSON(w, http.StatusOK, response)
}
