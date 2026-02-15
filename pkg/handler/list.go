package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	todo "github.com/mahmud-off/todo-app/pkg"
)

// Create List
//
//	@Summary		Create List
//	@Description	Creating new list
//	@Tags			list
//	@Accept			json
//	@Produce		json
//	@Param			title		path		string	true	"List title"
//	@Param			description	path		string	false	"List description"
//	@Success		200			{object}	map[string]interface{}
//	@Failure		400			{object}	errorResponse
//	@Failure		500			{object}	errorResponse
//	@Security		@securitydefinitions.oauth2.password OAuth2Password
//	@Router			/api/lists/ [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

// Get All lists
//
//	@Summary		Get All Lists
//	@Description	getting all lists
//	@Tags			list
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	getAllListsResponse
//	@Failure		500	{object}	errorResponse
//	@Security		OAuth2 password auth
//	@Router			/api/lists/ [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})

}

// Get List By ID
//
//	@Summary		Get List By Id
//	@Description	getting list using id
//	@Tags			list
//	@Accept			json
//	@Produce		json
//	@Param			id path int true "list Id"
//	@Success		200	{object}	todo.TodoList
//	@Failure		400	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Security		OAuth2 password auth
//	@Router			/api/lists/:id [get]
func (h *Handler) getListByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// Update List
//
//	@Summary		Update List
//	@Description	updating list info
//	@Tags			list
//	@Accept			json
//	@Produce		json
//	@Param			title		path		string	false	"List title"
//	@Param			description	path		string	false	"List description"
//	@Success		200	{object}	statusResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Security		OAuth2 password auth
//	@Router			/api/lists/:id [put]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateListInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TodoList.Update(userId, id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// Delete List
//
//	@Summary		Delete List
//	@Description	deletting list using id
//	@Tags			list
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	statusResponse
//	@Failure		400	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Security		OAuth2 password auth
//	@Router			/api/lists/:id [delete]
func (h *Handler) DeleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TodoList.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
