package routes

import (
    "arquitecturahex/src/products/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

func RegisterSubjectRoutes(router *gin.Engine, subjectController *controllers.SubjectController) {
    subjectGroup:= router.Group("/subject")
    {
        subjectGroup.GET("/", subjectController.GetAllSubjects)
        subjectGroup.POST("/", subjectController.CreateSubject)
        subjectGroup.PUT("/:id", subjectController.UpdateSubject)
	    subjectGroup.DELETE("/:id", subjectController.DeleteSubject)
    }
}