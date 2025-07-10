package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"talento-tech-go/handlers"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	// Configurar logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// Configurar Gin
	gin.SetMode(gin.ReleaseMode)
	if os.Getenv("GIN_MODE") == "debug" {
		gin.SetMode(gin.DebugMode)
	}

	// Crear router
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Configurar CORS
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	})

	// Aplicar CORS
	r.Use(func(c *gin.Context) {
		corsConfig.HandlerFunc(c.Writer, c.Request)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Rutas de la API
	setupRoutes(r, logger)

	// Puerto del servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Infof("Servidor iniciando en puerto %s", port)
	if err := r.Run(":" + port); err != nil {
		logger.Fatal("Error al iniciar el servidor:", err)
	}
}

func setupRoutes(r *gin.Engine, logger *logrus.Logger) {
	// Grupo de rutas API v1
	v1 := r.Group("/api/v1")
	{
		// Health check
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Servidor funcionando correctamente",
				"version": "1.0.0",
			})
		})

		// Rutas de usuarios
		users := v1.Group("/users")
		{
			users.GET("/", handlers.GetUsers)
			users.GET("/:id", handlers.GetUserByID)
			users.POST("/", handlers.CreateUser)
			users.PUT("/:id", handlers.UpdateUser)
			users.DELETE("/:id", handlers.DeleteUser)
		}

		// Rutas de archivos
		files := v1.Group("/files")
		{
			files.GET("/", handlers.GetFiles)
			files.GET("/:id", handlers.GetFileByID)
			files.GET("/:id/download", handlers.DownloadFile)
			files.POST("/", handlers.UploadFile)
			files.DELETE("/:id", handlers.DeleteFile)
		}

		// Rutas de autenticación
		auth := v1.Group("/auth")
		{
			auth.POST("/login", handlers.Login)
			auth.POST("/register", handlers.Register)
			auth.POST("/logout", handlers.Logout)
			auth.GET("/me", handlers.GetCurrentUser)
		}
	}

	// Ruta raíz
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Bienvenido a la API de Talento Tech",
			"version": "1.0.0",
			"docs":    "/api/v1/health",
		})
	})
} 