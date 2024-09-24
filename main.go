package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seu-usuario/go-backend-carros/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDatabase() {
	dsn := "host=aws-0-sa-east-1.pooler.supabase.com user=postgres.psuaxdrhaexvptzcpzyf password=dOpJfeX5Z2yvJNj8 dbname=postgres port=6543 sslmode=require"
	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // Desabilita o cache de statements preparados
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connection successfully opened")

	DB.AutoMigrate(&models.Carro{})
}

func main() {
	r := gin.Default()

	connectDatabase()

	r.GET("/carros", func(c *gin.Context) {
		var carros []models.Carro
		if result := DB.Find(&carros); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, carros)
	})

	r.GET("/carros/detalhados", func(c *gin.Context) {
		var carrosDetalhados []models.CarroDetalhado
		if result := DB.Find(&carrosDetalhados); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, carrosDetalhados)
	})

	r.GET("/carros/variacoes", func(c *gin.Context) {
		var carrosVariacao []models.CarroVariacao
		if result := DB.Find(&carrosVariacao); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, carrosVariacao)
	})

	r.Run(":8080")
}
