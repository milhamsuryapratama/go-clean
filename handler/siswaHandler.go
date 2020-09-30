package handler

import (
	"go-clean/domain"

	"github.com/gin-gonic/gin"
)

// SiswaHandler ...
type SiswaHandler struct {
	SiswaEntity domain.SiswaEntity
}

// SiswaHandlerFunc ...
func SiswaHandlerFunc(r *gin.RouterGroup, us domain.SiswaEntity) {
	handler := &SiswaHandler{
		SiswaEntity: us,
	}

	r.GET("/siswa", handler.GetSiswa)
}

// GetSiswa ...
func (a *SiswaHandler) GetSiswa(c *gin.Context) {
	siswa, _ := a.SiswaEntity.Get(c.Request.Context())
	c.JSON(200, gin.H{"data": siswa})
}
