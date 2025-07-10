package models

import (
	"time"
)

// File representa un archivo en el sistema
type File struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name" binding:"required"`
	OriginalName string   `json:"original_name" bson:"original_name"`
	Path        string    `json:"path" bson:"path"`
	Size        int64     `json:"size" bson:"size"`
	Type        string    `json:"type" bson:"type"`
	Extension   string    `json:"extension" bson:"extension"`
	Public      bool      `json:"public" bson:"public"`
	UserID      string    `json:"user_id" bson:"user_id"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

// FileUploadRequest representa la solicitud de subida de archivo
type FileUploadRequest struct {
	Name   string `json:"name" binding:"required"`
	Public bool   `json:"public"`
}

// FileResponse es la respuesta pública de un archivo
type FileResponse struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	OriginalName string    `json:"original_name"`
	Size         int64     `json:"size"`
	Type         string    `json:"type"`
	Extension    string    `json:"extension"`
	Public       bool      `json:"public"`
	UserID       string    `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DownloadURL  string    `json:"download_url,omitempty"`
}

// ToResponse convierte un File a FileResponse
func (f *File) ToResponse() FileResponse {
	return FileResponse{
		ID:           f.ID,
		Name:         f.Name,
		OriginalName: f.OriginalName,
		Size:         f.Size,
		Type:         f.Type,
		Extension:    f.Extension,
		Public:       f.Public,
		UserID:       f.UserID,
		CreatedAt:    f.CreatedAt,
		UpdatedAt:    f.UpdatedAt,
	}
} 