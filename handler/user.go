package handler

import (
	"crowdfunding-api/helper"
	"crowdfunding-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {

		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := user.FormatUser(newUser, "tokenexample")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
