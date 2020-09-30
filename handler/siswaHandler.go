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
	r.GET("/siswa/:id", handler.ShowSiswa)
}

// GetSiswa ...
func (a *SiswaHandler) GetSiswa(c *gin.Context) {
	siswa, _ := a.SiswaEntity.Get(c.Request.Context())
	c.JSON(200, gin.H{"data": siswa})
}

// ShowSiswa ...
func (a *SiswaHandler) ShowSiswa(c *gin.Context) {
	siswa, _ := a.SiswaEntity.Show(c.Param("id"))
	c.JSON(200, gin.H{"data": siswa})
}
