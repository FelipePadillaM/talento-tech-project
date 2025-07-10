package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"talento-tech-go/models"
)

// Simulación de base de datos en memoria
var users = make(map[string]models.User)

// GetUsers obtiene todos los usuarios
func GetUsers(c *gin.Context) {
	var userResponses []models.UserResponse
	
	for _, user := range users {
		userResponses = append(userResponses, user.ToResponse())
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userResponses,
		"count":   len(userResponses),
	})
}

// GetUserByID obtiene un usuario por ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	
	user, exists := users[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Usuario no encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user.ToResponse(),
	})
}

// CreateUser crea un nuevo usuario
func CreateUser(c *gin.Context) {
	var user models.User
	
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Verificar si el email ya existe
	for _, existingUser := range users {
		if existingUser.Email == user.Email {
			c.JSON(http.StatusConflict, gin.H{
				"success": false,
				"message": "El email ya está registrado",
			})
			return
		}
	}

	// Generar ID único
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Active = true
	
	// Asignar rol por defecto si no se especifica
	if user.Role == "" {
		user.Role = "user"
	}

	// Guardar usuario
	users[user.ID] = user

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Usuario creado exitosamente",
		"data":    user.ToResponse(),
	})
}

// UpdateUser actualiza un usuario existente
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	
	user, exists := users[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Usuario no encontrado",
		})
		return
	}

	var updateData models.User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Actualizar campos permitidos
	if updateData.Name != "" {
		user.Name = updateData.Name
	}
	if updateData.Email != "" {
		// Verificar que el email no esté en uso por otro usuario
		for existingID, existingUser := range users {
			if existingID != id && existingUser.Email == updateData.Email {
				c.JSON(http.StatusConflict, gin.H{
					"success": false,
					"message": "El email ya está en uso por otro usuario",
				})
				return
			}
		}
		user.Email = updateData.Email
	}
	if updateData.Role != "" {
		user.Role = updateData.Role
	}
	
	user.UpdatedAt = time.Now()
	users[id] = user

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Usuario actualizado exitosamente",
		"data":    user.ToResponse(),
	})
}

// DeleteUser elimina un usuario
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	
	user, exists := users[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Usuario no encontrado",
		})
		return
	}

	delete(users, id)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Usuario eliminado exitosamente",
		"data":    user.ToResponse(),
	})
} 