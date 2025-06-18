package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	_ "github.com/andres-dev4/go-flower-api/docs" // <-- Ajustar al nombre real de tu módulo
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var log = logrus.New()

// @title Flower API
// @version 1.0
// @description API que devuelve flores y saluda usuarios.
// @host localhost:8080
// @BasePath /
// @contact.name Andrés
// @contact.email andres@example.com

func main() {
	var Version = "dev" // Sobrescrito en tiempo de compilación
	// Configurar logrus
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
	log.Infof("Iniciando API - versión: %s", Version)

	router := gin.Default()

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Endpoints
	router.GET("/flowers", GetFlowers)
	router.POST("/greet", PostGreet)

	log.Info("Iniciando servidor en :8080")
	router.Run(":8080")
}

// GetFlowers godoc
// @Summary Obtener lista de flores
// @Produce json
// @Success 200 {object} map[string][]string
// @Router /flowers [get]
func GetFlowers(c *gin.Context) {
	log.WithFields(logrus.Fields{
		"endpoint": "/flowers",
		"method":   "GET",
	}).Info("Solicitud recibida get flowers")

	flowers := []string{"Rosa", "Tulipán", "Lirio", "Margarita"}
	c.JSON(http.StatusOK, gin.H{"flowers": flowers})
}

// PostGreet godoc
// @Summary Saludar a un usuario
// @Accept json
// @Produce json
// @Param input body map[string]string true "Nombre"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /greet [post]
func PostGreet(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
	}
	if err := c.BindJSON(&body); err != nil {
		log.WithError(err).Warn("Entrada inválida")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	log.WithFields(logrus.Fields{
		"endpoint": "/greet",
		"method":   "POST",
		"name":     body.Name,
	}).Info("Saludo generado postgree")

	message := "Hola, un gusto " + body.Name
	c.JSON(http.StatusOK, gin.H{"message": message})
}
