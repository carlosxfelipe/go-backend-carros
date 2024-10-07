package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/seu-usuario/go-backend-carros/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDatabase() {
	// Carregar variáveis do .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Montar o DSN usando variáveis de ambiente
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

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

	// Adicionando middleware CORS
	r.Use(cors.Default())

	connectDatabase()

	r.GET("/carros", func(c *gin.Context) {
		var carros []models.Carro
		limitParam := c.DefaultQuery("limit", "10")
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

		if limitParam == "all" {
			if result := DB.Find(&carros); result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}
		} else {
			limit, _ := strconv.Atoi(limitParam)
			offset := (page - 1) * limit
			if result := DB.Offset(offset).Limit(limit).Find(&carros); result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}
		}
		c.JSON(http.StatusOK, carros)
	})

	r.GET("/carros/detalhados", func(c *gin.Context) {
		var carrosDetalhados []models.CarroDetalhado
		limitParam := c.DefaultQuery("limit", "10")
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

		if limitParam == "all" {
			if result := DB.Find(&carrosDetalhados); result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}
		} else {
			limit, _ := strconv.Atoi(limitParam)
			offset := (page - 1) * limit
			if result := DB.Offset(offset).Limit(limit).Find(&carrosDetalhados); result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}
		}
		c.JSON(http.StatusOK, carrosDetalhados)
	})

	r.GET("/carros/variacoes", func(c *gin.Context) {
		var carrosVariacao []models.CarroVariacao
		limitParam := c.DefaultQuery("limit", "10")
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

		if limitParam == "all" {
			if result := DB.Find(&carrosVariacao); result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}
		} else {
			limit, _ := strconv.Atoi(limitParam)
			offset := (page - 1) * limit
			if result := DB.Offset(offset).Limit(limit).Find(&carrosVariacao); result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}
		}
		c.JSON(http.StatusOK, carrosVariacao)
	})

	r.GET("/carros/search", func(c *gin.Context) {
		var carros []models.Carro
		query := DB.Model(&models.Carro{})

		if tipo := c.Query("tipo"); tipo != "" {
			query = query.Where("tipo = ?", tipo)
		}
		if ano := c.Query("ano"); ano != "" {
			if anoInt, err := strconv.Atoi(ano); err == nil {
				query = query.Where("ano = ?", anoInt)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Ano inválido"})
				return
			}
		}
		if marca := c.Query("marca"); marca != "" {
			// Usando ILIKE para busca insensível a maiúsculas e minúsculas
			query = query.Where("marca ILIKE ?", "%"+marca+"%")
		}
		if modelo := c.Query("modelo"); modelo != "" {
			// Usando ILIKE para busca insensível a maiúsculas e minúsculas
			query = query.Where("modelo ILIKE ?", "%"+modelo+"%")
		}
		if combustivel := c.Query("combustivel"); combustivel != "" {
			query = query.Where("combustivel = ?", combustivel)
		}

		if result := query.Find(&carros); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, carros)
	})

	r.GET("/carros/precos", func(c *gin.Context) {
		var carrosDetalhados []models.CarroDetalhado
		query := DB.Model(&models.CarroDetalhado{})

		if precoMin := c.Query("precoMin"); precoMin != "" {
			query = query.Where("preco >= ?", precoMin)
		}
		if precoMax := c.Query("precoMax"); precoMax != "" {
			query = query.Where("preco <= ?", precoMax)
		}
		if dataReferencia := c.Query("dataReferencia"); dataReferencia != "" {
			query = query.Where("data_referencia = ?", dataReferencia)
		}

		if result := query.Find(&carrosDetalhados); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, carrosDetalhados)
	})

	r.Run(":8080")
}
