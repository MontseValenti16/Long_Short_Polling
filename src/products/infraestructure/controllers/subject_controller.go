package controllers

import (
    "arquitecturahex/src/products/application"
    "arquitecturahex/src/products/domain/entities"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type SubjectController struct {
    createSubjectUseCase *application.CreateSubjectUseCase
    viewSubjectUseCase   *application.ViewSubjectUseCase
    updateSubjectUseCase *application.UpdateSubjectUseCase
    deleteSubjectUseCase *application.DeleteSubjectUseCase
}

func NewSubjectController(
    createSubjectUseCase *application.CreateSubjectUseCase,
    viewSubjectUseCase *application.ViewSubjectUseCase,
    updateSubjectUseCase *application.UpdateSubjectUseCase,
    deleteSubjectUseCase *application.DeleteSubjectUseCase,
) *SubjectController {
    return &SubjectController{
        createSubjectUseCase: createSubjectUseCase,
        viewSubjectUseCase:   viewSubjectUseCase,
        updateSubjectUseCase: updateSubjectUseCase,
        deleteSubjectUseCase: deleteSubjectUseCase,
    }
}

func (s *SubjectController) CreateSubject(c *gin.Context) {
    var subject entities.Subject
    if err := c.ShouldBindJSON(&subject); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    if err := s.createSubjectUseCase.Execute(subject); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar la asignatura"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Asignatura creada correctamente"})
}

func (s *SubjectController) GetAllSubjects(c *gin.Context) {
    subjects, err := s.viewSubjectUseCase.Execute()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las asignaturas"})
        return
    }
    c.JSON(http.StatusOK, subjects)
}

func (s *SubjectController) UpdateSubject(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    var subject entities.Subject
    if err := c.ShouldBindJSON(&subject); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    subject.ID = id
    if err := s.updateSubjectUseCase.Execute(id, subject); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar la asignatura"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Asignatura actualizada correctamente"})
}

func (s *SubjectController) DeleteSubject(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    if err := s.deleteSubjectUseCase.Execute(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la asignatura"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Asignatura eliminada correctamente"})
}