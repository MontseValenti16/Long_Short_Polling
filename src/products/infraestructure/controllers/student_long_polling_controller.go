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
    c.Writer.Header().Set("Content-Type", "application/json")
    c.Writer.Header().Set("Transfer-Encoding", "chunked") 
    c.Writer.Flush() 

    for {
        time.Sleep(15 * time.Second) 

        s.mu.Lock()
        students, err := s.viewStudentUseCase.Execute()
        s.mu.Unlock()

        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los estudiantes"})
            return
        }

        if s.hasChanged(len(students)) {
            s.mu.Lock()
            s.lastStudents = len(students)
            s.mu.Unlock()
            c.SecureJSON(http.StatusOK, gin.H{"message": "Hubo cambios en los estudiantes"})
        } else {
            c.SecureJSON(http.StatusOK, gin.H{"message": "No hubo cambios en los estudiantes"})
        }
        
        c.Writer.Flush() 
    }
}

func (s *StudentLongPollingController) hasChanged(newCount int) bool {
    return newCount != s.lastStudents
}
