package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	todo_app "todo-app"
)

// createList godoc
// @Summary 			Create a new Todo list
// @Security 			ApiKeyAuth
// @Tags 				Lists
// @Description 		Create a new Todo list with the privided data
// @ID 					create-list
// @Accept 				json
// @Produce 			json
// @Param 				input body todo_app.TodoList true "list info"
// @Success 			200 {integer} integer 1
// @Failure 			400,404 {object} errorResponse
// @Failure				500 {object} errorResponse
// @Failure 			default {object} errorResponse
// @Router 				/api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo_app.TodoList
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
	Data []todo_app.TodoList `json:"data"`
}

// getAllLists godoc
// @Summary 			List all Todo lists
// @Security 			ApiKeyAuth
// @Tags 				Lists
// @Description 		Get a list of all Todo lists
// @ID 					get-lists
// @Accept 				json
// @Produce 			json
// @Success 			200 {object} getAllListsResponse
// @Failure 			400,404 {object} errorResponse
// @Failure				500 {object} errorResponse
// @Failure 			default {object} errorResponse
// @Router 				/api/lists [get]
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

// getListById godoc
// @Summary 			Get a List by ID
// @Security 			ApiKeyAuth
// @Tags 				Lists
// @Description 		Get single Todo list by providing its unique ID
// @ID 					get-list
// @Accept 				json
// @Produce 			json
// @Param 				id path integer true "User ID"
// @Success 			200 {object}  todo_app.TodoList
// @Failure 			400,404 {object} errorResponse
// @Failure				500 {object} errorResponse
// @Failure 			default {object} errorResponse
// @Router 				/api/lists/{id} [get]
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// updateList godoc
// @Summary 			Update an existing Todo list
// @Security 			ApiKeyAuth
// @Tags 				Lists
// @Description 		Update an existing Todo list with the provided data
// @ID 					update-list
// @Accept 				json
// @Produce 			json
// @Param 				id path integer true "User ID"
// @Param 				new_input body todo_app.TodoList true "new list info"
// @Success 			200 {object} statusResponse
// @Failure 			400,404 {object} errorResponse
// @Failure				500 {object} errorResponse
// @Failure 			default {object} errorResponse
// @Router 				/api/lists/{id} [put]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	var input todo_app.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// deleteList godoc
// @Summary				Delete an existing Todo list
// @Security 			ApiKeyAuth
// @Tags 				Lists
// @Description			Delete an existing Todo list by ID
// @ID 					delete-list
// @Accept 				json
// @Produce 			json
// @Param 				id path integer true "User ID"
// @Success 			200 {object} statusResponse
// @Failure 			400,404 {object} errorResponse
// @Failure				500 {object} errorResponse
// @Failure 			default {object} errorResponse
// @Router 				/api/lists/{id} [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	err = h.services.TodoList.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
