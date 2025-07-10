package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"talento-tech-go/models"
)

// Simulación de base de datos en memoria para archivos
var files = make(map[string]models.File)

// GetFiles obtiene todos los archivos
func GetFiles(c *gin.Context) {
	var fileResponses []models.FileResponse
	
	for _, file := range files {
		fileResponse := file.ToResponse()
		fileResponse.DownloadURL = fmt.Sprintf("/api/v1/files/%s/download", file.ID)
		fileResponses = append(fileResponses, fileResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    fileResponses,
		"count":   len(fileResponses),
	})
}

// GetFileByID obtiene un archivo por ID
func GetFileByID(c *gin.Context) {
	id := c.Param("id")
	
	file, exists := files[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Archivo no encontrado",
		})
		return
	}

	fileResponse := file.ToResponse()
	fileResponse.DownloadURL = fmt.Sprintf("/api/v1/files/%s/download", file.ID)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    fileResponse,
	})
}

// UploadFile sube un nuevo archivo
func UploadFile(c *gin.Context) {
	// Obtener el archivo del formulario
	uploadedFile, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "No se pudo obtener el archivo",
			"error":   err.Error(),
		})
		return
	}

	// Obtener datos adicionales del formulario
	name := c.PostForm("name")
	if name == "" {
		name = uploadedFile.Filename
	}
	
	publicStr := c.PostForm("public")
	public := publicStr == "true"

	// Crear directorio de uploads si no existe
	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error al crear directorio de uploads",
			"error":   err.Error(),
		})
		return
	}

	// Generar nombre único para el archivo
	fileID := uuid.New().String()
	extension := filepath.Ext(uploadedFile.Filename)
	fileName := fileID + extension
	filePath := filepath.Join(uploadDir, fileName)

	// Guardar archivo en el sistema
	if err := c.SaveUploadedFile(uploadedFile, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error al guardar el archivo",
			"error":   err.Error(),
		})
		return
	}

	// Crear registro del archivo
	file := models.File{
		ID:           fileID,
		Name:         name,
		OriginalName: uploadedFile.Filename,
		Path:         filePath,
		Size:         uploadedFile.Size,
		Type:         uploadedFile.Header.Get("Content-Type"),
		Extension:    strings.ToLower(extension),
		Public:       public,
		UserID:       "system", // Por ahora hardcodeado, debería venir del contexto de autenticación
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Guardar en memoria
	files[fileID] = file

	fileResponse := file.ToResponse()
	fileResponse.DownloadURL = fmt.Sprintf("/api/v1/files/%s/download", file.ID)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Archivo subido exitosamente",
		"data":    fileResponse,
	})
}

// DeleteFile elimina un archivo
func DeleteFile(c *gin.Context) {
	id := c.Param("id")
	
	file, exists := files[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Archivo no encontrado",
		})
		return
	}

	// Eliminar archivo del sistema
	if err := os.Remove(file.Path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error al eliminar el archivo del sistema",
			"error":   err.Error(),
		})
		return
	}

	// Eliminar de memoria
	delete(files, id)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Archivo eliminado exitosamente",
		"data":    file.ToResponse(),
	})
}

// DownloadFile permite descargar un archivo
func DownloadFile(c *gin.Context) {
	id := c.Param("id")
	
	file, exists := files[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Archivo no encontrado",
		})
		return
	}

	// Verificar si el archivo existe en el sistema
	if _, err := os.Stat(file.Path); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Archivo no encontrado en el sistema",
		})
		return
	}

	// Servir el archivo
	c.File(file.Path)
} 