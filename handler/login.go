package handler

import (
	"encoding/json"
	"github.com/Elias2389/api-rest/authorization"
	"github.com/Elias2389/api-rest/model"
	"net/http"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp := newResponse(Error, "estructura no válida", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	if !isLoginValid(&data) {
		resp := newResponse(Error, "usuario o contraseña no válidos", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(Error, "no se pudo generar el token", nil)
		responseJSON(w, http.StatusInternalServerError, resp)
		return
	}

	dataToken := map[string]string{"token": token}
	resp := newResponse(Message, "Ok", dataToken)
	responseJSON(w, http.StatusOK, resp)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "contacto@ae.com" && data.Password == "123456"
}
