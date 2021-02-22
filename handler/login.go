package handler

import (
	"net/http"

	"github.com/JulianDavidGamboa/basic-go-api/cmd/authorization"
	"github.com/JulianDavidGamboa/basic-go-api/model"
	"github.com/labstack/echo/v4"
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
	// err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp := newResponse(Error, "Estructura no válida", nil)
		// responseJSON(w, http.StatusBadRequest, resp)
		return c.JSON(http.StatusBadRequest, resp)
	}

	if !isLoginValid(&data) {
		resp := newResponse(Error, "Usuario o contraseña no válidos", nil)
		// responseJSON(w, http.StatusBadRequest, resp)
		return c.JSON(http.StatusBadRequest, resp)
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(Error, "No se pudo generar el token", nil)
		// responseJSON(w, http.StatusInternalServerError, resp)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	dataToken := map[string]string{"token": token}

	resp := newResponse(Message, "Ok", dataToken)
	// responseJSON(w, http.StatusOK, resp)
	return c.JSON(http.StatusOK, resp)

}

func isLoginValid(data *model.Login) bool {
	return data.Email == "julian.gamboa@gmail.com" && data.Password == "123456"
}
