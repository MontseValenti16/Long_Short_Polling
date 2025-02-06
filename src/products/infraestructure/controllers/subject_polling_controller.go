package controllers

import (
    "arquitecturahex/src/products/application"
    "github.com/gin-gonic/gin"
    "net/http"
    "sync"
    "time"
)

type SubjectPollingController struct {
    viewSubjectUseCase *application.ViewSubjectUseCase
    lastSubjects       int
    mu                 sync.Mutex
    changeChan         chan bool
}

func NewSubjectPollingController(viewSubjectUseCase *application.ViewSubjectUseCase) *SubjectPollingController {
    controller := &SubjectPollingController{
        viewSubjectUseCase: viewSubjectUseCase,
        changeChan:         make(chan bool),
    }
    go controller.startPolling()
    return controller
}

func (s *SubjectPollingController) startPolling() {
    ticker := time.NewTicker(5 * time.Second) // Intervalo de polling
    defer ticker.Stop()

    for {
        <-ticker.C
        s.mu.Lock()
        subjects, err := s.viewSubjectUseCase.Execute()
        if err != nil {
            s.mu.Unlock()
            continue
        }

        if s.hasChanged(len(subjects)) {
            s.lastSubjects = len(subjects)
            s.changeChan <- true
        } else {
            s.changeChan <- false
        }
        s.mu.Unlock()
    }
}

func (s *SubjectPollingController) PollSubjects(c *gin.Context) {
    select {
    case changed := <-s.changeChan:
        if changed {
            c.JSON(http.StatusOK, gin.H{"message": "Hubo cambios en las asignaturas"})
        } else {
            c.JSON(http.StatusOK, gin.H{"message": "No hubo cambios en las asignaturas"})
        }
    case <-time.After(5 * time.Second):
        c.JSON(http.StatusOK, gin.H{"message": "No hubo cambios en las asignaturas"})
    }
}

func (s *SubjectPollingController) hasChanged(newCount int) bool {
    return newCount != s.lastSubjects
}