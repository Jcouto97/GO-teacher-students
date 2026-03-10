package handler

import (
	"goProject/internal/models"
	"goProject/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TeacherHandler struct {
	service service.TeacherService
}

func NewTeacherHandler(service service.TeacherService) *TeacherHandler {
	return &TeacherHandler{service: service}
}

func (h *TeacherHandler) GetAllTeachers(c *gin.Context) {
	teachers, err := h.service.GetAllTeachers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch teachers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": teachers})
}

func (h *TeacherHandler) GetTeacher(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher ID"})
		return
	}

	teacher, err := h.service.GetTeacherByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": teacher})
}

func (h *TeacherHandler) CreateTeacher(c *gin.Context) {
	var req models.CreateTeacherRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teacher, err := h.service.CreateTeacher(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create teacher"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": teacher})
}

func (h *TeacherHandler) UpdateTeacher(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher ID"})
		return
	}

	var req models.UpdateTeacherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teacher, err := h.service.UpdateTeacher(uint(id), &req)
	if err != nil {
		if err.Error() == "teacher not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update teacher"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": teacher})
}

func (h *TeacherHandler) DeleteTeacher(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher ID"})
		return
	}

	if err := h.service.DeleteTeacher(uint(id)); err != nil {
		if err.Error() == "teacher not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete teacher"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted successfully"})
}

