package controllers

import (
    "arquitecturahex/src/products/application"
    "github.com/gin-gonic/gin"
    "sync"
    "time"
)

type StudentPollingController struct {
    viewStudentUseCase *application.ViewStudentUseCase
    lastStudents       int
    mu                 sync.Mutex
    changeChan         chan bool
}

func NewStudentPollingController(viewStudentUseCase *application.ViewStudentUseCase) *StudentPollingController {
    controller := &StudentPollingController{
        viewStudentUseCase: viewStudentUseCase,
        changeChan:         make(chan bool),
    }
    go controller.startPolling()
    return controller
}

func (s *StudentPollingController) startPolling() {
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()

    for {
        <-ticker.C
        s.mu.Lock()
        students, err := s.viewStudentUseCase.Execute()
        if err != nil {
            s.mu.Unlock()
            continue
        }

        if s.hasChanged(len(students)) {
            s.lastStudents = len(students)
            s.changeChan <- true
        } else {
            s.changeChan <- false
        }
        s.mu.Unlock()
    }
}

func (s *StudentPollingController) PollStudents(c *gin.Context) {
    c.Writer.Header().Set("Content-Type", "text/event-stream")
    c.Writer.Header().Set("Cache-Control", "no-cache")
    c.Writer.Header().Set("Connection", "keep-alive")
    c.Writer.Flush()

    for {
        select {
        case changed := <-s.changeChan:
            if changed {
                c.SSEvent("message", "Hubo cambios en los estudiantes")
            } else {
                c.SSEvent("message", "No hubo cambios en los estudiantes")
            }
            c.Writer.Flush()
        case <-time.After(5 * time.Second):
            c.SSEvent("message", "No hubo cambios en los estudiantes")
            c.Writer.Flush()
        }
    }
}

func (s *StudentPollingController) hasChanged(newCount int) bool {
    return newCount != s.lastStudents
}
