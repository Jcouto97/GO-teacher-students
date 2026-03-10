package handler

import (
	"goProject/internal/models"
	"goProject/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubjectHandler struct {
	service service.SubjectService
}

func NewSubjectHandler(service service.SubjectService) *SubjectHandler {
	return &SubjectHandler{service: service}
}

func (h *SubjectHandler) GetAllSubjects(c *gin.Context) {
	subjects, err := h.service.GetAllSubjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subjects"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subjects})
}

func (h *SubjectHandler) GetSubject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject ID"})
		return
	}

	subject, err := h.service.GetSubjectByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subject})
}

func (h *SubjectHandler) CreateSubject(c *gin.Context) {
	var req models.CreateSubjectRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subject, err := h.service.CreateSubject(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subject"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": subject})
}

func (h *SubjectHandler) UpdateSubject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject ID"})
		return
	}

	var req models.UpdateSubjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subject, err := h.service.UpdateSubject(uint(id), &req)
	if err != nil {
		if err.Error() == "subject not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update subject"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subject})
}

func (h *SubjectHandler) DeleteSubject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject ID"})
		return
	}

	if err := h.service.DeleteSubject(uint(id)); err != nil {
		if err.Error() == "subject not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete subject"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subject deleted successfully"})
}

