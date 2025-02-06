package routes

import (
    "arquitecturahex/src/products/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

func RegisterStudentRoutes(router *gin.Engine, studentController *controllers.StudentController) {
    studentGroup := router.Group("/students")
    {
        studentGroup.GET("/", studentController.GetAllStudents)
        studentGroup.POST("/", studentController.CreateStudent)
        studentGroup.PUT("/:id", studentController.UpdateStudent)
        studentGroup.DELETE("/:id", studentController.DeleteStudent)
        studentGroup.GET("/:id", studentController.GetStudentByID)
    }
}