package handler

import (
	"github.com/Elias2389/api-rest/authorization"
	"github.com/Elias2389/api-rest/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(c echo.Context) error {
	data := model.Login{}
	err := c.Bind(&data)
	if err != nil {
		resp := newResponse(Error, "estructura no válida", nil)
		return c.JSON(http.StatusOK, resp)
	}

	if !isLoginValid(&data) {
		resp := newResponse(Error, "usuario o contraseña no válidos", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(Error, "no se pudo generar el token", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	dataToken := map[string]string{"token": token}
	resp := newResponse(Message, "Ok", dataToken)
	return c.JSON(http.StatusOK, resp)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "contacto@ae.com" && data.Password == "123456"
}
