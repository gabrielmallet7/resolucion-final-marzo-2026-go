package main

import (
	"final_marzo_2026/database"
	"final_marzo_2026/handlers"
	"final_marzo_2026/middleware"
	"final_marzo_2026/repositories"
	"final_marzo_2026/services"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router               *gin.Engine
	entrenamientoHandler *handlers.EntrenamientoHandler
	rendimientoHandler   *handlers.RendimientoHandler
	mejoraHandler        *handlers.MejoraHandler
	reporteHandler       *handlers.ReporteHandler
)

func main() {
	router = gin.Default()

	databaseInstance, err := dependencies()
	if err != nil {
		log.Fatalf("no se pudo conectar a MongoDB: %v", err)
	}
	defer func() {
		if err := databaseInstance.Disconnect(); err != nil {
			log.Printf("no se pudo cerrar la conexión con MongoDB: %v", err)
		}
	}()

	mappingRoutes()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("no se pudo iniciar el servidor: %v", err)
	}
}

func mappingRoutes() {
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/promedio-velocidad", entrenamientoHandler.CalcularPromedioVelocidad)
		api.POST("/variabilidad-rendimiento", rendimientoHandler.CalcularVariabilidadRendimiento)
		api.POST("/proyeccion-mejora", mejoraHandler.CalcularProyeccionMejora)
		api.GET("/reportes", reporteHandler.ObtenerReportes)
	}
}

func dependencies() (database.DB, error) {
	var databaseInstance database.DB
	databaseInstance = database.NewMongoDB()
	if err := databaseInstance.Connect(); err != nil {
		return nil, err
	}

	var entrenamientoRepo repositories.EntrenamientoRepositoryInterface
	var rendimientoRepo repositories.RendimientoRepositoryInterface

	var entrenamientoService services.EntrenamientoServiceInterface
	var rendimientoService services.RendimientoServiceInterface
	var mejoraService services.MejoraServiceInterface
	var reporteService services.ReporteServiceInterface

	entrenamientoRepo = repositories.NewEntrenamientoRepository(databaseInstance)
	rendimientoRepo = repositories.NewRendimientoRepository(databaseInstance)

	entrenamientoService = services.NewEntrenamientoService(entrenamientoRepo)
	rendimientoService = services.NewRendimientoService(rendimientoRepo)
	mejoraService = services.NewMejoraService()
	reporteService = services.NewReporteService(entrenamientoRepo, rendimientoRepo)

	entrenamientoHandler = handlers.NewEntrenamientoHandler(entrenamientoService)
	rendimientoHandler = handlers.NewRendimientoHandler(rendimientoService)
	mejoraHandler = handlers.NewMejoraHandler(mejoraService)
	reporteHandler = handlers.NewReporteHandler(reporteService)

	return databaseInstance, nil
}
