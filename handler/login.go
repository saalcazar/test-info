package handler

import (
	"encoding/json"
	"net/http"

	"github.com/saalcazar/ceadlbk-info/autorization"
	"github.com/saalcazar/ceadlbk-info/model"
)

type login struct {
	storage Storagelogin
}

func NewServiceLogin(s Storagelogin) *login {
	return &login{s}
}

func (l *login) Login(lo *model.Login) (bool, []*model.DataUser, error) {
	return l.storage.Login(lo)
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		resp := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}
	data := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp := newResponse(Error, "Estructura no valida", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	isValid, user, err := isLoginValid(&data, l.storage)
	if err != nil {
		errorMsg := "Error al verificar el inicio de sesión" + err.Error()
		resp := newResponse(Error, errorMsg, nil)
		responseJSON(w, http.StatusInternalServerError, resp)
		return
	}

	if !isValid {
		resp := newResponse(Error, "Usuario o contraseña no válidos", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	token, err := autorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(Error, "no se pudo generar el token", nil)
		responseJSON(w, http.StatusInternalServerError, resp)
		return
	}

	dataToken := map[string]string{"token": token, "user": user[0].Name, "profile": user[0].Profile, "signature": user[0].Signature, "nameProject": user[0].NameProyect}
	resp := newResponse(Message, "Ok", dataToken)
	responseJSON(w, http.StatusOK, resp)
	//return

}

func isLoginValid(data *model.Login, storage Storagelogin) (bool, []*model.DataUser, error) {
	return storage.Login(data)
}
