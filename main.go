package main

import (
    "log"
    "arquitecturahex/src/products/application"
    "arquitecturahex/src/products/infraestructure/controllers"
    "arquitecturahex/src/products/infraestructure/db"
    "arquitecturahex/src/products/infraestructure/repositories"
    "arquitecturahex/src/products/infraestructure/routes"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "os"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error al cargar el archivo .env: %v", err)
    }

    db.InitDB()
    defer db.CloseDB()

    router := gin.Default()

    studentRepo := repositories.NewStudentRepositoryImpl()
    subjectRepo := repositories.NewSubjectRepositoryImpl()

    createStudentUseCase := application.NewCreateStudentUseCase(studentRepo)
    viewStudentUseCase := application.NewViewStudentUseCase(studentRepo)
    updateStudentUseCase := application.NewUpdateStudentUseCase(studentRepo)
    deleteStudentUseCase := application.NewDeleteStudentUseCase(studentRepo)
    viewStudentIDUseCase := application.NewViewStudentIDUseCase(studentRepo)

    createSubjectUseCase := application.NewCreateSubjectUseCase(subjectRepo)
    viewSubjectUseCase := application.NewViewSubjectUseCase(subjectRepo)
    updateSubjectUseCase := application.NewUpdateSubjectUseCase(subjectRepo)
    deleteSubjectUseCase := application.NewDeleteSubjectUseCase(subjectRepo)

    studentController := controllers.NewStudentController(
        createStudentUseCase,
        viewStudentUseCase,
        updateStudentUseCase,
        deleteStudentUseCase,
        viewStudentIDUseCase,
    )

    subjectController := controllers.NewSubjectController(
        createSubjectUseCase,
        viewSubjectUseCase,
        updateSubjectUseCase,
        deleteSubjectUseCase,
    )

    studentPollingController := controllers.NewStudentPollingController(viewStudentUseCase)
    subjectPollingController := controllers.NewSubjectPollingController(viewSubjectUseCase)
    studentLongPollingController := controllers.NewStudentLongPollingController(viewStudentUseCase)

    routes.RegisterStudentRoutes(router, studentController)
    routes.RegisterSubjectRoutes(router, subjectController)
    routes.RegisterPollingRoutes(router, studentPollingController, subjectPollingController, studentLongPollingController)

    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" 
    }

    router.Run(":" + port)
}