package routes

import (
    "arquitecturahex/src/products/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

func RegisterPollingRoutes(router *gin.Engine, studentPollingController *controllers.StudentPollingController, subjectPollingController *controllers.SubjectPollingController, studentLongPollingController *controllers.StudentLongPollingController) {
    pollingGroup := router.Group("/polling")
    {
        pollingGroup.GET("/students", studentPollingController.PollStudents)
        pollingGroup.GET("/subjects", subjectPollingController.PollSubjects)
        pollingGroup.GET("/students/long", studentLongPollingController.LongPollStudents) 
    }
}
