package handler

import (
	"github.com/gin-gonic/gin"

	"final_project_3/helper"
	"final_project_3/user"
)

type userHanlder struct {
	service  user.IService
	response helper.IResponse
}

func NewUserHandler(service user.IService) *userHanlder {
	newResponse := helper.NewResponse
	return &userHanlder{service: service, response: newResponse}
}

func (h *userHanlder) UserRegister(c *gin.Context) {
	var input user.UserRegisterInput
	var response *helper.Response

	if err := c.ShouldBindJSON(&input); err != nil {
		response := h.response.BadRequestResponse(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	user, err := h.service.CreateUser(input)
	if err != nil {
		response := h.response.InternalServerError(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	response = h.response.SuccessCreateResponse(user, "Success register user")
	h.response.WriteJsonRespnse(c, response)
}

func (h *userHanlder) UserLogin(c *gin.Context) {
	var input user.UserLoginInput
	var response *helper.Response
	errUserNotFound := "Invalid username or password"

	if err := c.ShouldBindJSON(&input); err != nil {
		response := h.response.BadRequestResponse(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	user, err := h.service.UserLogin(input)
	if err != nil {
		response := h.response.UnauthorizedResponse(err, errUserNotFound)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	response = h.response.SuccessLoginResponse(nil, user.Token, "Success Login")
	h.response.WriteJsonRespnse(c, response)
}

func (h *userHanlder) UserUpdate(c *gin.Context) {
	var input user.UserUpdateInput
	var response *helper.Response

	if err := c.ShouldBindJSON(&input); err != nil {
		response := h.response.BadRequestResponse(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.EmailCurrentUser = currentUser.Email

	user, err := h.service.UpdateUser(input)
	if err != nil {
		response := h.response.InternalServerError(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	response = h.response.SuccessResponse(user, "Success update user")
	h.response.WriteJsonRespnse(c, response)
}

func (h *userHanlder) DeleteUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	if err := h.service.DeleteUser(userId); err != nil {
		response := h.response.InternalServerError(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	response := h.response.SuccessResponse(nil, "Success update user")
	h.response.WriteJsonRespnse(c, response)
}

func (h *userHanlder) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetUsers()
	if err != nil {
		response := h.response.InternalServerError(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	response := h.response.SuccessResponse(users, "Success get all users")
	h.response.WriteJsonRespnse(c, response)
}
