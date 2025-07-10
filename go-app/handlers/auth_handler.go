package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"talento-tech-go/models"
)

// Simulación de sesiones en memoria
var sessions = make(map[string]string) // token -> userID

// Login maneja la autenticación de usuarios
func Login(c *gin.Context) {
	var loginReq models.LoginRequest
	
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Buscar usuario por email
	var foundUser models.User
	userFound := false
	
	for _, user := range users {
		if user.Email == loginReq.Email {
			foundUser = user
			userFound = true
			break
		}
	}

	if !userFound {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Credenciales inválidas",
		})
		return
	}

	// Verificar contraseña (en producción debería usar bcrypt)
	if foundUser.Password != loginReq.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Credenciales inválidas",
		})
		return
	}

	// Verificar si el usuario está activo
	if !foundUser.Active {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Usuario inactivo",
		})
		return
	}

	// Generar token de sesión
	token := uuid.New().String()
	sessions[token] = foundUser.ID

	// Configurar cookie
	c.SetCookie("session_token", token, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login exitoso",
		"data": gin.H{
			"user":  foundUser.ToResponse(),
			"token": token,
		},
	})
}

// Register maneja el registro de nuevos usuarios
func Register(c *gin.Context) {
	var registerReq models.RegisterRequest
	
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Verificar si el email ya existe
	for _, existingUser := range users {
		if existingUser.Email == registerReq.Email {
			c.JSON(http.StatusConflict, gin.H{
				"success": false,
				"message": "El email ya está registrado",
			})
			return
		}
	}

	// Crear nuevo usuario
	user := models.User{
		ID:        uuid.New().String(),
		Email:     registerReq.Email,
		Password:  registerReq.Password, // En producción debería hashearse
		Name:      registerReq.Name,
		Role:      "user",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Guardar usuario
	users[user.ID] = user

	// Generar token de sesión automáticamente
	token := uuid.New().String()
	sessions[token] = user.ID

	// Configurar cookie
	c.SetCookie("session_token", token, 3600, "/", "", false, true)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Usuario registrado exitosamente",
		"data": gin.H{
			"user":  user.ToResponse(),
			"token": token,
		},
	})
}

// Logout maneja el cierre de sesión
func Logout(c *gin.Context) {
	// Obtener token de la cookie
	token, err := c.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "No hay sesión activa",
		})
		return
	}

	// Eliminar sesión
	delete(sessions, token)

	// Eliminar cookie
	c.SetCookie("session_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logout exitoso",
	})
}

// GetCurrentUser obtiene la información del usuario actual
func GetCurrentUser(c *gin.Context) {
	// Obtener token de la cookie
	token, err := c.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "No hay sesión activa",
		})
		return
	}

	// Buscar usuario por token
	userID, exists := sessions[token]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Sesión inválida",
		})
		return
	}

	user, exists := users[userID]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
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