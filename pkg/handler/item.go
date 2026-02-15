package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	todo "github.com/mahmud-off/todo-app/pkg"
)

// Create Item
//
//	@Summary		Create Item
//	@Description	Creating new item
//	@Tags			item
//	@Accept			json
//	@Produce		json
//	@Param			title		path		string	true	"List title"
//	@Param			description	path		string	false	"List description"
//	@Param			done path bool false "flag that task done"
//	@Success		200			{object}	map[string]interface{}
//	@Failure		400			{object}	string
//	@Failure		500			{object}	string
//	@Security		OAuth2 password auth
//	@Router			/api/lists/:id/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

// Get All items
//
//	@Summary		Get All Items
//	@Description	getting all items
//	@Tags			item
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	todo.TodoItem
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Security		OAuth2 password auth
//	@Router			/api/lists/:id/items [get]
func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// Get Item By ID
//
//	@Summary		Get Item By Id
//	@Description	getting Item using id
//	@Tags			item
//	@Accept			json
//	@Produce		json
//	@Param			id path int true "item Id"
//	@Success		200	{object}	todo.TodoItem
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Security		OAuth2 password auth
//	@Router			/api/lists/:id/items/:id [get]
func (h *Handler) getItemByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	ItemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	item, err := h.services.TodoItem.GetById(userId, ItemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// Update Item
//
//	@Summary		Update Item
//	@Description	updating item info
//	@Tags			item
//	@Accept			json
//	@Produce		json
//	@Param			title		path		string	false	"Item title"
//	@Param			description	path		string	false	"Item description"
//	@Param			done path bool false "flag that task done"
//	@Success		200	{object}	statusResponse
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Security		OAuth2 password auth
//	@Router			/api/lists/:id/items/:id [put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateItemInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TodoItem.Update(userId, id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// Delete Item
//
//	@Summary		Delete Item
//	@Description	deletting item using id
//	@Tags			item
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	statusResponse
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Security		OAuth2 password auth
//	@Router			/api/lists/:id/items/:id [delete]
func (h *Handler) DeleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	ItemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.services.TodoItem.Delete(userId, ItemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
