package handler

import (
	"encoding/json"
	"net/http"

	"github.com/saalcazar/ceadlbk-info/autorization"
	"github.com/saalcazar/ceadlbk-info/model"
)

type superlogin struct {
	storage StorageSuperlogin
}

func NewServiceSuperLogin(s StorageSuperlogin) *superlogin {
	return &superlogin{s}
}

func (l *superlogin) SuperLogin(lo *model.SuperLogin) bool {
	return l.storage.SuperLogin(lo)
}

func (l *superlogin) superlogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		resp := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}
	data := model.SuperLogin{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp := newResponse(Error, "Estructura no valida", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	if !isSuperLoginValid(&data, l.storage) {
		resp := newResponse(Error, "usuario o contraseña no validos", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	token, err := autorization.GenerateSuperToken(&data)
	if err != nil {
		resp := newResponse(Error, "no se pudo generar el token", nil)
		responseJSON(w, http.StatusInternalServerError, resp)
		return
	}

	dataToken := map[string]string{"token": token}
	resp := newResponse(Message, "Ok", dataToken)
	responseJSON(w, http.StatusOK, resp)
	//return

}

func isSuperLoginValid(data *model.SuperLogin, storage StorageSuperlogin) bool {
	return storage.SuperLogin(data)
}
