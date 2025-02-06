package controllers

import (
    "arquitecturahex/src/products/application"
    "github.com/gin-gonic/gin"
    "net/http"
    "sync"
    "time"
)

type StudentLongPollingController struct {
    viewStudentUseCase *application.ViewStudentUseCase
    lastStudents       int
    mu                 sync.Mutex
}

func NewStudentLongPollingController(viewStudentUseCase *application.ViewStudentUseCase) *StudentLongPollingController {
    return &StudentLongPollingController{viewStudentUseCase: viewStudentUseCase}
}

func (s *StudentLongPollingController) LongPollStudents(c *gin.Context) {
    timeout := time.After(15 * time.Second)
    ticker := time.Tick(5 * time.Second)

    for {
        select {
        case <-timeout:
            c.JSON(http.StatusOK, gin.H{"message": "No hubo cambios en los estudiantes"})
            return
        case <-ticker:
            s.mu.Lock()
            students, err := s.viewStudentUseCase.Execute()
            if err != nil {
                s.mu.Unlock()
                c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los estudiantes"})
                return
            }

            if s.hasChanged(len(students)) {
                s.lastStudents = len(students)
                s.mu.Unlock()
                c.JSON(http.StatusOK, gin.H{"message": "Hubo cambios en los estudiantes"})
                return
            }
            s.mu.Unlock()
        }
    }
}

func (s *StudentLongPollingController) hasChanged(newCount int) bool {
    return newCount != s.lastStudents
}