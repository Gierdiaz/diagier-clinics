package handler

import (
	"net/http"

	"github.com/Gierdiaz/diagier-clinics/internal/domain/patients"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PatientsHandler struct {
	service patients.PatientService
}

func NewPatientsHandler(service patients.PatientService) *PatientsHandler {
	return &PatientsHandler{service: service}
}

func (handler *PatientsHandler) GetAllPatients(c *gin.Context) {
	patients, err := handler.service.GetAllPatients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patients)
}

func (handler *PatientsHandler) GetPatientByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	patient, err := handler.service.GetPatientByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patient)
}

func (handler *PatientsHandler) CreatePatient(c *gin.Context) {
	var dto patients.PatientDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	patient, err := handler.service.CreatePatient(&dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, patient)
}

func (handler *PatientsHandler) UpdatePatient(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var dto patients.PatientDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := handler.service.UpdatePatient(id, &dto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (handler *PatientsHandler) DeletePatient(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := handler.service.DeletePatient(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}