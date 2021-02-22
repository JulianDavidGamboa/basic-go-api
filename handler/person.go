package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/JulianDavidGamboa/basic-go-api/model"
	"github.com/labstack/echo/v4"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(c echo.Context) error {

	data := model.Person{}

	// err := json.NewDecoder(r.Body).Decode(&data)

	err := c.Bind(&data)

	if err != nil {
		response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		// responseJSON(w, http.StatusInternalServerError, response)
		return c.JSON(http.StatusInternalServerError, response)
	}

	err = p.storage.Create(&data)

	if err != nil {
		response := newResponse(Error, "Hubo un problema al crear la persona", nil)
		// responseJSON(w, http.StatusInternalServerError, response)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Persona creada correctamente", nil)
	// responseJSON(w, http.StatusCreated, response)
	return c.JSON(http.StatusCreated, response)

}

func (p *person) update(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		response := newResponse(Error, "El id debe ser un número", nil)
		// responseJSON(w, http.StatusBadRequest, response)
		return c.JSON(http.StatusBadRequest, response)

	}

	data := model.Person{}

	// err = json.NewDecoder(r.Body).Decode(&data)

	err = c.Bind(&data)

	if err != nil {
		response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		// responseJSON(w, http.StatusBadRequest, response)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Update(ID, &data)

	if err != nil {

		response := newResponse(Error, "Hubo un problema al obtener las personas", nil)
		// responseJSON(w, http.StatusInternalServerError, response)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Persona actualizada correctamente", nil)
	// responseJSON(w, http.StatusOK, response)
	return c.JSON(http.StatusOK, response)

}

func (p *person) delete(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		response := newResponse(Error, "El id debe ser un número", nil)
		// responseJSON(w, http.StatusBadRequest, response)
		return c.JSON(http.StatusBadRequest, response)

	}

	err = p.storage.Delete(ID)

	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "El ID de la persona no existe", nil)
		// responseJSON(w, http.StatusBadRequest, response)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err != nil {
		response := newResponse(Error, "Ocurrió un error al eliminar el registro", nil)
		// responseJSON(w, http.StatusBadRequest, response)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := newResponse(Message, "Ok", nil)
	// responseJSON(w, http.StatusOK, response)
	return c.JSON(http.StatusOK, response)

}

func (p *person) getAll(c echo.Context) error {

	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "Hubo un problema al obtener las personas", nil)
		//responseJSON(w, http.StatusInternalServerError, response)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Ok", data)
	// responseJSON(w, http.StatusOK, response)
	return c.JSON(http.StatusOK, response)
}
