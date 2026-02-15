package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/mahmud-off/todo-app/pkg"
)

// Sign Up
//
//	@Summary		Sign Up
//	@Description	Creating new user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			name		path		string	true	"Account name"
//	@Param			user_name	path		string	true	"Account user_name"
//	@Param			password	path		string	true	"Account password"
//	@Success		200			{object}	todo.User
//	@Failure		400			{object}	errorResponse
//	@Failure		500			{object}	errorResponse
//	@Router			/auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Sign In
//
//	@Summary		Sign In
//	@Description	Authorization
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user_name	path		string	true	"Account user_name"
//	@Param			password	path		string	true	"Account password"
//	@Success		200			{object}	map[string]interface{}
//	@Failure		400			{object}	handler.errorResponse
//	@Failure		500			{object}	handler.errorResponse
//	@Router			/auth/sign-in [post]
func (h *Handler) singIn(c *gin.Context) {
	var input SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
