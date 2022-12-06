package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"final_project_3/helper"
	"final_project_3/task"
	"final_project_3/user"
)

type taskHandler struct {
	service  task.Iservice
	response helper.IResponse
}

func NewTaskHandler(service task.Iservice) *taskHandler {
	newResponse := helper.NewResponse
	return &taskHandler{service: service, response: newResponse}
}

func (h *taskHandler) UpdateStatusCategory(c *gin.Context) {
	var input task.TaskUpdateCategoryInput

	taskId, err := strconv.Atoi(c.Param("taksId"))
	if err != nil {
		response := h.response.BadRequestResponse(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	currentUserId := currentUser.ID

	if err := c.ShouldBindJSON(&input); err != nil {
		response := h.response.BadRequestResponse(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	task, err := h.service.UpdateCategoryTask(input, taskId, currentUserId)
	if err != nil {
		response := h.response.InternalServerError(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	response := h.response.SuccessResponse(task, "Success update status task")
	h.response.WriteJsonRespnse(c, response)
}

func (h *taskHandler) UpdateStatusTask(c *gin.Context) {
	var input task.TaskUpdateStatusInput

	taskId, err := strconv.Atoi(c.Param("taksId"))
	if err != nil {
		response := h.response.BadRequestResponse(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	currentUserId := currentUser.ID

	if err := c.ShouldBindJSON(&input); err != nil {
		response := h.response.BadRequestResponse(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	task, err := h.service.UpdateStatusTask(input, taskId, currentUserId)
	if err != nil {
		response := h.response.InternalServerError(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	response := h.response.SuccessResponse(task, "Success update status task")
	h.response.WriteJsonRespnse(c, response)
}

func (h *taskHandler) DeleteTask(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("taksId"))
	if err != nil {
		response := h.response.BadRequestResponse(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	if err := h.service.DeleteTask(taskId); err != nil {
		response := h.response.InternalServerError(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	response := h.response.SuccessResponse(nil, "Success delete task")
	h.response.WriteJsonRespnse(c, response)
}

func (h *taskHandler) UpdateTask(c *gin.Context) {
	var input task.TaskUpdateInput

	taskId, err := strconv.Atoi(c.Param("taksId"))
	if err != nil {
		response := h.response.BadRequestResponse(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response := h.response.BadRequestResponse(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	task, err := h.service.UpdateTask(input, taskId)
	if err != nil {
		response := h.response.InternalServerError(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	response := h.response.SuccessResponse(task, "Success update task")
	h.response.WriteJsonRespnse(c, response)
}

func (h *taskHandler) GetTasks(c *gin.Context) {
	task, err := h.service.GetTasks()
	if err != nil {
		response := h.response.InternalServerError(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	response := h.response.SuccessResponse(task, "Success get all task")
	h.response.WriteJsonRespnse(c, response)
}

func (h *taskHandler) CreateTask(c *gin.Context) {
	var input task.TaskCreateInput
	var response *helper.Response

	if err := c.ShouldBindJSON(&input); err != nil {
		response := h.response.BadRequestResponse(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	currnetUserId := currentUser.ID

	task, err := h.service.CreateTask(input, currnetUserId)
	if err != nil {
		response := h.response.InternalServerError(err)
		h.response.WriteJsonRespnse(c, response)
		return
	}

	response = h.response.SuccessResponse(task, "Success update task")
	h.response.WriteJsonRespnse(c, response)
}
