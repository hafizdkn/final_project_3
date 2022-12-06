package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"final_project_3/category"
	"final_project_3/helper"
	"final_project_3/user"
)

type categoryHandler struct {
	service  category.IService
	response helper.IResponse
}

func NewCategoryHandler(service category.IService) *categoryHandler {
	newResponse := helper.NewResponse
	return &categoryHandler{service: service, response: newResponse}
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {
	var input category.CategoryCreateInput
	var response *helper.Response

	if err := c.ShouldBindJSON(&input); err != nil {
		response = h.response.BadRequestResponse(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	role := currentUser.Role

	category, err := h.service.CreateCategory(input, role)
	if err != nil {
		response = h.response.UnauthorizedResponse(err, "")
		h.response.WriteJsonRespnse(c, response)
		return
	}

	response = h.response.SuccessCreateResponse(category, "Success create category")
	h.response.WriteJsonRespnse(c, response)
}

func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	var input category.CategoryUpdateInput
	var response *helper.Response

	currentUser := c.MustGet("currentUser").(user.User)
	currentUserId := currentUser.ID

	categoryId, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		response = h.response.BadRequestResponse(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response = h.response.BadRequestResponse(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	category, err := h.service.UpdateCategory(input, categoryId, currentUserId)
	if err != nil {
		if err.Error() == "Unauthorized" {
			response = h.response.UnauthorizedResponse(err, "Unauthorized")
			h.response.AbortJsonRespnse(c, response)
			return
		}

		response = h.response.InternalServerError(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	response = h.response.SuccessResponse(category, "Success update category")
	h.response.WriteJsonRespnse(c, response)
}

func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	var response *helper.Response

	currentUser := c.MustGet("currentUser").(user.User)
	_ = currentUser.ID

	categoryId, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		response = h.response.BadRequestResponse(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	err = h.service.DeleteCategory(categoryId)
	if err != nil {
		if err.Error() == "Unauthorized" {
			response = h.response.UnauthorizedResponse(err, "Unauthorized")
			h.response.AbortJsonRespnse(c, response)
			return
		}

		response = h.response.InternalServerError(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	response = h.response.SuccessResponse(nil, "Success delete category")
	h.response.WriteJsonRespnse(c, response)
}

func (h *categoryHandler) GetCategorys(c *gin.Context) {
	categorys, err := h.service.GetCategorys()
	if err != nil {
		response := h.response.InternalServerError(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	response := h.response.SuccessResponse(categorys, "Success get categorys")
	h.response.WriteJsonRespnse(c, response)
}
